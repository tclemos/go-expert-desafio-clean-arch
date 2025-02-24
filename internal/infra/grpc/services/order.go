package services

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/entity"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/grpc/gen/pb"
)

type CreateOrderUseCase interface {
	CreateOrder(ctx context.Context, createOrder entity.CreateOrderDTO) (*entity.Order, error)
	UpdateOrder(ctx context.Context, order entity.Order) (*entity.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

type FindOrderUseCase interface {
	FindOrder(ctx context.Context, id string) (*entity.Order, error)
	FindOrders(ctx context.Context, customerName string, status string) ([]entity.Order, error)
}

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	createOrderUseCase CreateOrderUseCase
	findOrderUseCase   FindOrderUseCase
}

func NewOrderService(createOrderUseCase CreateOrderUseCase, findOrderUseCase FindOrderUseCase) *OrderService {
	return &OrderService{
		createOrderUseCase: createOrderUseCase,
		findOrderUseCase:   findOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	order, err := s.createOrderUseCase.CreateOrder(ctx, entity.CreateOrderDTO{
		CustomerName: req.CustomerName,
		Amount:       int(req.Amount),
		Items:        req.Items,
	})
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Order: &pb.Order{
			Id:           order.ID,
			CustomerName: order.CustomerName,
			Status:       order.Status,
			Amount:       int32(order.Amount),
			Items:        order.Items,
		},
	}, nil
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.OrderResponse, error) {
	order, err := s.createOrderUseCase.UpdateOrder(ctx, entity.Order{
		ID:           req.Id,
		CustomerName: req.CustomerName,
		Amount:       int(req.Amount),
		Items:        req.Items,
		Status:       req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{
		Order: &pb.Order{
			Id:           order.ID,
			CustomerName: order.CustomerName,
			Status:       order.Status,
			Amount:       int32(order.Amount),
			Items:        order.Items,
		},
	}, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*empty.Empty, error) {
	err := s.createOrderUseCase.DeleteOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.findOrderUseCase.FindOrders(ctx, req.CustomerName, req.Status)
	if err != nil {
		return nil, err
	}

	ordersResponse := make([]*pb.Order, 0, len(orders))
	for _, order := range orders {
		ordersResponse = append(ordersResponse, &pb.Order{
			Id:           order.ID,
			CustomerName: order.CustomerName,
			Status:       order.Status,
			Amount:       int32(order.Amount),
			Items:        order.Items,
		})
	}

	return &pb.ListOrdersResponse{
		Orders: ordersResponse,
	}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	order, err := s.findOrderUseCase.FindOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Order: &pb.Order{
			Id:           order.ID,
			CustomerName: order.CustomerName,
			Status:       order.Status,
			Amount:       int32(order.Amount),
			Items:        order.Items,
		},
	}, nil
}
