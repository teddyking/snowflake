package middleware_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/middleware"

	"context"
	"errors"

	"google.golang.org/grpc"
)

var _ = Describe("Middleware", func() {
	Describe("Logging", func() {
		Describe("Server", func() {
			var fakeHandler grpc.UnaryHandler

			BeforeEach(func() {
				fakeHandler = func(ctx context.Context, req interface{}) (interface{}, error) {
					return true, nil
				}
			})

			It("invokes the provided handler", func() {
				res, err := WithServerLogging(nil, nil, &grpc.UnaryServerInfo{}, fakeHandler)
				Expect(err).NotTo(HaveOccurred())

				Expect(res).To(BeTrue())
			})
		})

		Describe("Client", func() {
			var fakeInvoker grpc.UnaryInvoker

			BeforeEach(func() {
				fakeInvoker = func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
					return errors.New("error-invoking")
				}
			})

			It("invokes the provided invoker", func() {
				err := WithClientLogging(nil, "", nil, nil, nil, fakeInvoker)
				Expect(err).To(MatchError("error-invoking"))
			})
		})
	})
})
