package services

import (
	"context"
)

type ProdServer struct {
}

func (p *ProdServer) GetProdStock(ctx context.Context, in *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 100}, nil
}

func (p *ProdServer) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {

	Prodres := []*ProdResponse{
		{ProdStock:101},
		{ProdStock:102},
		{ProdStock:103},
		{ProdStock:103},
	}

	return &ProdResponseList{Prods:Prodres}, nil
}

func (p *ProdServer) GetProdStocksByArea(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	var stock int32
	switch req.ProdArea {
	case ProdAreas_A:
		stock = 100
	case ProdAreas_B:
		stock = 200
	case ProdAreas_C:
		stock = 300
	default:
		stock = 1000
	}
	return &ProdResponse{ProdStock: stock}, nil
}
