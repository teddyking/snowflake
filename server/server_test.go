package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/server"

	"context"
	"errors"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/server/serverfakes"
)

var _ = Describe("Server", func() {
	var (
		fakeStore       *serverfakes.FakeStore
		snowflakeServer *Server
	)

	BeforeEach(func() {
		fakeStore = new(serverfakes.FakeStore)

		snowflakeServer = New(fakeStore)
	})

	Describe("Create", func() {
		var (
			ctx context.Context
			req *api.CreateRequest
		)

		BeforeEach(func() {
			ctx = context.Background()
			req = &api.CreateRequest{Summary: &api.SuiteSummary{Name: "cake"}}

			_, err := snowflakeServer.Create(ctx, req)
			Expect(err).NotTo(HaveOccurred())
		})

		It("creates the provided summary in the store", func() {
			Expect(fakeStore.CreateCallCount()).To(Equal(1))
			storedSummary := fakeStore.CreateArgsForCall(0)

			Expect(storedSummary).To(Equal(req.Summary))
		})

		When("the store returns an error", func() {
			BeforeEach(func() {
				fakeStore.CreateReturns(errors.New("error-creating-summary"))
			})

			It("returns the error", func() {
				_, err := snowflakeServer.Create(ctx, req)
				Expect(err).To(MatchError("error-creating-summary"))
			})
		})
	})
})
