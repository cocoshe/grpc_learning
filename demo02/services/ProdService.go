package services

import "context"

type ProdService struct {
}

func (this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	switch req.ProdAres {
	case ProdAreas_B:
		stock = 31
	case ProdAreas_C:
		stock = 50
	default:
		stock = 102
	}
	return &ProdResponse{ProdStock: stock}, nil
}

func (this *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdRequestList, error) {
	Prodres := []*ProdResponse{
		&ProdResponse{ProdStock: 12},
		&ProdResponse{ProdStock: 13},
		&ProdResponse{ProdStock: 14},
	}
	return &ProdRequestList{Prodres: Prodres}, nil
}

func (this *ProdService) GetProdInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error) {
	return &ProdModel{ProdId: 101, ProdName: "测试商品", ProdPrice: 20.5}, nil
}
