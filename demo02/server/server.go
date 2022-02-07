package main

import (
	"google.golang.org/grpc"
	"grpc_learning/demo02/services"
	"log"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	lis, err := net.Listen("tcp", ":9955")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("start server")
	err = rpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
