package interceptors

import (
	"context"
	"log"
	"user_ms/backend/core/internal/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// https://dev.to/techschoolguru/use-grpc-interceptor-for-authorization-with-jwt-1c5h

type AuthInterceptor struct {
	jwtManager      *repository.JWTManager
	accessibleRoles map[string][]string
}

func NewAuthInterceptor(jwtManager *repository.JWTManager, accessibleRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, accessibleRoles}
}

func (interceptor *AuthInterceptor) JWTUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		return handler(ctx, req)
	}
}

func AccessibleRoles() map[string][]string {
	const userServicePath = "/api.UserAPI/"

	return map[string][]string{
		userServicePath + "Create":   {"admin"},
		userServicePath + "Retrieve": {"user", "admin"},
		userServicePath + "Update":   {"user", "admin"},
		userServicePath + "Delete":   {"user", "admin"},
		userServicePath + "List":     {"admin"},
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {

	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		// everyone can access
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, role := range accessibleRoles {
		if role == claims.Role {
			return nil
		}
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
