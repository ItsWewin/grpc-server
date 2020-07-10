package main

import (
	"google.golang.org/grpc"
	"log"
	"mayihahah.com/grpc/helper"
	"mayihahah.com/grpc/services"
	"net"
)

func main() {
	creds, err := helper.GetServerCreds()
	if err != nil {
		log.Fatal(err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdServer))

	services.RegisterOrderServiceServer(rpcServer, new(services.OrderServer))

	services.RegisterUserServiceServer(rpcServer, new(services.UserServer))

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("server listen failed: %s", err)
	}
	rpcServer.Serve(lis)
}