package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("snowflakeweb Integration", func() {
	var (
		env        []string
		webSession *gexec.Session
	)

	BeforeEach(func() {
		env = []string{}
	})

	JustBeforeEach(func() {
		webSession = startSnowflakeWeb(env...)
	})

	AfterEach(func() {
		webSession.Kill()
	})

	Describe("listen port", func() {
		It("listens on TCP port 2930 by default", func() {
			ensureConnectivityToPort(2930)
		})

		When("the PORT env var is set", func() {
			BeforeEach(func() {
				env = []string{"PORT=8080"}
			})

			It("listens on the port specified by the PORT env var", func() {
				ensureConnectivityToPort(8080)
			})
		})
	})

	Describe("server port", func() {
		// super lame tests ...

		It("connects to the server on TCP port 2929 by default", func() {
			Eventually(webSession.Err).Should(gbytes.Say("connecting to snowflake server on port: 2929"))
		})

		When("the SERVERPORT env var is set", func() {
			BeforeEach(func() {
				env = []string{"SERVERPORT=2000"}
			})

			It("connects on the port specified by the SERVERPORT env var", func() {
				Eventually(webSession.Err).Should(gbytes.Say("connecting to snowflake server on port: 2000"))
			})
		})
	})
})
