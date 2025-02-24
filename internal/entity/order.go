package entity

import "context"

type Order struct {
	ID           string
	CustomerName string
	Status       string
	Amount       int
	Items        []string
}

type OrderRepository interface {
	ListOrders(ctx context.Context, customerName, status string) ([]Order, error)
	GetOrder(ctx context.Context, id string) (*Order, error)
	CreateOrder(ctx context.Context, order Order) (*Order, error)
	UpdateOrder(ctx context.Context, order Order) (*Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

var (
	OrderStatusActive   = "active"
	OrderStatusComplete = "complete"
)

type CreateOrderDTO struct {
	CustomerName string
	Amount       int
	Items        []string
}
