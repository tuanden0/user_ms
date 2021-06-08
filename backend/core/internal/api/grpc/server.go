package grpc

import (
	"log"
	"net"
	"time"
	"user_ms/backend/core/api"
	"user_ms/backend/core/internal/api/interceptors"
	"user_ms/backend/core/internal/configs"
	"user_ms/backend/core/internal/repository"
	"user_ms/backend/core/internal/services"

	"google.golang.org/grpc"
)

const (
	netStr        = "tcp"
	addrStr       = "0.0.0.0:50001"
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

func StartGRPCServer() {

	log.Printf("Starting server... on %s\n", addrStr)

	// Init GRPC Server
	lis, err := net.Listen(netStr, addrStr)
	if err != nil {
		log.Fatalln("Unable to start GRPC server: ", err.Error())
	}

	// Connect DB
	db := configs.ConnectDatabase()

	// Create User Services
	userAuthRepo := repository.NewJWTManager(secretKey, tokenDuration, db)
	userRepo := repository.NewUserMng(db)
	userService := services.NewUserService(userRepo)
	userAuthenService := services.NewUserAuthenService(userAuthRepo)

	// Init GRPC and Interceptor (middleware)
	// auth := interceptors.NewAuthInterceptor(userAuthRepo, interceptors.AccessibleRoles())
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.LogUnaryInterceptor,
			// auth.JWTUnaryInterceptor(),
		),
	)
	api.RegisterUserAPIServer(s, userService)
	api.RegisterAuthenServiceServer(s, userAuthenService)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
