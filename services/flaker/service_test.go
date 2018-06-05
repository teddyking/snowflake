package flaker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/services/flaker"

	"context"
	"errors"

	"github.com/teddyking/snowflake/api"
)

var _ = Describe("Flaker Service", func() {
	var (
		flakeService               *Service
		fakeFlakeAnalyser          FlakeAnalyser
		fakeFlakeAnalyserSummaries []*api.SuiteSummary
	)

	BeforeEach(func() {
		fakeFlakeAnalyser = func(summaries []*api.SuiteSummary) ([]*api.Flake, error) {
			fakeFlakeAnalyserSummaries = summaries

			return []*api.Flake{
				&api.Flake{Name: "It is a flake"},
				&api.Flake{Name: "It is another flake"},
			}, nil
		}
	})

	JustBeforeEach(func() {
		flakeService = New(fakeFlakeAnalyser)
	})

	Describe("List", func() {
		var (
			ctx            context.Context
			req            *api.FlakeListRequest
			suiteSummaries []*api.SuiteSummary
		)

		BeforeEach(func() {
			ctx = context.Background()

			suiteSummaries = []*api.SuiteSummary{
				&api.SuiteSummary{Codebase: "test-codebase-1"},
				&api.SuiteSummary{Codebase: "test-codebase-2"},
			}
		})

		JustBeforeEach(func() {
			req = &api.FlakeListRequest{
				Suites: suiteSummaries,
			}
		})

		It("passes suite summaries to the flake analyser", func() {
			_, err := flakeService.List(ctx, req)
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeFlakeAnalyserSummaries).To(Equal(suiteSummaries))
		})

		It("uses the flake analyser to detect flakes", func() {
			res, err := flakeService.List(ctx, req)
			Expect(err).NotTo(HaveOccurred())

			Expect(res.Flakes).To(HaveLen(2))
			Expect(res.Flakes[0].Name).To(Equal("It is a flake"))
		})

		When("the flake analyser returns an error", func() {
			BeforeEach(func() {
				fakeFlakeAnalyser = func(summaries []*api.SuiteSummary) ([]*api.Flake, error) {
					return []*api.Flake{}, errors.New("error-analysing-flakes")
				}
			})

			It("returns the error", func() {
				_, err := flakeService.List(ctx, req)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
