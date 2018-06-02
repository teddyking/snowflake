package suite_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/services/suite"

	"context"
	"errors"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/services/suite/suitefakes"
)

var _ = Describe("SuiteService", func() {
	var (
		fakeStore    *suitefakes.FakeStore
		suiteService *SuiteService
	)

	BeforeEach(func() {
		fakeStore = new(suitefakes.FakeStore)

		suiteService = New(fakeStore)
	})

	Describe("Create", func() {
		var (
			ctx context.Context
			req *api.SuiteCreateRequest
		)

		BeforeEach(func() {
			ctx = context.Background()
			req = &api.SuiteCreateRequest{Summary: &api.SuiteSummary{Name: "cake"}}

			_, err := suiteService.Create(ctx, req)
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
				_, err := suiteService.Create(ctx, req)
				Expect(err).To(MatchError("error-creating-summary"))
			})
		})
	})

	Describe("List", func() {
		var (
			ctx context.Context
			req *api.SuiteListRequest
			res *api.SuiteListResponse
		)

		BeforeEach(func() {
			var err error
			ctx = context.Background()
			req = &api.SuiteListRequest{}

			fakeStore.ListReturns([]*api.SuiteSummary{&api.SuiteSummary{Name: "cake"}}, nil)

			res, err = suiteService.List(ctx, req)
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
				_, err := suiteService.List(ctx, req)
				Expect(err).To(MatchError("error-listing-summaries"))
			})
		})
	})

	Describe("Get", func() {
		var (
			ctx context.Context
			req *api.SuiteGetRequest
			res *api.SuiteGetResponse
		)

		BeforeEach(func() {
			var err error
			ctx = context.Background()
			req = &api.SuiteGetRequest{
				Codebase: "test-codebase",
				Commit:   "test-commit",
				Location: "test-location",
			}

			fakeStore.GetReturns(&api.Test{Name: "cake"}, nil)

			res, err = suiteService.Get(ctx, req)
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
				_, err := suiteService.Get(ctx, req)
				Expect(err).To(MatchError("error-geting-test"))
			})
		})
	})
})
