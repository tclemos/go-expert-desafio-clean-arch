package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/entity"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/grpc/gen/pb"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/infra/webserver/dto"
)

type OrderHandlers struct {
	orderServiceClient pb.OrderServiceClient
}

func NewOrderHandlers(orderServiceClient pb.OrderServiceClient) *OrderHandlers {
	return &OrderHandlers{
		orderServiceClient: orderServiceClient,
	}
}

func (h *OrderHandlers) ListOrders(w http.ResponseWriter, r *http.Request) {
	customerName := r.URL.Query().Get("customer_name")
	status := r.URL.Query().Get("status")

	orderResponse, err := h.orderServiceClient.ListOrders(r.Context(), &pb.ListOrdersRequest{
		CustomerName: customerName,
		Status:       status,
	})
	if err != nil {
		render.Render(w, r, dto.ErrInternalServerError(err))
		return
	}

	orders := make([]dto.OrderDTO, 0, len(orderResponse.Orders))
	for _, order := range orderResponse.Orders {
		orders = append(orders, dto.OrderDTO{
			ID:           order.Id,
			CustomerName: order.CustomerName,
			Status:       order.Status,
			Amount:       int(order.Amount),
			Items:        order.Items,
		})
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, orders)
}

func (h *OrderHandlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	data := &dto.CreateOrderDTO{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, dto.ErrBadRequest(err))
		return
	}

	orderResponse, err := h.orderServiceClient.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerName: data.CustomerName,
		Items:        data.Items,
		Amount:       int32(data.Amount),
	})
	if err != nil {
		render.Render(w, r, dto.ErrInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, dto.OrderDTO{
		ID:           orderResponse.Order.Id,
		CustomerName: orderResponse.Order.CustomerName,
		Status:       orderResponse.Order.Status,
		Amount:       int(orderResponse.Order.Amount),
		Items:        orderResponse.Order.Items,
	})
}

func (h *OrderHandlers) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, dto.ErrBadRequest(fmt.Errorf("missing required id field")))
		return
	}

	_, err := uuid.Parse(id)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, dto.ErrBadRequest(fmt.Errorf("id must be a valid uuid")))
		return
	}

	orderResponse, err := h.orderServiceClient.GetOrder(r.Context(), &pb.GetOrderRequest{
		Id: id,
	})
	if err != nil {
		if strings.Contains(err.Error(), entity.ErrNotFound.Error()) {
			render.Status(r, http.StatusNotFound)
			render.Render(w, r, dto.ErrNotFound(fmt.Errorf("order not found")))
			return
		}
		render.Render(w, r, dto.ErrInternalServerError(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, dto.OrderDTO{
		ID:           orderResponse.Order.Id,
		CustomerName: orderResponse.Order.CustomerName,
		Status:       orderResponse.Order.Status,
		Amount:       int(orderResponse.Order.Amount),
		Items:        orderResponse.Order.Items,
	})
}

func (h *OrderHandlers) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	data := &dto.UpdateOrderDTO{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, dto.ErrBadRequest(err))
		return
	}

	orderResponse, err := h.orderServiceClient.UpdateOrder(r.Context(), &pb.UpdateOrderRequest{
		Id:           data.ID,
		CustomerName: data.CustomerName,
		Status:       data.Status,
		Items:        data.Items,
		Amount:       int32(data.Amount),
	})
	if err != nil {
		if strings.Contains(err.Error(), entity.ErrNotFound.Error()) {
			render.Status(r, http.StatusNotFound)
			render.Render(w, r, dto.ErrNotFound(fmt.Errorf("order not found")))
			return
		}
		render.Render(w, r, dto.ErrInternalServerError(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, dto.OrderDTO{
		ID:           orderResponse.Order.Id,
		CustomerName: orderResponse.Order.CustomerName,
		Status:       orderResponse.Order.Status,
		Amount:       int(orderResponse.Order.Amount),
		Items:        orderResponse.Order.Items,
	})
}

func (h *OrderHandlers) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, dto.ErrBadRequest(fmt.Errorf("missing required id field")))
		return
	}

	_, err := uuid.Parse(id)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, dto.ErrBadRequest(fmt.Errorf("id must be a valid uuid")))
		return
	}

	_, err = h.orderServiceClient.DeleteOrder(r.Context(), &pb.DeleteOrderRequest{
		Id: id,
	})
	if err != nil {
		if strings.Contains(err.Error(), entity.ErrNotFound.Error()) {
			render.Status(r, http.StatusNotFound)
			render.Render(w, r, dto.ErrNotFound(fmt.Errorf("order not found")))
			return
		}
		render.Render(w, r, dto.ErrInternalServerError(err))
		return
	}

	render.Status(r, http.StatusOK)
}
