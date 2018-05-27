package snowflake

import (
	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/reporter"
)

func NewReporter(codebase, commit string, client reporter.SuiteClient) *reporter.SnowflakeReporter {
	return &reporter.SnowflakeReporter{
		Summary: &api.SuiteSummary{
			Codebase: codebase,
			Commit:   commit,
		},
		Client: client,
	}
}
