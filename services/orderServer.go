package services

import (
	"context"
)

type OrderServer struct {
}

func (o *OrderServer) NewOrder(ctx context.Context, order *OrderMain) (*OrderResponse, error) {
	return &OrderResponse{
		Status:        "0",
		Message:       "create order succceed",
	}, nil
}
