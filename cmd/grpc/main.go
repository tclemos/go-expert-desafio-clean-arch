package main

import (
	"github.com/tclemos/go-expert-desafio-clean-arch/config"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/grpc"
)

func main() {
	var grpcConfig config.GRPCConfig
	if err := config.LoadConfig(".env", &grpcConfig); err != nil {
		panic(err)
	}

	var dbConfig config.DBConfig
	if err := config.LoadConfig(".env", &dbConfig); err != nil {
		panic(err)
	}

	grpc.Start(grpcConfig, dbConfig)
}
