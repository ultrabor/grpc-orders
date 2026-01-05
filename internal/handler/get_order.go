package handler

import (
	"context"

	pb "github.com/ultrabor/grpc-orders/proto/orderpb"
)

func (h *OrderHandler) GetOrder(
	ctx context.Context,
	req *pb.GetOrderRequest,
) (*pb.OrderResponse, error) {

	o, err := h.service.GetOrder(ctx, req.Id)

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
