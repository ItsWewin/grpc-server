package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"mayihahah.com/grpc/helper"
	"mayihahah.com/grpc/services"
	"net/http"
)

func main() {
	creds, err := helper.GetClientCreds()
	if err != nil {
		log.Fatal(err)
	}

	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	err = services.RegisterProdServiceHandlerFromEndpoint(context.Background(),
		gwmux, "localhost:8081", opt)
	if err != nil {
		log.Fatal(err)
	}
	
	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           gwmux,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
