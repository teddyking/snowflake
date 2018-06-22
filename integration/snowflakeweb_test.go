package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"

	"github.com/onsi/gomega/gexec"
	"github.com/shirou/gopsutil/net"
)

var _ = Describe("snowflakeweb Integration", func() {
	var (
		webEnv     []string
		webSession *gexec.Session
	)

	BeforeEach(func() {
		webEnv = []string{}
	})

	JustBeforeEach(func() {
		webSession = startSnowflakeWeb(webEnv...)
	})

	AfterEach(func() {
		webSession.Kill()
	})

	Describe("listen port", func() {
		It("listens on TCP port 2930 by default", func() {
			ensureConnectivityToPort(2930)
		})

		When("the PORT env var is set", func() {
			var webPort int

			BeforeEach(func() {
				webPort = 6000 + GinkgoParallelNode()
				webEnv = []string{fmt.Sprintf("PORT=%d", webPort)}
			})

			It("listens on the port specified by the PORT env var", func() {
				ensureConnectivityToPort(webPort)
			})
		})
	})

	Describe("server port", func() {
		var (
			serverPort    int
			webPort       int
			serverEnv     []string
			serverSession *gexec.Session
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
			ensureConnectivityToPort(webPort)
		})

		AfterEach(func() {
			serverSession.Kill()
		})

		It("connects to the server on the port specified by the SERVERPORT env var", func() {
			Expect(connectionEstablishedTo(net.Addr{IP: "127.0.0.1", Port: uint32(serverPort)})).To(BeTrue())
		})
	})
})

func connectionEstablishedTo(remoteAddress net.Addr) bool {
	conns, err := net.Connections("tcp4")
	Expect(err).NotTo(HaveOccurred())

	for _, conn := range conns {
		if conn.Status == "ESTABLISHED" {
			if conn.Raddr == remoteAddress {
				return true
			}
		}
	}

	return false
}
