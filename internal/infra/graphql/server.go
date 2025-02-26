package graphql

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tclemos/go-expert-desafio-clean-arch/config"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/graphql/graph"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/grpc/gen/pb"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start(config config.GraphQLConfig) {
	// grpc service clients
	grpcAddr := fmt.Sprintf("%s:%d", config.GRPCHost, config.GRPCPort)
	grpcConn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server %s: %v", grpcAddr, err)
	}
	defer grpcConn.Close()
	orderServiceClient := pb.NewOrderServiceClient(grpcConn)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(orderServiceClient)}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	log.Println("GraphQL server is running on:", addr)
	http.ListenAndServe(addr, nil)

}
