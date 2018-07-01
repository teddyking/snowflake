package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"net/http"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Logging", func() {
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
		webPort = 6000 + GinkgoParallelNode()
		serverEnv = []string{fmt.Sprintf("PORT=%d", serverPort)}
		webEnv = []string{fmt.Sprintf("SERVERPORT=%d", serverPort), fmt.Sprintf("PORT=%d", webPort)}
	})

	JustBeforeEach(func() {
		serverSession = startSnowflakeServer(serverEnv...)
		webSession = startSnowflakeWeb(webEnv...)
		ensureConnectivityToPort(serverPort)
		ensureConnectivityToPort(webPort)
	})

	AfterEach(func() {
		serverSession.Kill()
		webSession.Kill()
	})

	It("logs at info level", func() {
		Expect(serverSession.Out).To(gbytes.Say(`"level":"info"`))
		Expect(webSession.Out).To(gbytes.Say(`"level":"info"`))
	})

	When("the DEBUG serverEnv var is set to true", func() {
		BeforeEach(func() {
			serverEnv = append(serverEnv, "DEBUG=true")
			webEnv = append(webEnv, "DEBUG=true")
		})

		It("logs at debug level", func() {
			Expect(serverSession.Out).To(gbytes.Say(`"level":"debug"`))
			Expect(webSession.Out).To(gbytes.Say(`"level":"debug"`))
		})

		It("logs server requests", func() {
			http.Get(fmt.Sprintf("http://0.0.0.0:%d", webPort))

			Expect(serverSession.Out).To(gbytes.Say(`"level":"debug","method":"/api.Flaker/List"`))
			Expect(webSession.Out).To(gbytes.Say(`"level":"debug","method":"/api.Flaker/List"`))
		})
	})
})
