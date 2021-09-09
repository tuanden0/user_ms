package interceptors

import (
	"context"
	"log"
	"user_ms/backend/core/internal/constant"
	"user_ms/backend/core/internal/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// https://dev.to/techschoolguru/use-grpc-interceptor-for-authorization-with-jwt-1c5h

type AuthInterceptor struct {
	jwtManager *repository.JWTManager
	// accessibleRoles map[string][]string
}

func NewAuthInterceptor(jwtManager *repository.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{jwtManager}
}

func (interceptor *AuthInterceptor) JWTUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		newCTX, err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		return handler(newCTX, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (context.Context, error) {

	// Bypass login
	loginPath := "/api.AuthenService/Login"
	if method == loginPath {
		return ctx, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	newCTX := context.WithValue(ctx, constant.JWTKey, repository.UserClaims{
		Id:       claims.GetID(),
		Username: claims.GetUserName(),
		Role:     claims.GetRole(),
	})

	return newCTX, nil
}
