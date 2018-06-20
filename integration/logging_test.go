package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Logging", func() {
	Describe("Server logging", func() {
		var (
			port          int
			env           []string
			serverSession *gexec.Session
		)

		BeforeEach(func() {
			port = 5000 + GinkgoParallelNode()
			env = []string{fmt.Sprintf("PORT=%d", port)}
		})

		JustBeforeEach(func() {
			serverSession = startSnowflakeServer(env...)
			ensureConnectivityToPort(port)
		})

		AfterEach(func() {
			serverSession.Kill()
		})

		FIt("logs at info level", func() {
			Expect(serverSession.Out).To(gbytes.Say(`"level":"info"`))
		})

		When("the DEBUG env var is set to true", func() {
			BeforeEach(func() {
				env = append(env, "DEBUG=true")
			})

			It("logs all debug messages, as well as error messages", func() {
				Expect(serverSession.Out).To(gbytes.Say("debug"))
			})
		})
	})
})
