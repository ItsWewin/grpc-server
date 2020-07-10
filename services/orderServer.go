package services

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

type OrderServer struct {
}

func (o *OrderServer) NewOrder(ctx context.Context, request *OrderRequest) (*OrderResponse, error) {
	fmt.Println(request.OrderMain)

	if err := request.OrderMain.Validate(); err != nil {
		return &OrderResponse{Status: "1", Message: err.Error()}, nil
	}

	return &OrderResponse{
		Status:        "0",
		Message:       "create order succceed",
	}, nil
}

func (o *OrderServer) GetOrder(ctx context.Context, req *OrderQuery ) (*OrderMain, error) {
	fmt.Println(req)

	t := timestamp.Timestamp{Seconds: time.Now().Unix()}
	return &OrderMain{
		OrderId:       req.OrderId,
		OrderNo:       fmt.Sprintf("order-no-%d", req.OrderId),
		UserId:        111,
		OrderMoney:    123.456,
		OrderTime:     &t,
	}, nil
}
