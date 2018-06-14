package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"

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

	Describe("server port", func() {
		var port int

		BeforeEach(func() {
			port = 5000 + GinkgoParallelNode()
			env = []string{fmt.Sprintf("SERVERPORT=%d", port)}
		})

		// super lame test ...
		It("connects on the port specified by the SERVERPORT env var", func() {
			Eventually(webSession.Err).Should(gbytes.Say(fmt.Sprintf("connecting to snowflake server on port: %d", port)))
		})
	})

	Describe("static dir", func() {
		// super lame tests ...
		It("serves static assets from the path sepecified by STATICDIR", func() {
			Eventually(webSession.Err).Should(gbytes.Say("serving static assets from: ../web/static"))
		})

		When("the STATICDIR env var is set", func() {
			BeforeEach(func() {
				env = []string{fmt.Sprintf("STATICDIR=staticdir")}
			})

			It("serves static assets from the path sepecified by STATICDIR", func() {
				Eventually(webSession.Err).Should(gbytes.Say("serving static assets from: staticdir"))
			})
		})
	})
})
