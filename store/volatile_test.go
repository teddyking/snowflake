package store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/store"

	"github.com/teddyking/snowflake/api"
)

var _ = Describe("Volatile", func() {
	Describe("Create and List", func() {
		It("stores and retrieves summaries in/from memory", func() {
			volatileStore := NewVolatileStore()
			summary := &api.SuiteSummary{Name: "cake"}

			Expect(volatileStore.Create(summary)).To(Succeed())
			Expect(volatileStore.List()).To(HaveLen(1))

			summaries, err := volatileStore.List()
			Expect(err).NotTo(HaveOccurred())

			Expect(summaries[0].Name).To(Equal("cake"))
		})
	})
})
