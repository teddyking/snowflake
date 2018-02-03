package snowflake

import (
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
