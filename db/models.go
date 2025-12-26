package db

import "time"

type User struct {
	ID        string    `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Name      *string   `db:"name" json:"name,omitempty"`
	Password  string    `db:"password" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Product struct {
	ID         string    `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Price      float64   `db:"price" json:"price"`
	OutOfStock bool      `db:"out_of_stock" json:"out_of_stock"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

type Order struct {
	ID         string    `db:"id" json:"id"`
	UserID     string    `db:"user_id" json:"user_id"`
	ProductID  string    `db:"product_id" json:"product_id"`
	Quantity   int       `db:"quantity" json:"quantity"`
	TotalPrice float64   `db:"total_price" json:"total_price"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
