package repository

import (
	"context"
	"errors"
	"maps"
	"slices"
	"sync"

	"github.com/ultrabor/grpc-orders/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) (*model.Order, error)
	GetByID(ctx context.Context, id int64) (*model.Order, error)
	List(ctx context.Context) ([]*model.Order, error)
}

type InMemoryOrderRepo struct {
	mu     sync.Mutex
	lastID int64
	orders map[int64]*model.Order
}

func NewInMemoryOrderRepository() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders: make(map[int64]*model.Order),
	}
}

func (r *InMemoryOrderRepo) Create(ctx context.Context, order *model.Order) (*model.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.lastID++
	order.ID = r.lastID
	r.orders[r.lastID] = order
	return order, nil
}

func (r *InMemoryOrderRepo) GetByID(ctx context.Context, id int64) (*model.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	order, ok := r.orders[id]
	if !ok {
		return nil, errors.New("element doesnt exist")
	}

	return order, nil
}

func (r *InMemoryOrderRepo) List(ctx context.Context) ([]*model.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	orders := slices.Collect(maps.Values(r.orders))

	return orders, nil
}
