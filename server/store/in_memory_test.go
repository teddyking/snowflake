package store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/server/store"

	"github.com/teddyking/snowflake"
)

var _ = Describe("InMemory", func() {
	var store Store

	BeforeEach(func() {
		store = NewInMemory()
	})

	Describe("All", func() {
		It("returns all suites", func() {
			suites, err := store.All()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(suites)).To(Equal(0))
		})
	})

	Describe("Save", func() {
		It("saves the suite to the store", func() {
			suite := snowflake.Suite{Name: "A Sweet Suite"}
			Expect(store.Save(suite)).To(Succeed())

			savedSuites, err := store.All()
			Expect(err).NotTo(HaveOccurred())

			Expect(len(savedSuites)).To(Equal(1))
			Expect(savedSuites[0].Name).To(Equal("A Sweet Suite"))
		})
	})
})
