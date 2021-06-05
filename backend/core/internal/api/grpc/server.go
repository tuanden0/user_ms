package grpc

import (
	"log"
	"net"
	"user_ms/backend/core/api"
	"user_ms/backend/core/internal/models"
	"user_ms/backend/core/internal/repository"
	"user_ms/backend/core/internal/services"

	"google.golang.org/grpc"
)

const (
	netStr  = "tcp"
	addrStr = ":50001"
)

func StartGRPCServer() {

	log.Printf("Starting server... on %s\n", addrStr)

	// Init GRPC Server
	lis, err := net.Listen(netStr, addrStr)
	if err != nil {
		log.Fatalln("Unable to start GRPC server: ", err.Error())
	}

	// Connect DB
	db := models.ConnectDatabase()

	// Create User Services
	userRepo := repository.NewUserMng(db)
	userService := services.Init(userRepo)

	// Init GRPC
	s := grpc.NewServer()
	api.RegisterUserAPIServer(s, userService)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
