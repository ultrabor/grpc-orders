package handler

import (
	"context"

	pb "github.com/ultrabor/grpc-orders/proto/orderpb"
)

func (h *OrderHandler) CreateOrder(
	ctx context.Context,
	req *pb.CreateOrderRequest,
) (*pb.OrderResponse, error) {

	o, err := h.service.CreateOrder(ctx, req.UserId, req.Amount)

	if err != nil {
		return nil, err
	}

	order := &pb.Order{
		Id:     o.ID,
		UserId: o.UserID,
		Amount: o.Amount,
		Status: o.Status,
	}

	return &pb.OrderResponse{
		Order: order,
	}, nil
}
