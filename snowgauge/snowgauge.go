package snowgauge

import (
	"fmt"

	"github.com/teddyking/snowflake/api"
)

// Flakes returns flaky tests, where a flaky test is defined as one that has
// both passed and failed on the same codebase and same git commit.
// Flakes must only be passed SuiteSummaries for a single codebase and commit.
// The "algorithm" used here is basic AF right now.
func Flakes(summaries []*api.SuiteSummary) ([]*api.Test, error) {
	var tests []*api.Test
	var failures []*api.Test

	if len(summaries) < 2 {
		return []*api.Test{}, nil
	}

	expectedCodebase := summaries[0].Codebase
	expectedCommit := summaries[0].Commit

	for _, summary := range summaries {
		if err := checkValidCodebase(expectedCodebase, summary.Codebase); err != nil {
			return []*api.Test{}, err
		}

		if err := checkValidCommit(expectedCommit, summary.Commit); err != nil {
			return []*api.Test{}, err
		}

		for _, test := range summary.Tests {
			tests = append(tests, test)

			if test.State == api.Test_FAILED {
				failures = append(failures, test)
			}
		}
	}

	if len(failures) < 1 {
		return []*api.Test{}, nil
	}

	return flakesFrom(tests, failures), nil
}

func checkValidCodebase(expectedCodebase, codebase string) error {
	if expectedCodebase != codebase {
		return fmt.Errorf("cannot detect flakes across different codebases - '%s' and '%s'", expectedCodebase, codebase)
	}

	return nil
}

func checkValidCommit(expectedCommit, commit string) error {
	if expectedCommit != commit {
		return fmt.Errorf("cannot detect flakes across different commits - '%s' and '%s'", expectedCommit, commit)
	}

	return nil
}

func flakesFrom(tests, failures []*api.Test) []*api.Test {
	var flakes []*api.Test

	for _, failedTest := range failures {
		for _, test := range tests {
			if failedTest.Location == test.Location {
				if test.State == api.Test_PASSED {
					flakes = append(flakes, test)
				}
			}
		}
	}

	return flakes
}
