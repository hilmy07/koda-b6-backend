package models

import "time"

type Order struct {
	Id         int     `json:"order_id" db:"id"`
	Cart_id    int     `json:"cart_id" db:"cart_id"`
	Total 	   float64 `json:"total" db:"total"`
	Status     int     `json:"status" db:"status"`
	Fullname   string  `json:"fullname" db:"fullname"`
	Phone      string  `json:"phone" db:"phone"`
	Email      string  `json:"email" db:"email"`
	Address    string  `json:"address" db:"address"`
	Delivery   string  `json:"delivery" db:"delivery"`
	Created_at time.Time
	Updated_at time.Time
}