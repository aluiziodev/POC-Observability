package models

type CreateOrderRequest struct {
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}
