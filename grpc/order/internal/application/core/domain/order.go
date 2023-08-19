package domain

import (
	"time"
)

type Order struct {
	ID         int64       `json:"id"`
	CustomerId int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
}

type OrderItem struct {
	ProductCode int64 `json:"product_code"`
	UnitPrice   int64 `json:"unit_price"`
	Quantity    int64 `json:"quantity"`
}

func CreateOrder(customerId int64, orderItems []OrderItem) Order {
	return Order{
		Status:     "Pending",
		CreatedAt:  time.Now().Unix(),
		CustomerId: customerId,
		OrderItems: orderItems,
	}
}
