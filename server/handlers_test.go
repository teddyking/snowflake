package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/server"

	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/teddyking/snowflake"
	"github.com/teddyking/snowflake/server/store/storefakes"
)

var _ = Describe("Handlers", func() {
	Describe("SuitesHandler", func() {
		var (
			suitesHandler    *SuitesHandler
			responseRecorder *httptest.ResponseRecorder
			fakeStore        *storefakes.FakeStore
		)

		BeforeEach(func() {
			fakeStore = new(storefakes.FakeStore)
			suitesHandler = NewSuitesHandler(fakeStore)
			responseRecorder = httptest.NewRecorder()
		})

		Describe("List", func() {
			BeforeEach(func() {
				fakeStore.AllReturns([]snowflake.Suite{
					snowflake.Suite{
						Name: "A Sweet Suite",
						Tests: []*snowflake.Test{
							&snowflake.Test{
								Name: "A Sweet Test",
							},
						},
					},
				}, nil)
			})

			JustBeforeEach(func() {
				suitesHandler.List(responseRecorder, newTestRequest(""))
			})

			It("returns an HTTP 200", func() {
				Expect(responseRecorder.Code).To(Equal(http.StatusOK))
			})

			It("returns all stored Suites in JSON format", func() {
				var returnedSuites []snowflake.Suite
				Expect(json.Unmarshal(responseRecorder.Body.Bytes(), &returnedSuites)).To(Succeed())

				Expect(len(returnedSuites)).To(Equal(1))
				Expect(len(returnedSuites[0].Tests)).To(Equal(1))
				Expect(returnedSuites[0].Tests[0].Name).To(Equal("A Sweet Test"))
			})

			Context("when the store returns an error", func() {
				BeforeEach(func() {
					fakeStore.AllReturns([]snowflake.Suite{}, errors.New("store-error"))
				})

				It("returns an HTTP 500", func() {
					Expect(responseRecorder.Code).To(Equal(http.StatusInternalServerError))
				})
			})
		})

		Describe("Create", func() {
			var suite string

			BeforeEach(func() {
				suite = `{"name":"A Sweet Suite"}`
			})

			JustBeforeEach(func() {
				suitesHandler.Create(responseRecorder, newTestRequest(suite))
			})

			It("returns an HTTP 201", func() {
				Expect(responseRecorder.Code).To(Equal(http.StatusCreated))
			})

			It("saves the suite to the store", func() {
				Expect(fakeStore.SaveCallCount()).To(Equal(1))
				Expect(fakeStore.SaveArgsForCall(0).Name).To(Equal("A Sweet Suite"))
			})

			Context("when the request body is not valid json", func() {
				BeforeEach(func() {
					suite = `invalid-json`
				})

				It("returns an HTTP 500", func() {
					Expect(responseRecorder.Code).To(Equal(http.StatusInternalServerError))
				})
			})

			Context("when the store returns an error", func() {
				BeforeEach(func() {
					fakeStore.SaveReturns(errors.New("store-error"))
				})

				It("returns an HTTP 500", func() {
					Expect(responseRecorder.Code).To(Equal(http.StatusInternalServerError))
				})
			})
		})
	})
})

func newTestRequest(body string) *http.Request {
	reader := strings.NewReader(body)

	request, err := http.NewRequest("", "", reader)
	Expect(err).NotTo(HaveOccurred())

	return request
}
