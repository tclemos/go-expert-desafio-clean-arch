package dto

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/tclemos/go-expert-desafio-clean-arch/internal/entity"
)

type OrderDTO struct {
	ID           string   `json:"id"`
	CustomerName string   `json:"customer_name"`
	Status       string   `json:"status"`
	Amount       int      `json:"amount"`
	Items        []string `json:"items"`
}

type CreateOrderDTO struct {
	CustomerName string   `json:"customer_name"`
	Items        []string `json:"items"`
	Amount       int      `json:"amount"`
}

func (o *CreateOrderDTO) Bind(r *http.Request) error {
	if len(o.CustomerName) == 0 {
		return errors.New("missing required customer_name field")
	}

	if len(o.Items) == 0 {
		return errors.New("order must have at least one item")
	}

	if o.Amount <= 0 {
		return errors.New("order must have a positive amount")
	}
	return nil
}

type UpdateOrderDTO struct {
	ID           string   `json:"id"`
	CustomerName string   `json:"customer_name"`
	Status       string   `json:"status"`
	Items        []string `json:"items"`
	Amount       int      `json:"amount"`
}

func (o *UpdateOrderDTO) Bind(r *http.Request) error {
	if len(o.ID) == 0 {
		return errors.New("missing required id field")
	}

	if _, err := uuid.Parse(o.ID); err != nil {
		return errors.New("id must be a valid uuid")
	}

	if len(o.CustomerName) == 0 {
		return errors.New("missing required customer_name field")
	}

	if o.Status != entity.OrderStatusActive && o.Status != entity.OrderStatusComplete {
		return errors.New("status must be active or complete")
	}

	if len(o.Items) == 0 {
		return errors.New("order must have at least one item")
	}

	if o.Amount <= 0 {
		return errors.New("order must have a positive amount")
	}
	return nil
}
