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

	Describe("Get", func() {
		var (
			ctx context.Context
			req *api.GetRequest
			res *api.GetResponse
		)

		BeforeEach(func() {
			var err error
			ctx = context.Background()
			req = &api.GetRequest{
				Codebase: "test-codebase",
				Commit:   "test-commit",
				Location: "test-location",
			}

			fakeStore.GetReturns(&api.Test{Name: "cake"}, nil)

			res, err = snowflakeServer.Get(ctx, req)
			Expect(err).NotTo(HaveOccurred())
		})

		It("gets the test from the store", func() {
			Expect(fakeStore.GetCallCount()).To(Equal(1))

			codebase, commit, location := fakeStore.GetArgsForCall(0)
			Expect(codebase).To(Equal("test-codebase"))
			Expect(commit).To(Equal("test-commit"))
			Expect(location).To(Equal("test-location"))
		})

		It("returns the test", func() {
			Expect(res.Test.Name).To(Equal("cake"))
		})

		When("the store returns an error", func() {
			BeforeEach(func() {
				fakeStore.GetReturns(nil, errors.New("error-geting-test"))
			})

			It("returns the error", func() {
				_, err := snowflakeServer.Get(ctx, req)
				Expect(err).To(MatchError("error-geting-test"))
			})
		})
	})
})
