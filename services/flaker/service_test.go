package flaker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/services/flaker"

	"context"
	"errors"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/services/flaker/flakerfakes"
)

var _ = Describe("Flaker Service", func() {
	var (
		fakeStore                *flakerfakes.FakeStore
		fakeFlakeAnalyser        FlakeAnalyser
		fakeFlakeAnalyserReports []*api.Report
		flakerService            *Service
	)

	BeforeEach(func() {
		fakeFlakeAnalyser = func(reports []*api.Report) ([]*api.Flake, error) {
			fakeFlakeAnalyserReports = reports

			return []*api.Flake{
				&api.Flake{TestDescription: "It is a flake"},
				&api.Flake{TestDescription: "It is another flake"},
			}, nil
		}

		fakeStore = new(flakerfakes.FakeStore)
	})

	JustBeforeEach(func() {
		flakerService = New(fakeStore, fakeFlakeAnalyser)
	})

	Describe("List", func() {
		var (
			ctx     context.Context
			req     *api.FlakerListReq
			reports []*api.Report
		)

		BeforeEach(func() {
			ctx = context.Background()
			req = &api.FlakerListReq{}

			reports = []*api.Report{
				&api.Report{Description: "Integration Suite"},
				&api.Report{Description: "Integration Suite"},
			}

			fakeStore.ListReportsReturns(reports, nil)
		})

		It("retrieves all known reports from the store", func() {
			_, err := flakerService.List(ctx, req)
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeStore.ListReportsCallCount()).To(Equal(1))
		})

		It("sends the reports to the flake analyser", func() {
			_, err := flakerService.List(ctx, req)
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeFlakeAnalyserReports).To(Equal(reports))
		})

		It("returns the flakes", func() {
			res, err := flakerService.List(ctx, req)
			Expect(err).NotTo(HaveOccurred())

			Expect(res.Flakes).To(HaveLen(2))
			Expect(res.Flakes[0].TestDescription).To(Equal("It is a flake"))
		})

		When("the store returns an error", func() {
			BeforeEach(func() {
				fakeStore.ListReportsReturns(nil, errors.New("error-listing-reports"))
			})

			It("returns the error", func() {
				_, err := flakerService.List(ctx, req)
				Expect(err).To(MatchError("error-listing-reports"))
			})
		})

		When("the flake analyser returns an error", func() {
			BeforeEach(func() {
				fakeFlakeAnalyser = func(reports []*api.Report) ([]*api.Flake, error) {
					return []*api.Flake{}, errors.New("error-analysing-flakes")
				}
			})

			It("returns the error", func() {
				_, err := flakerService.List(ctx, req)
				Expect(err).To(MatchError("error-analysing-flakes"))
			})
		})
	})
})
