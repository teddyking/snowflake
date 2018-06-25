package middleware

import (
	"context"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

func WithServerLogging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.WithFields(log.Fields{"method": info.FullMethod}).Printf("request")

	return handler(ctx, req)
}

func WithClientLogging(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.WithFields(log.Fields{"method": method}).Printf("request")

	return invoker(ctx, method, req, reply, cc, opts...)
}
