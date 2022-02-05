package main

import (
	"google.golang.org/grpc"
	services2 "grpc_learning/demo01/services"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	services2.RegisterProdServiceServer(rpcServer, new(services2.ProdService))

	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)

}
