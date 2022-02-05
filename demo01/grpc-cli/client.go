package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_learning/demo01/services"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient := services.NewProdServiceClient(conn)
	ctx := context.Background()

	// 单值
	//prodRes, err := prodClient.GetProdStock(context.Background(),
	//	&services.ProdRequest{ProdId: 12})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(prodRes.ProdStock)

	// 切片
	//response, err := prodClient.GetProdStocks(ctx,
	//	&services.QuerySize{Size: 10})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(response.Prodres)

	// 枚举
	prodRes, err := prodClient.GetProdStock(ctx,
		&services.ProdRequest{ProdId: 12, ProdArea: services.ProdAreas_B})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)

}
