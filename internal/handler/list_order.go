package handler

import pb "github.com/ultrabor/grpc-orders/proto/orderpb"

func (h *OrderHandler) ListOrders(
	req *pb.Empty,
	stream pb.OrderService_ListOrdersServer,
) error {

	o, err := h.service.ListOrders(stream.Context())

	if err != nil {
		return err
	}

	for i := 0; i < len(o); i++ {
		err := stream.Send(&pb.OrderResponse{
			Order: &pb.Order{
				Id:     o[i].ID,
				UserId: o[i].UserID,
				Amount: o[i].Amount,
				Status: o[i].Status,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
