package handler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/web/handler"

	"html/template"
	"time"
)

var _ = Describe("CustomTemplateFuncs", func() {
	var (
		nl2br                  func(string) template.HTML
		humanizeTime           func(int64) string
		codebaseFromImportPath func(string) string
		trimCommit             func(string) string
	)

	BeforeSuite(func() {
		var ok bool

		nl2br, ok = CustomTemplateFuncs["nl2br"].(func(string) template.HTML)
		Expect(ok).To(BeTrue())

		humanizeTime, ok = CustomTemplateFuncs["humanizeTime"].(func(int64) string)
		Expect(ok).To(BeTrue())

		codebaseFromImportPath, ok = CustomTemplateFuncs["codebaseFromImportPath"].(func(string) string)
		Expect(ok).To(BeTrue())

		trimCommit, ok = CustomTemplateFuncs["trimCommit"].(func(string) string)
		Expect(ok).To(BeTrue())
	})

	DescribeTable("nl2br",
		func(input, expectedOutput string) {
			output := string(nl2br(input))
			Expect(expectedOutput).To(Equal(output))
		},
		Entry("simple usage", "Hello\nthere", "Hello<br />there"),
		Entry("multiple newline usage", "Hello\n\nthere", "Hello<br /><br />there"),
		Entry("safety usage 1", "<", "&lt;"),
		Entry("safety usage 2", ">", "&gt;"),
		Entry("safety usage 3", "<script>alert('hack');</script>", "&lt;script&gt;alert(&#39;hack&#39;);&lt;/script&gt;"),
	)

	DescribeTable("humanizeTime",
		func(input int64, expectedOutput string) {
			output := humanizeTime(input)
			Expect(expectedOutput).To(Equal(output))
		},
		Entry("time", int64(1528789164), time.Unix(1528789164, 0).Format(time.RFC1123)),
		Entry("time", int64(3376728000), time.Unix(3376728000, 0).Format(time.RFC1123)),
	)

	DescribeTable("codebaseFromImportPath",
		func(input string, expectedOutput string) {
			output := codebaseFromImportPath(input)
			Expect(expectedOutput).To(Equal(output))
		},
		Entry("suite in toplevel package", "github.com/teddyking/snowflake", "github.com/teddyking/snowflake"),
		Entry("suite in toplevel package with ending /", "github.com/teddyking/snowflake/", "github.com/teddyking/snowflake"),
		Entry("suite in a subpackage", "github.com/teddyking/snowflake/example/examplesuite", "github.com/teddyking/snowflake"),
		Entry("suite in a subpackage", "github.com/teddyking/snowflake/example/examplesuite/", "github.com/teddyking/snowflake"),
		Entry("bad import path", "github.com/teddyking", "github.com/teddyking"),
		Entry("bad import path", "github.com", "github.com"),
	)

	DescribeTable("trimCommit",
		func(input string, expectedOutput string) {
			output := trimCommit(input)
			Expect(expectedOutput).To(Equal(output))
		},
		Entry("long commit", "6e121a2c762a778907c94df7e774cb014531da8d", "6e121a2c"),
		Entry("short commit", "6", "6"),
		Entry("empty commit", "", ""),
	)
})
