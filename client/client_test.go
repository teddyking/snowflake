package client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/client"

	"net/http"

	"github.com/onsi/gomega/ghttp"
	"github.com/teddyking/snowflake"
)

var _ = Describe("Client", func() {
	Describe("PostSuite", func() {
		var (
			server          *ghttp.Server
			snowflakeClient *Client
		)

		BeforeEach(func() {
			server = ghttp.NewServer()
			snowflakeClient = New(server.URL())

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyJSONRepresenting(&snowflake.Suite{}),
					ghttp.RespondWith(http.StatusCreated, nil),
				),
			)
		})

		AfterEach(func() {
			server.Close()
		})

		It("POSTs the suite to the snowflake server", func() {
			testSuite := &snowflake.Suite{}

			Expect(snowflakeClient.PostSuite(testSuite)).To(Succeed())
			Expect(server.ReceivedRequests()).To(HaveLen(1))
		})

		Context("when the server does not return an HTTP 201", func() {
			BeforeEach(func() {
				server.Reset()

				server.AppendHandlers(
					ghttp.RespondWith(http.StatusInternalServerError, nil),
				)
			})

			It("returns an error", func() {
				testSuite := &snowflake.Suite{}

				Expect(snowflakeClient.PostSuite(testSuite)).To(MatchError("unexpected HTTP response code: 500"))
				Expect(server.ReceivedRequests()).To(HaveLen(1))
			})
		})
	})
})
