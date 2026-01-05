package handler

import (
	"github.com/ultrabor/grpc-orders/internal/service"
	pb "github.com/ultrabor/grpc-orders/proto/orderpb"
)

type OrderHandler struct {
	pb.UnimplementedOrderServiceServer
	service *service.OrderService
}

func NewOrderHandler(Service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: Service}
}
