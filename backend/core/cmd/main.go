package main

import (
	"user_ms/backend/core/internal/api/grpc"
	"user_ms/backend/core/internal/api/http"
)

func main() {
	go grpc.StartGRPCServer()
	http.StartGatewayServer()
}
