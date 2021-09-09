package interceptors

import (
	"context"
	"log"
	"user_ms/backend/core/internal/repository"

	"google.golang.org/grpc"
)

func LogUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	userClaim := repository.ParseUserOrAnonymousFromCTX(ctx)
	log.Printf("%v has call %v", userClaim.Username, info.FullMethod)
	return handler(ctx, req)
}
