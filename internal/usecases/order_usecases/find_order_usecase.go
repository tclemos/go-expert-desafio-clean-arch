package order_usecases

import (
	"context"

	"github.com/tclemos/go-expert-desafio-clean-arch/internal/entity"
)

type FindOrderUseCase struct {
	orderRepository entity.OrderRepository
}

func NewFindOrderUseCase(orderRepository entity.OrderRepository) *FindOrderUseCase {
	return &FindOrderUseCase{
		orderRepository: orderRepository,
	}
}

func (u *FindOrderUseCase) FindOrder(ctx context.Context, id string) (*entity.Order, error) {
	return u.orderRepository.GetOrder(ctx, id)
}

func (u *FindOrderUseCase) FindOrders(ctx context.Context, customerName string, status string) ([]entity.Order, error) {
	return u.orderRepository.ListOrders(ctx, customerName, status)
}
