package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func LogUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println(info.FullMethod)
	return handler(ctx, req)
}
