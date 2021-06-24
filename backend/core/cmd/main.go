package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"user_ms/backend/core/internal/api/grpc"
	"user_ms/backend/core/internal/api/http"
)

func main() {

	signChan := make(chan os.Signal, 1)
	go grpc.StartGRPCServer()
	go http.StartGatewayServer()
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
	<-signChan
	log.Println("Shutting down")

}
