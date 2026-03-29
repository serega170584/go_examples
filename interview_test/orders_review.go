package orders

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Customer struct {
	ID        string
	Name      string
	Address   string
	CreatedAt time.Time
}

type OrderItem struct {
	ProductID string
	Quantity  int
	Price     int
}

type Order struct {
	ID         string
	Customer   Customer
	Items      []OrderItem
	TotalSum   int
	Status     string
	ExternalID string
}

type OrdersRepo struct {
	DB *sql.DB
}

func New(db *sql.DB) OrdersRepo {
	return OrdersRepo{
		DB: db,
	}
}

func (o *OrdersRepo) GetOrdersByDate(ctx context.Context, date time.Time, onlyCompleted bool) ([]Order, error) {
	query := "SELECT id, customer_id, total_sum, status FROM orders WHERE order_date = ?"
	if onlyCompleted {
		query += " AND status = 'completed'"
	}

	rows, err := o.DB.Query(query, date.Format("2006-01-02"))
	if err != nil {
		return nil, errors.New("ошибка получения заказов")
	}

	var orders []Order

	for rows.Next() {
		var order Order
		var customerID string

		err := rows.Scan(&order.ID, &customerID, &order.TotalSum, &order.Status)
		if err != nil {
			return nil, err
		}

		customer, err := o.getCustomerById(ctx, customerID)
		if err != nil {
			return nil, errors.New("ошибка получения заказов")
		}
		order.Customer = customer

		itemRows, err := o.DB.QueryContext(ctx,
			"SELECT product_id, quantity, price FROM order_items WHERE order_id = ?", order.ID)
		if err != nil {
			return nil, errors.New("ошибка получения заказов")
		}

		var items []OrderItem
		for itemRows.Next() {
			var item OrderItem

			err := rows.Scan(&item.ProductID, &item.Quantity, &item.Price)
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}

		order.Items = items

		// Получаем внешний ID из CRM системы
		externalID, err := getExternalOrderIdFromCRM(order.ID)
		if err != nil {
			return nil, errors.New("ошибка получения заказов")
		}
		order.ExternalID = externalID

		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("ошибка получения заказов")
	}

	return orders, nil
}

func (o *OrdersRepo) getCustomerById(ctx context.Context, id string) (Customer, error) {
	row := o.DB.QueryRowContext(ctx,
		"SELECT id, name, address, created_at FROM customers WHERE id = ?", id)

	var customer Customer

	err := row.Scan(&customer.ID, &customer.Name, &customer.Address, &customer.CreatedAt)
	if err != nil {
		return Customer{}, err
	}

	return customer, nil
}

func getExternalOrderIdFromCRM(orderID string) (string, error) {
	resp, err := http.Get("https://api.crm-system.com/orders/" + orderID)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		ExternalID int `json:"external_id"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return strconv.Itoa(result.ExternalID), nil
}

func (o *OrdersRepo) UpdateOrderStatus(ctx context.Context, orderID, status string) error {
	_, err := o.DB.ExecContext(ctx,
		"UPDATE orders SET status = ? WHERE id = ?",
		status, orderID)
	if err != nil {
		return errors.New("ошибка обновления статуса")
	}

	return nil
}
