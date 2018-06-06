package handler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/web/handler"

	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
)

var _ = Describe("Handler", func() {
	var (
		staticDirPath    string
		responseRecorder *httptest.ResponseRecorder
		h                http.Handler
	)

	BeforeEach(func() {
		staticDirPath = filepath.Join("..", "static")
		responseRecorder = httptest.NewRecorder()
		h = New(staticDirPath)
	})

	It("handles GET /static/*", func() {
		h.ServeHTTP(responseRecorder, newRequest("GET", "/static/images/favicon.ico"))

		Expect(responseRecorder.Result().StatusCode).To(Equal(http.StatusOK))
	})
})

func newRequest(method, path string) *http.Request {
	u, err := url.Parse(path)
	Expect(err).NotTo(HaveOccurred())

	return &http.Request{Method: method, URL: u}
}
