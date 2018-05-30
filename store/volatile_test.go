package store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/store"

	"github.com/teddyking/snowflake/api"
)

var _ = Describe("Volatile", func() {
	var (
		volatileStore *VolatileStore
		summary       *api.SuiteSummary
	)

	BeforeEach(func() {
		summary = &api.SuiteSummary{
			Name:     "test-name",
			Codebase: "test-codebase",
			Commit:   "test-commit",
			Tests: []*api.Test{
				&api.Test{Name: "test-name1", Location: "test-location1"},
				&api.Test{Name: "test-name2", Location: "test-location2"},
				&api.Test{Name: "test-name3", Location: "test-location3"},
			},
		}

		volatileStore = NewVolatileStore()
	})

	Describe("Create and List", func() {
		It("stores and retrieves summaries in/from memory", func() {
			Expect(volatileStore.Create(summary)).To(Succeed())
			Expect(volatileStore.List()).To(HaveLen(1))

			summaries, err := volatileStore.List()
			Expect(err).NotTo(HaveOccurred())

			Expect(summaries[0].Name).To(Equal("test-name"))
		})
	})

	Describe("Get", func() {
		BeforeEach(func() {
			Expect(volatileStore.Create(summary)).To(Succeed())
		})

		It("retrieves a specific test from memory", func() {
			test, err := volatileStore.Get("test-codebase", "test-commit", "test-location2")
			Expect(err).NotTo(HaveOccurred())

			Expect(test.Name).To(Equal("test-name2"))
		})
	})
})
