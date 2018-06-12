package data

import "github.com/teddyking/snowflake/api"

var ReportsWithAFlake = []*api.Report{
	&api.Report{
		ImportPath:  "github.com/teddyking/snowflake/examples/examplesuite",
		Commit:      "6e121a2c762a778907c94df7e774cb014531da8d",
		Description: "Examplesuite Suite",
		Tests: []*api.Test{
			&api.Test{
				Description: "Example demonstrates usage of the snowflake reporter",
				Location:    "/Users/eking/go/src/github.com/teddyking/snowflake/examples/examplesuite/example_test.go:17",
				State:       api.Test_PASSED,
				StartedAt:   1528834316,
			},
			&api.Test{
				Description: "Example when it works again demonstrates usage of the snowflake reporter",
				Location:    "/Users/eking/go/src/github.com/teddyking/snowflake/examples/examplesuite/example_test.go:22",
				State:       api.Test_PASSED,
				StartedAt:   1528834316,
			},
			&api.Test{
				Description: "Example is a flaky test",
				Location:    "/Users/eking/go/src/github.com/teddyking/snowflake/examples/examplesuite/example_test.go:31",
				State:       api.Test_FAILED,
				Failure:     &api.Failure{Message: FailureMessage},
				StartedAt:   1528834316,
			},
		},
	},
	&api.Report{
		ImportPath:  "github.com/teddyking/snowflake/examples/examplesuite",
		Commit:      "6e121a2c762a778907c94df7e774cb014531da8d",
		Description: "Examplesuite Suite",
		Tests: []*api.Test{
			&api.Test{
				Description: "Example demonstrates usage of the snowflake reporter",
				Location:    "/Users/eking/go/src/github.com/teddyking/snowflake/examples/examplesuite/example_test.go:17",
				State:       api.Test_PASSED,
				StartedAt:   1528834318,
			},
			&api.Test{
				Description: "Example when it works again demonstrates usage of the snowflake reporter",
				Location:    "/Users/eking/go/src/github.com/teddyking/snowflake/examples/examplesuite/example_test.go:22",
				State:       api.Test_PASSED,
				StartedAt:   1528834318,
			},
			&api.Test{
				Description: "Example is a flaky test",
				Location:    "/Users/eking/go/src/github.com/teddyking/snowflake/examples/examplesuite/example_test.go:31",
				State:       api.Test_PASSED,
				StartedAt:   1528834316,
			},
		},
	},
	&api.Report{
		ImportPath:  "github.com/teddyking/snowflake/examples/examplesuite",
		Commit:      "6e121a2c762a778907c94df7e774cb014531da8d",
		Description: "Examplesuite Suite",
		Tests: []*api.Test{
			&api.Test{
				Description: "Example demonstrates usage of the snowflake reporter",
				Location:    "/Users/eking/go/src/github.com/teddyking/snowflake/examples/examplesuite/example_test.go:17",
				State:       api.Test_PASSED,
				StartedAt:   1528834319,
			},
			&api.Test{
				Description: "Example when it works again demonstrates usage of the snowflake reporter",
				Location:    "/Users/eking/go/src/github.com/teddyking/snowflake/examples/examplesuite/example_test.go:22",
				State:       api.Test_PASSED,
				StartedAt:   1528834319,
			},
			&api.Test{
				Description: "Example is a flaky test",
				Location:    "/Users/eking/go/src/github.com/teddyking/snowflake/examples/examplesuite/example_test.go:31",
				State:       api.Test_PASSED,
				StartedAt:   1528834319,
			},
		},
	},
}

var FailureMessage = "Expected\n    <string>: omg\n    \nto equal\n    <string>: notflake\n/Users/eking/go/src/github.com/teddyking/snowflake/examples/examplesuite/example_test.go:35"
