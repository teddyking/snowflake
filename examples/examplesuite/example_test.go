package examplesuite_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/examples/examplesuite"

	"io/ioutil"
	"path/filepath"
)

var _ = Describe("Example", func() {
	testCake := func() {
		Expect(TheCake()).To(Equal("a lie"))
	}

	It("demonstrates usage of the snowflake reporter", func() {
		testCake()
	})

	Context("when it works", func() {
		It("again demonstrates usage of the snowflake reporter", func() {
			testCake()
		})

		It("again again demonstrates usage of the snowflake reporter", func() {
			testCake()
		})
	})

	It("is a flaky test", func() {
		flakeFileContents, err := ioutil.ReadFile(filepath.Join("assets", "flake.txt"))
		Expect(err).NotTo(HaveOccurred())

		Expect(string(flakeFileContents)).To(Equal("notflake"))
	})
})
