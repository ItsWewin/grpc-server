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
		&ProdResponse{ProdStock:101},
		&ProdResponse{ProdStock:102},
		&ProdResponse{ProdStock:103},
		&ProdResponse{ProdStock:103},
	}

	return &ProdResponseList{Prods:Prodres}, nil
}
