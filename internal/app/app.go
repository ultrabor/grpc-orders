package app

import (
	"net"

	"github.com/ultrabor/grpc-orders/internal/handler"
	"github.com/ultrabor/grpc-orders/internal/interceptor"
	"github.com/ultrabor/grpc-orders/internal/repository"
	"github.com/ultrabor/grpc-orders/internal/service"
	"github.com/ultrabor/grpc-orders/pkg/logger"
	pb "github.com/ultrabor/grpc-orders/proto/orderpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start() {

	l := logger.New()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		l.Error("probrem with listing port", "err", err)
	}

	grpcServer := grpc.NewServer(
		interceptor.ChainUnary(l, interceptor.LoggingUnary(l)))

	repo := repository.NewInMemoryOrderRepository()
	s := service.NewOrderService(repo)

	pb.RegisterOrderServiceServer(grpcServer, handler.NewOrderHandler(s))

	reflection.Register(grpcServer)

	l.Info("gRPC server started on :50051")
	l.Error("error on grpc server", "err", grpcServer.Serve(lis))
}
