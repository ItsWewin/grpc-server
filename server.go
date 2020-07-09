package main

import (
	"google.golang.org/grpc"
	"log"
	"mayihahah.com/grpc/helper"
	"mayihahah.com/grpc/services"
	"net"
)

func main() {
	//creds, err := credentials.NewServerTLSFromFile("certificate/backend.cert", "certificate/backend.key")
	//if err != nil {
	//	log.Fatal(err)
	//}
	
	//cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//certPool := x509.NewCertPool()
	//
	//ca, err := ioutil.ReadFile("cert/ca.pem")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//certPool.AppendCertsFromPEM(ca)
	//
	//creds := credentials.NewTLS(&tls.Config{
	//	Certificates:                []tls.Certificate{cert},
	//	ClientAuth:                  tls.RequireAndVerifyClientCert,
	//	ClientCAs:                   certPool,
	//})

	creds, err := helper.GetServerCreds()
	if err != nil {
		log.Fatal(err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdServer))

	services.RegisterOrderServiceServer(rpcServer, new(services.OrderServer))

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("server listen failed: %s", err)
	}
	rpcServer.Serve(lis)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(r)
	//	rpcServer.ServeHTTP(w, r)
	//})
	//
	//httpServer := &http.Server{
	//	Addr:              ":8081",
	//	Handler:           mux,
	//}

	//err = httpServer.ListenAndServe()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// httpServer.ListenAndServeTLS("certificate/backend.cert", "certificate/backend.key")
}