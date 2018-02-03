package snowflake_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

var _ = Describe("SnowflakeReporter", func() {
	Describe("SpecSuiteWillBegin", func() {
		var (
			r            *SnowflakeReporter
			suiteSummary *types.SuiteSummary
			ginkgoConfig config.GinkgoConfigType
		)

		BeforeEach(func() {
			r = NewReporter()

			ginkgoConfig = config.GinkgoConfigType{}
			suiteSummary = &types.SuiteSummary{
				SuiteDescription: "A Sweet Suite",
			}

			r.SpecSuiteWillBegin(ginkgoConfig, suiteSummary)
		})

		It("records the suite name", func() {
			Expect(r.Suite.Name).To(Equal("A Sweet Suite"))
		})
	})
})
