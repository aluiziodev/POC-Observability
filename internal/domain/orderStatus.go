package domain

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "confirmed"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

func (status OrderStatus) IsValid() bool {
	switch status {
	case OrderStatusPending,
		OrderStatusProcessing,
		OrderStatusShipped,
		OrderStatusDelivered,
		OrderStatusCancelled:
		return true
	default:
		return false
	}
}
