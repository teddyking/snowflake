package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/onsi/gomega/gexec"
	"github.com/teddyking/snowflake/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var _ = Describe("Security", func() {
	Describe("TLS communication", func() {
		var (
			serverPort    int
			webPort       int
			serverEnv     []string
			webEnv        []string
			serverSession *gexec.Session
			webSession    *gexec.Session
		)

		BeforeEach(func() {
			serverPort = 5000 + GinkgoParallelNode()
			serverEnv = []string{fmt.Sprintf("PORT=%d", serverPort)}
			webPort = 6000 + GinkgoParallelNode()
			webEnv = []string{fmt.Sprintf("SERVERPORT=%d", serverPort), fmt.Sprintf("PORT=%d", webPort)}
		})

		JustBeforeEach(func() {
			serverSession = startSnowflakeServer(serverEnv...)
			ensureConnectivityToPort(serverPort)
			webSession = startSnowflakeWeb(webEnv...)
			ensureConnectivityToPort(webPort)
		})

		AfterEach(func() {
			serverSession.Kill()
			webSession.Kill()
		})

		When("the TLSKEYPATH and TLSCRTPATH env vars point to a valid .key and .crt", func() {
			var (
				tmpCertDir string
				tlsCrtPath string
			)

			BeforeEach(func() {
				skipIfNotOnPath("certstrap")

				var err error
				tmpCertDir, err = ioutil.TempDir("", "snowflake-test-certs")
				Expect(err).NotTo(HaveOccurred())

				var tlsKeyPath string
				tlsKeyPath, tlsCrtPath, err = generateKeyAndCrt(tmpCertDir)
				Expect(err).NotTo(HaveOccurred())

				serverEnv = append(serverEnv, []string{fmt.Sprintf("TLSKEYPATH=%s", tlsKeyPath), fmt.Sprintf("TLSCRTPATH=%s", tlsCrtPath)}...)
				webEnv = append(webEnv, []string{fmt.Sprintf("TLSCRTPATH=%s", tlsCrtPath)}...)
			})

			AfterEach(func() {
				Expect(os.RemoveAll(tmpCertDir)).To(Succeed())
			})

			It("configures snowflake for communication over TLS", func() {
				creds, err := credentials.NewClientTLSFromFile(tlsCrtPath, "")
				Expect(err).NotTo(HaveOccurred())

				hostAddress := fmt.Sprintf("localhost:%d", serverPort)
				conn, err := grpc.Dial(hostAddress, grpc.WithTransportCredentials(creds))
				Expect(err).NotTo(HaveOccurred())

				// ensure server is configured for TLS by making an arbitrary API request
				flakerService := api.NewFlakerClient(conn)
				_, err = flakerService.List(context.Background(), &api.FlakerListReq{})
				Expect(err).NotTo(HaveOccurred())

				// ensure web is configured for TLS by checking for an HTTP 200
				webAddress := fmt.Sprintf("localhost:%d", webPort)
				res, err := http.Get(fmt.Sprintf("http://%s", webAddress))
				Expect(err).NotTo(HaveOccurred())
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
})

func generateKeyAndCrt(outDir string) (string, string, error) {
	certstrapCommands := [][]string{
		[]string{"init", "--passphrase", "", "--common-name", "snowflake test ca"},
		[]string{"request-cert", "--passphrase", "", "--domain", "localhost"},
		[]string{"sign", "--CA", "snowflake test ca", "localhost"},
	}

	for _, certstrapCommand := range certstrapCommands {
		if err := runCertstrap(outDir, certstrapCommand...); err != nil {
			return "", "", err
		}
	}

	tlsKeyPath := filepath.Join(outDir, "localhost.key")
	tlsCrtPath := filepath.Join(outDir, "localhost.crt")

	return tlsKeyPath, tlsCrtPath, nil
}

func runCertstrap(outDir string, args ...string) error {
	command := exec.Command(
		"certstrap",
		append([]string{"--depot-path", outDir}, args...)...,
	)
	command.Stdout = GinkgoWriter
	command.Stderr = GinkgoWriter

	return command.Run()
}
