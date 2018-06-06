package handler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/web/handler"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("IndexHandler", func() {
	var (
		responseRecorder *httptest.ResponseRecorder
		h                *IndexHandler
	)

	BeforeEach(func() {
		responseRecorder = httptest.NewRecorder()
		h = NewIndexHandler("../template/", nil)
	})

	JustBeforeEach(func() {
		h.HandleIndex(responseRecorder, nil)
	})

	It("returns an HTTP 200", func() {
		Expect(responseRecorder.Result().StatusCode).To(Equal(http.StatusOK))
	})

	It("renders the index.html template", func() {
		body := responseRecorder.Body.Bytes()
		Expect(body).To(ContainSubstring("<title>snowflake</title>"))
	})

	When("the index.html template cannot be parsed", func() {
		BeforeEach(func() {
			h = NewIndexHandler("../invalid/template/path/", nil)
		})

		It("returns an HTTP 500", func() {
			Expect(responseRecorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
		})
	})
})
