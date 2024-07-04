package types

import "time"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Order struct {
	ID          string    `json:"id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Timestamp   time.Time `json:"timestamp"`
	CustomerID  string    `json:"customer_id"`
}

type Customer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
