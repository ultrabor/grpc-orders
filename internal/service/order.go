package service

import (
	"context"
	"errors"

	"github.com/ultrabor/grpc-orders/internal/model"
	"github.com/ultrabor/grpc-orders/internal/repository"
)

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, user_id string, amount float64) (*model.Order, error) {
	if user_id == "" {
		return nil, errors.New("user_id cannot be empty")
	}
	if amount <= 0 {
		return nil, errors.New("amount is >= 0")
	}

	order := &model.Order{UserID: user_id, Amount: amount, Status: "Created"}

	create, err := s.repo.Create(ctx, order)

	if err != nil {
		return nil, err
	}

	return create, nil
}

func (s *OrderService) GetOrder(ctx context.Context, id int64) (*model.Order, error) {
	if id < 0 {
		return nil, errors.New("id cannot be < 0")
	}

	return s.repo.GetByID(ctx, id)
}

func (s *OrderService) ListOrders(ctx context.Context) ([]*model.Order, error) {
	return s.repo.List(ctx)
}
