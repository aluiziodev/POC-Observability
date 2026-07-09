package domain

import "time"

type Order struct {
	ID        string      `json:"id"`
	Item      string      `json:"item"`
	Quantity  int         `json:"quantity"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
