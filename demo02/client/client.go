package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_learning/demo02/services"
	"log"
)

func main() {
	conn, err := grpc.Dial(":9955", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	prodClient := services.NewProdServiceClient(conn)

	prodRes, err := prodClient.GetProdInfo(context.Background(), &services.ProdRequest{ProdId: 12})
	fmt.Println(prodRes)
	fmt.Println("prod_name: ", prodRes.ProdName)
}
