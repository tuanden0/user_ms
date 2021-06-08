package http

import (
	"context"
	"log"
	"net/http"
	"user_ms/backend/core/api"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc"
)

const (
	grpcServer = "0.0.0.0:50001"
	addrStr    = ":8000"
)

func StartGatewayServer() {

	// Connect to GRPC server
	conn, err := grpc.DialContext(
		context.Background(), grpcServer,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Unable to connect to GRPC server: ", err.Error())
	}

	// Create MUX
	mux := runtime.NewServeMux()

	// Create UserService Handler
	err = api.RegisterUserAPIHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register UserService gateway: ", err.Error())
	}

	// Create UserAuthService Handler
	err = api.RegisterAuthenServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register UserAuthService gateway: ", err.Error())
	}

	// Create HTTP Server
	gateway := &http.Server{
		Addr:    addrStr,
		Handler: mux,
	}

	log.Printf("Start GRPC HTTP Gateway Server on %s\n", addrStr)
	err = gateway.ListenAndServe()
	if err != nil {
		log.Fatalln("Unable to start gateway: ", err.Error())
	}

}
