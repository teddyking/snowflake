package reporter

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
	"github.com/teddyking/snowflake/api"
	"google.golang.org/grpc"
)

//go:generate counterfeiter . SuiteClient
type SuiteClient interface {
	Create(ctx context.Context, in *api.CreateRequest, opts ...grpc.CallOption) (*api.CreateResponse, error)
}

type SnowflakeReporter struct {
	Summary *api.SuiteSummary
	Client  SuiteClient
}

func (r *SnowflakeReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
	r.Summary.Name = summary.SuiteDescription
	r.Summary.StartedAt = time.Now().Unix()
}

func (r *SnowflakeReporter) SpecWillRun(specSummary *types.SpecSummary) {
	test := &api.Test{
		StartedAt: time.Now().Unix(),
		Location:  findIt(specSummary.ComponentCodeLocations),
	}

	r.Summary.Tests = append(r.Summary.Tests, test)
}

func (r *SnowflakeReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	test := r.findTestByLocation(findIt(specSummary.ComponentCodeLocations))

	test.Name = strings.Join(specSummary.ComponentTexts[1:], " ")
	test.FinishedAt = time.Now().Unix()
	test.State = ginkgoStateToTestState(specSummary.State)

	if specSummary.State == types.SpecStateFailed ||
		specSummary.State == types.SpecStatePanicked ||
		specSummary.State == types.SpecStateTimedOut {

		test.Failure = &api.Failure{
			Message: failureMessage(specSummary.Failure),
		}
	}
}

func (r *SnowflakeReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	r.Summary.FinishedAt = time.Now().Unix()

	ctx := context.Background()
	req := &api.CreateRequest{Summary: r.Summary}
	r.Client.Create(ctx, req)
}

func (r *SnowflakeReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {}
func (r *SnowflakeReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary)  {}

func failureMessage(failure types.SpecFailure) string {
	return fmt.Sprintf("%s\n%s", failure.Message, failure.Location.String())
}

func ginkgoStateToTestState(state types.SpecState) api.Test_State {
	switch state {
	case types.SpecStatePassed:
		return api.Test_PASSED
	case types.SpecStateFailed:
		return api.Test_FAILED
	case types.SpecStateSkipped:
		return api.Test_SKIPPED
	case types.SpecStatePending:
		return api.Test_PENDING
	case types.SpecStatePanicked:
		return api.Test_PANICKED
	case types.SpecStateTimedOut:
		return api.Test_TIMEDOUT
	case types.SpecStateInvalid:
		return api.Test_INVALID
	default:
		return api.Test_UNKNOWN
	}
}

// findIt will return the types.CodeLocation with the highest
// LineNumber from the provided locations as a string. In other words
// it will return "filepath:linenumber" for a test's "It(...)" statement.
// This is expected to be unique for every test.
func findIt(locations []types.CodeLocation) string {
	var highestLineNumber int
	var itCodeLocation types.CodeLocation

	for _, location := range locations {
		if location.LineNumber > highestLineNumber {
			highestLineNumber = location.LineNumber
			itCodeLocation = location
		}
	}

	return itCodeLocation.String()
}

func (r *SnowflakeReporter) findTestByLocation(location string) *api.Test {
	for _, test := range r.Summary.Tests {
		if test.Location == location {
			return test
		}
	}

	return nil
}
