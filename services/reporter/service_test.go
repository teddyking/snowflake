package reporter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/services/reporter"

	"context"
	"errors"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/services/reporter/reporterfakes"
)

var _ = Describe("ReporterService", func() {
	var (
		fakeStore       *reporterfakes.FakeStore
		reporterService *Service
	)

	BeforeEach(func() {
		fakeStore = new(reporterfakes.FakeStore)
		reporterService = New(fakeStore)
	})

	Describe("Create", func() {
		var (
			ctx context.Context
			req *api.ReporterCreateReq
		)

		BeforeEach(func() {
			ctx = context.Background()
			req = &api.ReporterCreateReq{Report: &api.Report{Description: "cake"}}

			_, err := reporterService.Create(ctx, req)
			Expect(err).NotTo(HaveOccurred())
		})

		It("creates the provided report in the store", func() {
			Expect(fakeStore.CreateReportCallCount()).To(Equal(1))
			storedReport := fakeStore.CreateReportArgsForCall(0)

			Expect(storedReport).To(Equal(req.Report))
		})

		When("the store returns an error", func() {
			BeforeEach(func() {
				fakeStore.CreateReportReturns(errors.New("error-creating-report"))
			})

			It("returns the error", func() {
				_, err := reporterService.Create(ctx, req)
				Expect(err).To(MatchError("error-creating-report"))
			})
		})
	})
})
