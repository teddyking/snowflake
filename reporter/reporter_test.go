package reporter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/reporter"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
	"github.com/teddyking/snowflake"
	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/reporter/reporterfakes"
)

var _ = Describe("SnowflakeReporter", func() {
	var (
		r            *SnowflakeReporter
		suiteClient  *reporterfakes.FakeSuiteClient
		codebase     string
		commit       string
		suiteSummary *types.SuiteSummary
		ginkgoConfig config.GinkgoConfigType
		specSummary  *types.SpecSummary
		setupSummary *types.SetupSummary
	)

	BeforeEach(func() {
		codebase = "github.com/teddyking/snowflake/reporter"
		commit = "aabb1122"
		suiteClient = new(reporterfakes.FakeSuiteClient)

		ginkgoConfig = config.GinkgoConfigType{}
		suiteSummary = &types.SuiteSummary{
			SuiteDescription: "A Sweet Suite",
		}

		r = snowflake.NewReporter(codebase, commit, suiteClient)
	})

	Describe("SpecSuiteWillBegin", func() {
		BeforeEach(func() {
			r.SpecSuiteWillBegin(ginkgoConfig, suiteSummary)
		})

		It("records the suite name", func() {
			Expect(r.Summary.Name).To(Equal("A Sweet Suite"))
		})

		It("records the time at which the suite started", func() {
			Expect(r.Summary.StartedAt).To(BeNumerically(">", 0))
		})
	})

	Describe("SpecWillRun", func() {
		BeforeEach(func() {
			specSummary = &types.SpecSummary{
				ComponentCodeLocations: []types.CodeLocation{
					types.CodeLocation{
						FileName:   "/some/file/path.go",
						LineNumber: 10,
					},
					types.CodeLocation{
						FileName:   "/some/file/path.go",
						LineNumber: 15,
					},
					types.CodeLocation{
						FileName:   "/some/file/path.go",
						LineNumber: 11,
					},
				},
			}
			r.SpecWillRun(specSummary)
		})

		It("adds the test to summary", func() {
			Expect(len(r.Summary.Tests)).To(Equal(1))
		})

		It("records the time at which the spec started", func() {
			Expect(r.Summary.Tests[0].StartedAt).To(BeNumerically(">", 0))
		})

		It("records the location of the test in the codebase", func() {
			Expect(r.Summary.Tests[0].Location).To(Equal("/some/file/path.go:15"))
		})
	})

	Describe("SpecDidComplete", func() {
		var (
			specState types.SpecState
			failure   types.SpecFailure
		)

		BeforeEach(func() {
			specState = types.SpecStatePassed
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

			r.SpecWillRun(specSummary)
			r.SpecDidComplete(specSummary)
		})

		It("doesn't append new tests to the summary", func() {
			Expect(len(r.Summary.Tests)).To(Equal(1))
		})

		It("records the spec name", func() {
			recordedSpecName := r.Summary.Tests[0].Name

			Expect(recordedSpecName).To(Equal("Integration CLI when passed an invalid flag exits with a status of 1"))
		})

		It("records the time at which the spec completed", func() {
			recordedFinishedAt := r.Summary.Tests[0].FinishedAt

			Expect(recordedFinishedAt).To(BeNumerically(">", 0))
		})

		Context("when the spec passed", func() {
			BeforeEach(func() {
				specState = types.SpecStatePassed
			})

			It("records the spec state", func() {
				recordedSpecState := r.Summary.Tests[0].State

				Expect(recordedSpecState).To(Equal(api.Test_PASSED))
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
				recordedSpecState := r.Summary.Tests[0].State

				Expect(recordedSpecState).To(Equal(api.Test_FAILED))
			})

			It("records the failure message", func() {
				recordedSpecFailureMessage := r.Summary.Tests[0].Failure.Message

				Expect(recordedSpecFailureMessage).To(Equal("Expected x to equal y\ngithub.com/teddyking/fail/fail_test.go:100"))
			})
		})

		Context("when the spec skipped", func() {
			BeforeEach(func() {
				specState = types.SpecStateSkipped
			})

			It("records the spec state", func() {
				recordedSpecState := r.Summary.Tests[0].State

				Expect(recordedSpecState).To(Equal(api.Test_SKIPPED))
			})
		})

		Context("when the spec pending", func() {
			BeforeEach(func() {
				specState = types.SpecStatePending
			})

			It("records the spec state", func() {
				recordedSpecState := r.Summary.Tests[0].State

				Expect(recordedSpecState).To(Equal(api.Test_PENDING))
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
				recordedSpecState := r.Summary.Tests[0].State

				Expect(recordedSpecState).To(Equal(api.Test_PANICKED))
			})

			It("records the failure message", func() {
				recordedSpecFailureMessage := r.Summary.Tests[0].Failure.Message

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
				recordedSpecState := r.Summary.Tests[0].State

				Expect(recordedSpecState).To(Equal(api.Test_TIMEDOUT))
			})

			It("records the failure message", func() {
				recordedSpecFailureMessage := r.Summary.Tests[0].Failure.Message

				Expect(recordedSpecFailureMessage).To(Equal("Timed out\ngithub.com/teddyking/fail/fail_test.go:100"))
			})
		})

		Context("when the spec invalid", func() {
			BeforeEach(func() {
				specState = types.SpecStateInvalid
			})

			It("records the spec state", func() {
				recordedSpecState := r.Summary.Tests[0].State

				Expect(recordedSpecState).To(Equal(api.Test_INVALID))
			})
		})
	})

	Describe("SpecSuiteDidEnd", func() {
		BeforeEach(func() {
			r.SpecSuiteWillBegin(ginkgoConfig, suiteSummary)
			r.SpecSuiteDidEnd(suiteSummary)
		})

		It("records the time at which the suite finished", func() {
			Expect(r.Summary.FinishedAt).To(BeNumerically(">", 0))
		})

		It("sends the summary to a snowflake server", func() {
			Expect(suiteClient.CreateCallCount()).To(Equal(1))
			_, sentCreateRequest, _ := suiteClient.CreateArgsForCall(0)
			Expect(sentCreateRequest.Summary.Name).To(Equal("A Sweet Suite"))
		})
	})

	Describe("BeforeSuiteDidRun", func() {
		It("is implemented but does nothing", func() {
			r.BeforeSuiteDidRun(setupSummary)
		})
	})

	Describe("AfterSuiteDidRun", func() {
		It("is implemented but does nothing", func() {
			r.AfterSuiteDidRun(setupSummary)
		})
	})
})

// TODO: make StartedAt/FinishedAt tests not suck
