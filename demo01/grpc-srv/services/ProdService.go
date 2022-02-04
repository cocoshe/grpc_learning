package services

import (
	"context"
	"log"
	//"google.golang.org/grpc"
)

type ProdService struct {
}

func (this *ProdService) mustEmbedUnimplementedProdServiceServer() {
	//TODO implement me
	log.Println("implement me")
}

//服务具体实现
func (this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20}, nil
}
