package order_usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/entity"
)

type CreateOrderUseCase struct {
	orderRepository entity.OrderRepository
}

func NewCreateOrderUseCase(orderRepository entity.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepository: orderRepository,
	}
}

func (u *CreateOrderUseCase) CreateOrder(ctx context.Context, createOrder entity.CreateOrderDTO) (*entity.Order, error) {
	order := entity.Order{
		ID:           uuid.New().String(),
		CustomerName: createOrder.CustomerName,
		Amount:       createOrder.Amount,
		Items:        createOrder.Items,
		Status:       entity.OrderStatusActive,
	}

	return u.orderRepository.CreateOrder(ctx, order)
}

func (u *CreateOrderUseCase) UpdateOrder(ctx context.Context, order entity.Order) (*entity.Order, error) {
	return u.orderRepository.UpdateOrder(ctx, order)
}

func (u *CreateOrderUseCase) DeleteOrder(ctx context.Context, id string) error {
	return u.orderRepository.DeleteOrder(ctx, id)
}
