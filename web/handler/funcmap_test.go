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
		nl2br        func(string) template.HTML
		humanizeTime func(int64) string
	)

	BeforeSuite(func() {
		var ok bool

		nl2br, ok = CustomTemplateFuncs["nl2br"].(func(string) template.HTML)
		Expect(ok).To(BeTrue())

		humanizeTime, ok = CustomTemplateFuncs["humanizeTime"].(func(int64) string)
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
})
