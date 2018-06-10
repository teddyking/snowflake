package integration_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/onsi/gomega/gexec"
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
			BeforeEach(func() {
				env = []string{"PORT=2930"}
			})

			It("listens on the port specified by the PORT env var", func() {
				ensureConnectivityToPort(2930)
			})
		})
	})
})
