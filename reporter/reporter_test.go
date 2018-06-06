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
		r               *SnowflakeReporter
		reporterService *reporterfakes.FakeReporterService
		importPath      string
		commit          string
		suiteSummary    *types.SuiteSummary
		ginkgoConfig    config.GinkgoConfigType
		specSummary     *types.SpecSummary
		setupSummary    *types.SetupSummary
	)

	BeforeEach(func() {
		importPath = "github.com/teddyking/snowflake/reporter"
		commit = "aabb1122"
		reporterService = new(reporterfakes.FakeReporterService)

		ginkgoConfig = config.GinkgoConfigType{}
		suiteSummary = &types.SuiteSummary{
			SuiteDescription: "A Sweet Suite",
		}

		r = snowflake.NewReporter(importPath, commit, reporterService)
	})

	Describe("SpecSuiteWillBegin", func() {
		BeforeEach(func() {
			r.SpecSuiteWillBegin(ginkgoConfig, suiteSummary)
		})

		It("records the suite description", func() {
			Expect(r.Report.Description).To(Equal("A Sweet Suite"))
		})

		It("records the time at which the suite started", func() {
			Expect(r.Report.StartedAt).To(BeNumerically(">", 0))
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

		It("adds the test to the report", func() {
			Expect(len(r.Report.Tests)).To(Equal(1))
		})

		It("records the time at which the spec started", func() {
			Expect(r.Report.Tests[0].StartedAt).To(BeNumerically(">", 0))
		})

		It("records the filepath:linenumber location of the test", func() {
			Expect(r.Report.Tests[0].Location).To(Equal("/some/file/path.go:15"))
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

		It("doesn't append new tests to the report", func() {
			Expect(len(r.Report.Tests)).To(Equal(1))
		})

		It("records the spec description", func() {
			recordedSpecDescription := r.Report.Tests[0].Description

			Expect(recordedSpecDescription).To(Equal("Integration CLI when passed an invalid flag exits with a status of 1"))
		})

		It("records the time at which the spec completed", func() {
			recordedFinishedAt := r.Report.Tests[0].FinishedAt

			Expect(recordedFinishedAt).To(BeNumerically(">", 0))
		})

		Context("when the spec passed", func() {
			BeforeEach(func() {
				specState = types.SpecStatePassed
			})

			It("records the spec state", func() {
				recordedSpecState := r.Report.Tests[0].State

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
				recordedSpecState := r.Report.Tests[0].State

				Expect(recordedSpecState).To(Equal(api.Test_FAILED))
			})

			It("records the failure message", func() {
				recordedSpecFailureMessage := r.Report.Tests[0].Failure.Message

				Expect(recordedSpecFailureMessage).To(Equal("Expected x to equal y\ngithub.com/teddyking/fail/fail_test.go:100"))
			})
		})

		Context("when the spec skipped", func() {
			BeforeEach(func() {
				specState = types.SpecStateSkipped
			})

			It("records the spec state", func() {
				recordedSpecState := r.Report.Tests[0].State

				Expect(recordedSpecState).To(Equal(api.Test_SKIPPED))
			})
		})

		Context("when the spec pending", func() {
			BeforeEach(func() {
				specState = types.SpecStatePending
			})

			It("records the spec state", func() {
				recordedSpecState := r.Report.Tests[0].State

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
				recordedSpecState := r.Report.Tests[0].State

				Expect(recordedSpecState).To(Equal(api.Test_PANICKED))
			})

			It("records the failure message", func() {
				recordedSpecFailureMessage := r.Report.Tests[0].Failure.Message

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
				recordedSpecState := r.Report.Tests[0].State

				Expect(recordedSpecState).To(Equal(api.Test_TIMEDOUT))
			})

			It("records the failure message", func() {
				recordedSpecFailureMessage := r.Report.Tests[0].Failure.Message

				Expect(recordedSpecFailureMessage).To(Equal("Timed out\ngithub.com/teddyking/fail/fail_test.go:100"))
			})
		})

		Context("when the spec invalid", func() {
			BeforeEach(func() {
				specState = types.SpecStateInvalid
			})

			It("records the spec state", func() {
				recordedSpecState := r.Report.Tests[0].State

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
			Expect(r.Report.FinishedAt).To(BeNumerically(">", 0))
		})

		It("sends the report to a snowflake server", func() {
			Expect(reporterService.CreateCallCount()).To(Equal(1))
			_, sentCreateRequest, _ := reporterService.CreateArgsForCall(0)
			Expect(sentCreateRequest.Report.Description).To(Equal("A Sweet Suite"))
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
