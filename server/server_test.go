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

	Describe("List", func() {
		var (
			ctx context.Context
			req *api.ListRequest
			res *api.ListResponse
		)

		BeforeEach(func() {
			var err error
			ctx = context.Background()
			req = &api.ListRequest{}

			fakeStore.ListReturns([]*api.SuiteSummary{&api.SuiteSummary{Name: "cake"}}, nil)

			res, err = snowflakeServer.List(ctx, req)
			Expect(err).NotTo(HaveOccurred())
		})

		It("retrieves all summaries from the store", func() {
			Expect(fakeStore.ListCallCount()).To(Equal(1))
		})

		It("returns the summaries", func() {
			Expect(res.SuiteSummaries).To(HaveLen(1))
			Expect(res.SuiteSummaries[0].Name).To(Equal("cake"))
		})

		When("the store returns an error", func() {
			BeforeEach(func() {
				fakeStore.ListReturns(nil, errors.New("error-listing-summaries"))
			})

			It("returns the error", func() {
				_, err := snowflakeServer.List(ctx, req)
				Expect(err).To(MatchError("error-listing-summaries"))
			})
		})
	})
})
