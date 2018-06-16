package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/onsi/gomega/gexec"
	"github.com/teddyking/snowflake/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var _ = Describe("snowflake Integration", func() {
	var (
		env           []string
		serverSession *gexec.Session
	)

	BeforeEach(func() {
		env = []string{}
	})

	JustBeforeEach(func() {
		serverSession = startSnowflakeServer(env...)
	})

	AfterEach(func() {
		serverSession.Kill()
	})

	Describe("listen port", func() {
		It("listens on TCP port 2929 by default", func() {
			ensureConnectivityToPort(2929)
		})

		When("the PORT env var is set", func() {
			var port int

			BeforeEach(func() {
				port = 5000 + GinkgoParallelNode()
				env = []string{fmt.Sprintf("PORT=%d", port)}
			})

			It("listens on the port specified by the PORT env var", func() {
				ensureConnectivityToPort(port)
			})
		})
	})

	Describe("TLS communication", func() {
		var port int

		BeforeEach(func() {
			port = 5000 + GinkgoParallelNode()
			env = []string{fmt.Sprintf("PORT=%d", port)}
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

				env = append(env, []string{fmt.Sprintf("TLSKEYPATH=%s", tlsKeyPath), fmt.Sprintf("TLSCRTPATH=%s", tlsCrtPath)}...)
			})

			JustBeforeEach(func() {
				ensureConnectivityToPort(port)
			})

			AfterEach(func() {
				Expect(os.RemoveAll(tmpCertDir)).To(Succeed())
			})

			It("works", func() {
				creds, err := credentials.NewClientTLSFromFile(tlsCrtPath, "")
				Expect(err).NotTo(HaveOccurred())

				hostAddress := fmt.Sprintf("localhost:%d", port)
				conn, err := grpc.Dial(hostAddress, grpc.WithTransportCredentials(creds), grpc.WithTimeout(time.Minute))
				Expect(err).NotTo(HaveOccurred())

				// ensure the connection is valid by making an arbitrary API request
				flakerService := api.NewFlakerClient(conn)
				_, err = flakerService.List(context.Background(), &api.FlakerListReq{})
				Expect(err).NotTo(HaveOccurred())
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
