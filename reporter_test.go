package snowflake_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake"

	"time"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

var _ = Describe("SnowflakeReporter", func() {
	var (
		r            *SnowflakeReporter
		suiteSummary *types.SuiteSummary
		ginkgoConfig config.GinkgoConfigType
		specSummary  *types.SpecSummary
	)

	BeforeEach(func() {
		r = NewReporter()
	})

	Describe("SpecSuiteWillBegin", func() {
		BeforeEach(func() {
			ginkgoConfig = config.GinkgoConfigType{}
			suiteSummary = &types.SuiteSummary{
				SuiteDescription: "A Sweet Suite",
			}
			r.SpecSuiteWillBegin(ginkgoConfig, suiteSummary)
		})

		It("records the suite name", func() {
			Expect(r.Suite.Name).To(Equal("A Sweet Suite"))
		})
	})

	Describe("SpecDidComplete", func() {
		var (
			timeBeforeTest int
			specState      types.SpecState
			failure        types.SpecFailure
		)

		BeforeEach(func() {
			specState = types.SpecStatePassed
			timeBeforeTest = time.Now().Nanosecond()
		})

		JustBeforeEach(func() {
			specSummary = &types.SpecSummary{
				ComponentTexts: []string{
					"[Top level]",
					"Integration",
					"CLI",
					"when passed an invalid flag",
					"exits with a status of 1",
				},
				State:   specState,
				Failure: failure,
			}
			r.SpecDidComplete(specSummary)
			Expect(len(r.Suite.Tests)).To(Equal(1))
		})

		It("records the spec name", func() {
			recordedSpecName := r.Suite.Tests[0].Name

			Expect(recordedSpecName).To(Equal("Integration CLI when passed an invalid flag exits with a status of 1"))
		})

		It("records the time at which the spec completed", func() {
			recordedCompletedAt := r.Suite.Tests[0].CompletedAt.Nanosecond()

			// TODO: make test not suck
			Expect(recordedCompletedAt).To(BeNumerically(">", timeBeforeTest))
		})

		Context("when the spec passed", func() {
			BeforeEach(func() {
				specState = types.SpecStatePassed
			})

			It("records the spec state", func() {
				recordedSpecState := r.Suite.Tests[0].State

				Expect(recordedSpecState).To(Equal("passed"))
			})
		})

		Context("when the spec failed", func() {
			BeforeEach(func() {
				specState = types.SpecStateFailed
				failure = types.SpecFailure{
					Message: "Expected x to equal y",
					Location: types.CodeLocation{
						FileName:   "github.com/teddyking/fail/fail_test.go",
						LineNumber: 100,
					},
				}
			})

			It("records the spec state", func() {
				recordedSpecState := r.Suite.Tests[0].State

				Expect(recordedSpecState).To(Equal("failed"))
			})

			It("records the failure message", func() {
				recordedSpecFailureMessage := r.Suite.Tests[0].Failure.Message

				Expect(recordedSpecFailureMessage).To(Equal("Expected x to equal y\ngithub.com/teddyking/fail/fail_test.go:100"))
			})
		})

		Context("when the spec skipped", func() {
			BeforeEach(func() {
				specState = types.SpecStateSkipped
			})

			It("records the spec state", func() {
				recordedSpecState := r.Suite.Tests[0].State

				Expect(recordedSpecState).To(Equal("skipped"))
			})
		})

		Context("when the spec pending", func() {
			BeforeEach(func() {
				specState = types.SpecStatePending
			})

			It("records the spec state", func() {
				recordedSpecState := r.Suite.Tests[0].State

				Expect(recordedSpecState).To(Equal("pending"))
			})
		})

		Context("when the spec panicked", func() {
			BeforeEach(func() {
				specState = types.SpecStatePanicked
				failure = types.SpecFailure{
					Message: "Panicked",
					Location: types.CodeLocation{
						FileName:   "github.com/teddyking/fail/fail_test.go",
						LineNumber: 100,
					},
				}
			})

			It("records the spec state", func() {
				recordedSpecState := r.Suite.Tests[0].State

				Expect(recordedSpecState).To(Equal("panicked"))
			})

			It("records the failure message", func() {
				recordedSpecFailureMessage := r.Suite.Tests[0].Failure.Message

				Expect(recordedSpecFailureMessage).To(Equal("Panicked\ngithub.com/teddyking/fail/fail_test.go:100"))
			})
		})

		Context("when the spec timed out", func() {
			BeforeEach(func() {
				specState = types.SpecStateTimedOut
				failure = types.SpecFailure{
					Message: "Timed out",
					Location: types.CodeLocation{
						FileName:   "github.com/teddyking/fail/fail_test.go",
						LineNumber: 100,
					},
				}
			})

			It("records the spec state", func() {
				recordedSpecState := r.Suite.Tests[0].State

				Expect(recordedSpecState).To(Equal("timedout"))
			})

			It("records the failure message", func() {
				recordedSpecFailureMessage := r.Suite.Tests[0].Failure.Message

				Expect(recordedSpecFailureMessage).To(Equal("Timed out\ngithub.com/teddyking/fail/fail_test.go:100"))
			})
		})

		Context("when the spec invalid", func() {
			BeforeEach(func() {
				specState = types.SpecStateInvalid
			})

			It("records the spec state", func() {
				recordedSpecState := r.Suite.Tests[0].State

				Expect(recordedSpecState).To(Equal("invalid"))
			})
		})
	})
})
