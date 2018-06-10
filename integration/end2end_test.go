package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("end-to-end functionality", func() {
	var (
		pathToExampleSuite string
		serverSession      *gexec.Session
		webSession         *gexec.Session
		serverPort         int
		webPort            int
	)

	BeforeEach(func() {
		pathToExampleSuite = filepath.Join("..", "examples", "examplesuite")
		serverPort = 2929
		webPort = 2930

		serverSession = startSnowflakeServer(fmt.Sprintf("PORT=%d", serverPort))
		webSession = startSnowflakeWeb(fmt.Sprintf("SERVERPORT=%d", serverPort), fmt.Sprintf("PORT=%d", webPort))

		ensureConnectivityToPort(serverPort)
		ensureConnectivityToPort(webPort)
	})

	AfterEach(func() {
		webSession.Kill()
		serverSession.Kill()
	})

	It("detects flakes and displays them on the web UI", func() {
		runGinkgoOn(pathToExampleSuite, serverPort, false)
		checkWebUIFor("Found 0 flakes", webPort)

		runGinkgoOn(pathToExampleSuite, serverPort, false)
		checkWebUIFor("Found 0 flakes", webPort)

		introduceFlakeTo(pathToExampleSuite)
		defer removeFlakeFrom(pathToExampleSuite)

		runGinkgoOn(pathToExampleSuite, serverPort, true)
		checkWebUIFor("Found 1 flakes", webPort)
	})
})

func runGinkgoOn(suitePath string, serverPort int, expectToFail bool) {
	command := exec.Command("ginkgo", suitePath)
	command.Stdout = GinkgoWriter
	command.Stderr = GinkgoWriter

	err := command.Run()

	if !expectToFail {
		Expect(err).NotTo(HaveOccurred())
	}
}

func checkWebUIFor(content string, port int) {
	res, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
	Expect(err).NotTo(HaveOccurred())
	Expect(res.StatusCode).To(Equal(http.StatusOK))

	body, err := ioutil.ReadAll(res.Body)
	Expect(err).NotTo(HaveOccurred())
	Expect(res.Body.Close()).To(Succeed())

	Expect(string(body)).To(ContainSubstring(content))
}

func introduceFlakeTo(suitePath string) {
	flakeFilePath := filepath.Join(suitePath, "assets", "flake.txt")
	Expect(ioutil.WriteFile(flakeFilePath, []byte("flake"), 0644)).To(Succeed())
}

func removeFlakeFrom(suitePath string) {
	flakeFilePath := filepath.Join(suitePath, "assets", "flake.txt")
	Expect(ioutil.WriteFile(flakeFilePath, []byte("notflake"), 0644)).To(Succeed())
}
