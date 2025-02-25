package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/entity"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/graphql/graph/model"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/grpc/gen/pb"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, request *model.CreateOrderRequest) (*model.Order, error) {
	if request == nil {
		return nil, fmt.Errorf("missing required fields")
	}

	orderResponse, err := r.orderServiceClient.CreateOrder(ctx, &pb.CreateOrderRequest{
		CustomerName: request.CustomerName,
		Items:        request.Items,
		Amount:       request.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &model.Order{
		ID:           orderResponse.Order.Id,
		CustomerName: orderResponse.Order.CustomerName,
		Status:       orderResponse.Order.Status,
		Amount:       orderResponse.Order.Amount,
		Items:        orderResponse.Order.Items,
	}, nil
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, request *model.UpdateOrderRequest) (*model.Order, error) {
	if request == nil {
		return nil, fmt.Errorf("missing required fields")
	}

	orderResponse, err := r.orderServiceClient.UpdateOrder(ctx, &pb.UpdateOrderRequest{
		Id:           request.ID,
		CustomerName: request.CustomerName,
		Status:       request.Status,
		Items:        request.Items,
		Amount:       request.Amount,
	})
	if err != nil {
		if strings.Contains(err.Error(), entity.ErrNotFound.Error()) {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}

	return &model.Order{
		ID:           orderResponse.Order.Id,
		CustomerName: orderResponse.Order.CustomerName,
		Status:       orderResponse.Order.Status,
		Amount:       orderResponse.Order.Amount,
		Items:        orderResponse.Order.Items,
	}, nil
}

// DeleteOrder is the resolver for the deleteOrder field.
func (r *mutationResolver) DeleteOrder(ctx context.Context, request *model.DeleteOrderRequest) (*model.Empty, error) {
	if request == nil {
		return nil, fmt.Errorf("missing required fields")
	}

	id := request.ID
	if len(id) == 0 {
		return nil, fmt.Errorf("missing required id field")
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("id must be a valid uuid")
	}

	_, err = r.orderServiceClient.DeleteOrder(ctx, &pb.DeleteOrderRequest{
		Id: id,
	})
	if err != nil {
		if strings.Contains(err.Error(), entity.ErrNotFound.Error()) {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}

	return &model.Empty{}, nil
}

// ListOrders is the resolver for the listOrders field.
func (r *queryResolver) ListOrders(ctx context.Context, request *model.ListOrdersRequest) ([]*model.Order, error) {
	if request == nil {
		request = &model.ListOrdersRequest{}
	}

	customerName := ""
	if request.CustomerName != nil {
		customerName = *request.CustomerName
	}

	status := ""
	if request.Status != nil {
		status = *request.Status
	}

	orderResponse, err := r.orderServiceClient.ListOrders(ctx, &pb.ListOrdersRequest{
		CustomerName: customerName,
		Status:       status,
	})
	if err != nil {
		return nil, err
	}

	orders := make([]*model.Order, 0, len(orderResponse.Orders))
	for _, order := range orderResponse.Orders {
		orders = append(orders, &model.Order{
			ID:           order.Id,
			CustomerName: order.CustomerName,
			Status:       order.Status,
			Amount:       order.Amount,
			Items:        order.Items,
		})
	}

	return orders, nil
}

// GetOrder is the resolver for the getOrder field.
func (r *queryResolver) GetOrder(ctx context.Context, request *model.GetOrderRequest) (*model.Order, error) {
	if request == nil || len(request.ID) == 0 {
		return nil, fmt.Errorf("missing required id field")
	}
	id := request.ID

	_, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("id must be a valid uuid")
	}

	orderResponse, err := r.orderServiceClient.GetOrder(ctx, &pb.GetOrderRequest{
		Id: id,
	})
	if err != nil {
		if strings.Contains(err.Error(), entity.ErrNotFound.Error()) {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}

	return &model.Order{
		ID:           orderResponse.Order.Id,
		CustomerName: orderResponse.Order.CustomerName,
		Status:       orderResponse.Order.Status,
		Amount:       orderResponse.Order.Amount,
		Items:        orderResponse.Order.Items,
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
