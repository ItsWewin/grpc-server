package services

import (
	"context"
)

type ProdServer struct {
}

func (p *ProdServer) GetProdStock(ctx context.Context, in *ProdRequest) (*ProdResponse, error) {

	return &ProdResponse{ProdStock: 100}, nil
}
