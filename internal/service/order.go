package service

import (
	"context"
	"errors"
	"order-service/internal/domain"
	"order-service/internal/models"
	"order-service/internal/repository"

	"github.com/google/uuid"
)

func CreateOrder(ctx context.Context, input models.CreateOrderInput) (*domain.Order, error) {
	input.Validate()

	order := &domain.Order{
		ID:       uuid.NewString(),
		Item:     input.Item,
		Quantity: input.Quantity,
		Status:   domain.OrderStatusPending,
	}

	orderRepo := repository.NewOrderRepository()
	if err := orderRepo.Create(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func ListOrders(ctx context.Context) ([]*domain.Order, error) {
	orderRepo := repository.NewOrderRepository()
	orders, err := orderRepo.GetAll(ctx)
	return orders, err
}

func GetOrder(ctx context.Context, ID string) (*domain.Order, error) {
	orderRepo := repository.NewOrderRepository()
	order, err := orderRepo.GetById(ctx, ID)
	return order, err
}

func UpdateOrderStatus(ctx context.Context, ID string, status domain.OrderStatus) (*domain.Order, error) {
	if !status.IsValid() {
		return nil, errors.New("Status invalido!!")
	}
	orderRepo := repository.NewOrderRepository()
	order, err := orderRepo.UpdateStatus(ctx, ID, status)
	return order, err

}
