package graph

import "github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/grpc/gen/pb"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	orderServiceClient pb.OrderServiceClient
}

func NewResolver(orderServiceClient pb.OrderServiceClient) *Resolver {
	return &Resolver{
		orderServiceClient: orderServiceClient,
	}
}
