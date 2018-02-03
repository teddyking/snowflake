package snowflake

import (
	"fmt"
	"strings"
	"time"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

type SnowflakeReporter struct {
	Suite Suite
}

func NewReporter() *SnowflakeReporter {
	return &SnowflakeReporter{}
}

func (r *SnowflakeReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
	r.Suite.Name = summary.SuiteDescription
}

func (r *SnowflakeReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	test := &Test{
		Name:        strings.Join(specSummary.ComponentTexts[1:], " "),
		CompletedAt: time.Now(),
		State:       stateToString(specSummary.State),
	}

	if specSummary.State == types.SpecStateFailed ||
		specSummary.State == types.SpecStatePanicked ||
		specSummary.State == types.SpecStateTimedOut {

		test.Failure = Failure{
			Message: failureMessage(specSummary.Failure),
		}
	}

	r.Suite.Tests = append(r.Suite.Tests, test)
}

func (r *SnowflakeReporter) SpecSuiteDidEnd(summary *types.SuiteSummary)        {}
func (r *SnowflakeReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {}
func (r *SnowflakeReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary)  {}
func (r *SnowflakeReporter) SpecWillRun(specSummary *types.SpecSummary)         {}

func failureMessage(failure types.SpecFailure) string {
	return fmt.Sprintf("%s\n%s", failure.Message, failure.Location.String())
}

func stateToString(state types.SpecState) string {
	switch state {
	case types.SpecStatePassed:
		return "passed"
	case types.SpecStateFailed:
		return "failed"
	case types.SpecStateSkipped:
		return "skipped"
	case types.SpecStatePending:
		return "pending"
	case types.SpecStatePanicked:
		return "panicked"
	case types.SpecStateTimedOut:
		return "timedout"
	case types.SpecStateInvalid:
		return "invalid"
	default:
		return "unknown"
	}
}
