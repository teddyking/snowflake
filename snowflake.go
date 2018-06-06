package snowflake

import (
	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/reporter"
)

func NewReporter(importPath, commit string, client reporter.ReporterService) *reporter.SnowflakeReporter {
	return &reporter.SnowflakeReporter{
		Report: &api.Report{
			ImportPath: importPath,
			Commit:     commit,
		},
		ReporterService: client,
	}
}
