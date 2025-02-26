package main

import (
	"github.com/tclemos/go-expert-desafio-clean-arch/config"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/graphql"
)

func main() {
	var graphQLConfig config.GraphQLConfig
	if err := config.LoadConfig(".env", &graphQLConfig); err != nil {
		panic(err)
	}

	graphql.Start(graphQLConfig)
}
