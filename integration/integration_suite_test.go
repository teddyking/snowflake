package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"net"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/onsi/gomega/gexec"
)

var pathToSnowflake, pathToSnowflakeWeb string

var _ = BeforeSuite(func() {
	var err error

	pathToSnowflake, err = gexec.Build("github.com/teddyking/snowflake/cmd/snowflake")
	Expect(err).NotTo(HaveOccurred())

	pathToSnowflakeWeb, err = gexec.Build("github.com/teddyking/snowflake/cmd/snowflakeweb")
	Expect(err).NotTo(HaveOccurred())

	SetDefaultEventuallyTimeout(time.Hour)
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

func startSnowflakeServer(env ...string) *gexec.Session {
	command := exec.Command(pathToSnowflake)
	command.Stdout = GinkgoWriter
	command.Stderr = GinkgoWriter

	if len(env) > 0 {
		command.Env = env
	}

	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}

func startSnowflakeWeb(env ...string) *gexec.Session {
	command := exec.Command(pathToSnowflakeWeb)
	command.Stdout = GinkgoWriter
	command.Stderr = GinkgoWriter
	command.Env = []string{fmt.Sprintf("STATICDIR=%s", filepath.Join("..", "web", "static"))}

	if len(env) > 0 {
		command.Env = append(command.Env, env...)
	}

	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}

func ensureConnectivityToPort(port int) {
	Eventually(func() error {
		_, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
		return err
	}).Should(Succeed())
}

func skipIfNotOnPath(executable string) {
	_, err := exec.LookPath("certstrap")
	if err != nil {
		Skip(fmt.Sprintf("this test requires the '%s' executable to be found on $PATH", executable))
	}
}
