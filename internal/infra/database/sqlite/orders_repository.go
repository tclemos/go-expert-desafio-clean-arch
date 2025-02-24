package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/tclemos/go-expert-desafio-clean-arch/internal/entity"
)

type OrdersRepository struct {
	db *sql.DB
}

func NewOrdersRepository(db *sql.DB) *OrdersRepository {
	return &OrdersRepository{db}
}

func (r *OrdersRepository) ListOrders(ctx context.Context, customerName, status string) ([]entity.Order, error) {
	query := `SELECT id, customer_name, status, amount, items FROM orders WHERE 1=1`

	args := []interface{}{}
	if customerName != "" {
		query += ` AND customer_name = ?`
		args = append(args, customerName)
	}

	if status != "" {
		query += ` AND status = ?`
		args = append(args, status)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := []entity.Order{}
	for rows.Next() {
		order := entity.Order{}
		var items string
		err := rows.Scan(&order.ID, &order.CustomerName, &order.Status, &order.Amount, &items)
		order.Items = strings.Split(items, ",")
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrdersRepository) GetOrder(ctx context.Context, id string) (*entity.Order, error) {
	query := `SELECT id, customer_name, status, amount, items FROM orders WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)
	order := entity.Order{}
	var items string
	err := row.Scan(&order.ID, &order.CustomerName, &order.Status, &order.Amount, &items)
	order.Items = strings.Split(items, ",")
	if errors.Is(err, sql.ErrNoRows) {
		return nil, entity.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrdersRepository) CreateOrder(ctx context.Context, order entity.Order) (*entity.Order, error) {
	query := `INSERT INTO orders (id, customer_name, status, amount, items) VALUES (?, ?, ?, ?, ?)`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	items := strings.Join(order.Items, ",")
	_, err = statement.ExecContext(ctx, order.ID, order.CustomerName, order.Status, order.Amount, items)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrdersRepository) UpdateOrder(ctx context.Context, order entity.Order) (*entity.Order, error) {
	query := `UPDATE orders SET customer_name = ?, status = ?, amount = ?, items = ? WHERE id = ?`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	items := strings.Join(order.Items, ",")
	result, err := statement.ExecContext(ctx, order.CustomerName, order.Status, order.Amount, items, order.ID)
	if err != nil {
		return nil, err
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return nil, entity.ErrNotFound
	}
	return &order, nil
}

func (r *OrdersRepository) DeleteOrder(ctx context.Context, id string) error {
	query := `DELETE FROM orders WHERE id = ?`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	result, err := statement.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return entity.ErrNotFound
	}
	return nil
}
