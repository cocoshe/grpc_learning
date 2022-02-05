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
	var stock int32 = 0
	if req.ProdArea == ProdAreas_A {
		stock = 30
	} else if req.ProdArea == ProdAreas_B {
		stock = 31
	} else {
		stock = 50
	}
	return &ProdResponse{ProdStock: stock}, nil
}

func (this *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	Prodres := []*ProdResponse{
		&ProdResponse{ProdStock: 28},
		&ProdResponse{ProdStock: 29},
		&ProdResponse{ProdStock: 30},
		&ProdResponse{ProdStock: 31},
	}
	return &ProdResponseList{
		Prodres: Prodres,
	}, nil
}
