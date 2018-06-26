package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"
	"fmt"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/teddyking/snowflake/api"
	"google.golang.org/grpc"
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

		It("logs at info level", func() {
			Expect(serverSession.Out).To(gbytes.Say(`"level":"info"`))
		})

		When("the DEBUG env var is set to true", func() {
			BeforeEach(func() {
				env = append(env, "DEBUG=true")
			})

			It("logs at debug level", func() {
				Expect(serverSession.Out).To(gbytes.Say(`"level":"debug"`))
			})

			It("logs server requests", func() {
				conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
				Expect(err).NotTo(HaveOccurred())

				flakerService := api.NewFlakerClient(conn)
				flakerService.List(context.Background(), &api.FlakerListReq{})

				Expect(serverSession.Out).To(gbytes.Say(`"level":"debug","method":"/api.Flaker/List"`))
			})
		})
	})
})
