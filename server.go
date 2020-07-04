package main

import (
	"google.golang.org/grpc"
	"log"
	"mayihahah.com/grpc/services"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()

	services.RegisterProdServiceServer(rpcServer, new(services.ProdServer))

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("server listen failed: %s", err)
	}

	rpcServer.Serve(lis)
}