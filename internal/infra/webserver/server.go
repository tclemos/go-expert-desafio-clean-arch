package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/tclemos/go-expert-desafio-clean-arch/config"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/grpc/gen/pb"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/webserver/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start(config config.RESTConfig) {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// headers
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// grpc service clients
	grpcAddr := fmt.Sprintf("%s:%d", config.GRPCHost, config.GRPCPort)
	grpcConn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server %s: %v", grpcAddr, err)
	}
	defer grpcConn.Close()
	orderClient := pb.NewOrderServiceClient(grpcConn)

	// handlers
	orderHandlers := handlers.NewOrderHandlers(orderClient)

	// REST routes
	r.Route("/orders", func(r chi.Router) {
		r.Get("/", orderHandlers.ListOrders)
		r.Get("/{id}", orderHandlers.GetOrder)

		r.Post("/", orderHandlers.CreateOrder)
		r.Put("/", orderHandlers.UpdateOrder)
		r.Delete("/{id}", orderHandlers.DeleteOrder)
	})

	// start REST server
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	log.Println("REST server is running on:", addr)
	http.ListenAndServe(addr, r)
}
