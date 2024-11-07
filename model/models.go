package model

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"-"`
}

type Order struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	OrderStatus  string    `json:"order_status"`
	TotalPrice   float64   `json:"total_price"`
	DiscountCode string    `json:"discount_code,omitempty"`
	Rating       int       `json:"rating,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
