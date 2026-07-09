package repository

import (
	"context"
	"errors"
	"order-service/internal/domain"
	"sync"
	"time"
)

type OrderRepository struct {
	mu     sync.Mutex
	orders map[string]*domain.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]*domain.Order),
	}
}

func (repo *OrderRepository) Create(ctx context.Context, order *domain.Order) error {

	if err := ctx.Err(); err != nil {
		return err
	}
	repo.mu.Lock()
	defer repo.mu.Unlock()

	now := time.Now().UTC()
	order.CreatedAt = now
	order.UpdatedAt = now

	stored := *order
	repo.orders[order.ID] = &stored
	return nil
}

func (repo *OrderRepository) GetAll(ctx context.Context) ([]*domain.Order, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	orders := make([]*domain.Order, 0, len(repo.orders))
	for _, order := range repo.orders {
		copied := *order
		orders = append(orders, &copied)
	}
	return orders, nil
}

func (repo *OrderRepository) GetById(ctx context.Context, ID string) (*domain.Order, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	order, ok := repo.orders[ID]
	if !ok {
		return nil, errors.New("Nao existe pedido com esse ID")
	}

	copied := *order
	return &copied, nil

}
