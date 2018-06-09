package snowgauge

import (
	"github.com/teddyking/snowflake/api"
)

// Flakes searches through a bunch of reports and returns only the flaky tests.
// A flake is defined as a test in the same codebase, on the same commit,
// that has both PASSED and !PASSED states for repeated runs of the test.
// The "algorithm" used is basic AF right now.
func Flakes(reports []*api.Report) ([]*api.Flake, error) {
	var flakes []*api.Flake

	if len(reports) < 2 {
		return []*api.Flake{}, nil
	}

	importPath := reports[0].ImportPath
	commit := reports[0].Commit
	SuiteDescription := reports[0].Description

	testsByTestLocation := sortTestsByTestLocation(reports)

	for _, tests := range testsByTestLocation {
		states := make([]api.Test_State, len(tests))

		var successes, failures, startedAt int64
		for i, test := range tests {
			states[i] = test.State
			if states[i] == api.Test_PASSED {
				successes++
			} else {
				failures++
			}

			if i == 0 {
				startedAt = test.StartedAt
			}
			if startedAt > test.StartedAt {
				startedAt = test.StartedAt
			}

			if i == len(tests)-1 {
				sameStateForAllTests, failedTest := sameStates(tests)
				if !sameStateForAllTests {
					flakes = append(flakes, &api.Flake{
						ImportPath:       importPath,
						Commit:           commit,
						SuiteDescription: SuiteDescription,
						TestDescription:  failedTest.Description,
						Location:         failedTest.Location,
						Failure:          failedTest.Failure,
						Successes:        successes,
						Failures:         failures,
						StartedAt:        startedAt,
					})
				}
			}
		}
	}

	return flakes, nil
}

// sortTestByTestLocation sorts all of the tests across all provided reports
// into a map[string][]*api.Test, where the map key is the test's location.
// Test location is unique for each test.
func sortTestsByTestLocation(reports []*api.Report) map[string][]*api.Test {
	testsByLocation := make(map[string][]*api.Test)

	for _, report := range reports {
		for _, test := range report.Tests {
			_, ok := testsByLocation[test.Location]
			if !ok {
				testsByLocation[test.Location] = []*api.Test{}
			}

			testsByLocation[test.Location] = append(testsByLocation[test.Location], test)
		}
	}

	return testsByLocation
}

// sameStates returns true if all provided tests have the same state.
// If they don't, it returns false and a test with a FAILED state.
func sameStates(tests []*api.Test) (bool, *api.Test) {
	states := make([]api.Test_State, len(tests))
	var failedTest *api.Test

	for i, test := range tests {
		if test.State != api.Test_PASSED {
			failedTest = test
		}

		states[i] = test.State
	}

	return same(states), failedTest
}

// same returns true if all provided states are the same.
func same(states []api.Test_State) bool {
	for i := 1; i < len(states); i++ {
		if states[i] != states[0] {
			return false
		}
	}
	return true
}
