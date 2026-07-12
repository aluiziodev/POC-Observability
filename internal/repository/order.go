package repository

import (
	"context"
	"database/sql"
	"order-service/internal/domain"
	"time"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (repo *OrderRepository) Create(ctx context.Context, order *domain.Order) error {

	if err := ctx.Err(); err != nil {
		return err
	}

	now := time.Now().UTC()
	order.CreatedAt = now
	order.UpdatedAt = now

	_, err := repo.db.ExecContext(ctx, `
		INSERT INTO orders (id, item, quantity, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, order.ID, order.Item, order.Quantity, order.Status, order.CreatedAt, order.UpdatedAt)
	return err
}

func (repo *OrderRepository) GetAll(ctx context.Context) ([]*domain.Order, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	rows, err := repo.db.QueryContext(ctx, `
		SELECT id, item, quantity, status, created_at, updated_at
		FROM orders
		ORDER BY status
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*domain.Order, 0)
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.ID, &order.Item, &order.Quantity, &order.Status,
			&order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (repo *OrderRepository) GetById(ctx context.Context, ID string) (*domain.Order, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	var order domain.Order

	err := repo.db.QueryRowContext(ctx, `
		SELECT id, item, quantity, status, created_at, updated_at 
		FROM orders WHERE id=$1
	`, ID).Scan(&order.ID, &order.Item, &order.Quantity, &order.Status,
		&order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &order, nil

}

func (repo *OrderRepository) UpdateStatus(ctx context.Context, ID string, status domain.OrderStatus) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	now := time.Now().UTC()

	_, err := repo.db.ExecContext(ctx, `
        UPDATE orders 
		SET status = $1, updated_at = $2
		WHERE id = $2
	`, status, now, ID)

	return err
}
