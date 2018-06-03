package snowgauge_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/snowgauge"

	"github.com/teddyking/snowflake/api"
)

var _ = Describe("Snowgauge", func() {
	Describe("Flakes", func() {
		var (
			suiteSummaries []*api.SuiteSummary
		)

		BeforeEach(func() {
			suiteSummaries = []*api.SuiteSummary{
				&api.SuiteSummary{
					Codebase: "test-codebase",
					Commit:   "1",
					Tests: []*api.Test{
						&api.Test{
							Name:     "test-name",
							State:    api.Test_PASSED,
							Location: "/some/path_test.go:1",
						},
						&api.Test{
							Name:     "test-name",
							State:    api.Test_PASSED,
							Location: "/some/path_test.go:2",
						},
					},
				},
				&api.SuiteSummary{
					Codebase: "test-codebase",
					Commit:   "1",
					Tests: []*api.Test{
						&api.Test{
							Name:     "test-name",
							State:    api.Test_FAILED,
							Location: "/some/path_test.go:1",
						},
						&api.Test{
							Name:     "test-name",
							State:    api.Test_PASSED,
							Location: "/some/path_test.go:2",
						},
					},
				},
			}

		})

		It("returns all flaky tests for a given codebase and commit", func() {
			flakes, err := Flakes(suiteSummaries)
			Expect(err).NotTo(HaveOccurred())

			Expect(flakes).To(HaveLen(1))
			Expect(flakes[0].Name).To(Equal("test-name"))
			Expect(flakes[0].Location).To(Equal("/some/path_test.go:1"))
		})

		When("passed less than 2 summaries", func() {
			BeforeEach(func() {
				suiteSummaries = suiteSummaries[:1]
			})

			It("returns no flakes and doesn't error", func() {
				flakes, err := Flakes(suiteSummaries)
				Expect(err).NotTo(HaveOccurred())

				Expect(len(flakes)).To(Equal(0))
			})
		})

		When("passed suite summaries from two different codebases", func() {
			BeforeEach(func() {
				suiteSummaries = append(suiteSummaries, &api.SuiteSummary{
					Codebase: "test-codebase-2",
					Commit:   "1",
				})
			})

			It("returns an error", func() {
				_, err := Flakes(suiteSummaries)
				Expect(err).To(MatchError("cannot detect flakes across different codebases - 'test-codebase' and 'test-codebase-2'"))
			})
		})

		When("passed suite summaries with two different commits", func() {
			BeforeEach(func() {
				suiteSummaries = append(suiteSummaries, &api.SuiteSummary{
					Codebase: "test-codebase",
					Commit:   "2",
				})
			})

			It("returns an error", func() {
				_, err := Flakes(suiteSummaries)
				Expect(err).To(MatchError("cannot detect flakes across different commits - '1' and '2'"))
			})
		})
	})
})
