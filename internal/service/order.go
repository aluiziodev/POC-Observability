package service

import (
	"context"
	"errors"
	"order-service/internal/connection"
	"order-service/internal/domain"
	"order-service/internal/models"
	"order-service/internal/observability"
	"order-service/internal/repository"

	"github.com/google/uuid"
)

func CreateOrder(ctx context.Context, input models.CreateOrderInput) (*domain.Order, error) {
	db, err := connection.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	input.Validate()

	order := &domain.Order{
		ID:       uuid.NewString(),
		Item:     input.Item,
		Quantity: input.Quantity,
		Status:   domain.OrderStatusPending,
	}

	orderRepo := repository.NewOrderRepository(db)
	if err := orderRepo.Create(ctx, order); err != nil {
		return nil, err
	}

	observability.OrdersCreatedTotal.Inc()
	return order, nil
}

func ListOrders(ctx context.Context) ([]*domain.Order, error) {
	db, err := connection.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	orderRepo := repository.NewOrderRepository(db)
	orders, err := orderRepo.GetAll(ctx)
	return orders, err
}

func GetOrder(ctx context.Context, ID string) (*domain.Order, error) {
	db, err := connection.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	orderRepo := repository.NewOrderRepository(db)
	order, err := orderRepo.GetById(ctx, ID)
	return order, err
}

func UpdateOrderStatus(ctx context.Context, ID string, status domain.OrderStatus) (*domain.Order, error) {
	db, err := connection.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if !status.IsValid() {
		return nil, errors.New("Status invalido!!")
	}
	orderRepo := repository.NewOrderRepository(db)
	if err := orderRepo.UpdateStatus(ctx, ID, status); err != nil {
		return nil, err
	}

	order, err := orderRepo.GetById(ctx, ID)
	return order, err

}
