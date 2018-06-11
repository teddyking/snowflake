package handler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/web/handler"

	"errors"
	"net/http"
	"net/http/httptest"
	"path/filepath"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/web/handler/handlerfakes"
)

var _ = Describe("IndexHandler", func() {
	var (
		fakeFlakerService *handlerfakes.FakeFlakerService
		responseRecorder  *httptest.ResponseRecorder
		templatePath      string
		h                 *IndexHandler
	)

	BeforeEach(func() {
		fakeFlakerService = new(handlerfakes.FakeFlakerService)
		responseRecorder = httptest.NewRecorder()
		templatePath = filepath.Join("..", "static", "templates")

		fakeFlakerService.ListReturns(&api.FlakerListRes{
			Flakes: []*api.Flake{
				&api.Flake{TestDescription: "test-flake"},
			},
		}, nil)
	})

	JustBeforeEach(func() {
		h = NewIndexHandler(templatePath, fakeFlakerService)
		h.HandleIndex(responseRecorder, nil)
	})

	It("returns an HTTP 200", func() {
		Expect(responseRecorder.Result().StatusCode).To(Equal(http.StatusOK))
	})

	It("retrieves flakes using the flaker service", func() {
		Expect(fakeFlakerService.ListCallCount()).To(Equal(1))
	})

	It("renders flakes into index.html via the template", func() {
		body := string(responseRecorder.Body.Bytes())

		Expect(body).To(ContainSubstring("<title>snowflake</title>"))
		Expect(body).To(ContainSubstring("test-flake"))
	})

	When("the flaker service returns an error", func() {
		BeforeEach(func() {
			fakeFlakerService.ListReturns(&api.FlakerListRes{}, errors.New("error-retrieving-flakes"))
		})

		It("returns an HTTP 500", func() {
			Expect(responseRecorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
		})
	})
})
