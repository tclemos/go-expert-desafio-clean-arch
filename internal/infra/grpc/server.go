package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/tclemos/go-expert-desafio-clean-arch/config"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/database/sqlite"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/grpc/gen/pb"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/grpc/services"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/usecases/order_usecases"
	"google.golang.org/grpc"
)

func Start(grpcConfig config.GRPCConfig, dbConfig config.DBConfig) {
	// db connection
	db := sqlite.MustOpenConn(dbConfig)
	defer db.Close()

	// repositories
	orderRepository := sqlite.NewOrdersRepository(db)

	// use cases
	createOrderUseCase := order_usecases.NewCreateOrderUseCase(orderRepository)
	findOrderUseCase := order_usecases.NewFindOrderUseCase(orderRepository)

	// create GRPC server
	addr := fmt.Sprintf("%s:%d", grpcConfig.Host, grpcConfig.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// GRPC services
	orderService := services.NewOrderService(createOrderUseCase, findOrderUseCase)
	pb.RegisterOrderServiceServer(s, orderService)

	// start GRPC server
	log.Println("REST server is running on:", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
