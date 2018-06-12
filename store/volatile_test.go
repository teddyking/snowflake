package store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/store"

	"github.com/teddyking/snowflake/api"
)

var _ = Describe("Volatile", func() {
	Describe("CreateReport and List", func() {
		var (
			volatileStore *VolatileStore
			report        *api.Report
		)

		BeforeEach(func() {
			report = &api.Report{
				Description: "test-desc",
				ImportPath:  "test-import-path",
				Commit:      "test-commit",
				Tests: []*api.Test{
					&api.Test{Description: "test-name1", Location: "test-location1"},
					&api.Test{Description: "test-name2", Location: "test-location2"},
					&api.Test{Description: "test-name3", Location: "test-location3"},
				},
			}

			volatileStore = NewVolatileStore()
		})

		It("stores and retrieves reports in/from memory", func() {
			Expect(volatileStore.CreateReport(report)).To(Succeed())
			Expect(volatileStore.ListReports()).To(HaveLen(1))

			reports, err := volatileStore.ListReports()
			Expect(err).NotTo(HaveOccurred())

			Expect(reports[0].Description).To(Equal("test-desc"))
		})
	})

	Describe("NewVolatileStore", func() {
		It("returns a volatile store containing 0 reports", func() {
			volatileStore := NewVolatileStore()
			Expect(volatileStore.ListReports()).To(HaveLen(0))
		})

		Context("when passed a WithInitialReports storeOpt", func() {
			It("returns a volatile store containing the initial reports", func() {
				volatileStore := NewVolatileStore(
					WithInitialReports([]*api.Report{&api.Report{}, &api.Report{}}),
				)
				Expect(volatileStore.ListReports()).To(HaveLen(2))
			})
		})
	})
})
