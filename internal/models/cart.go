package models

import "time"

type Cart struct {
	Id         int    `json:"cart_id" db:"id"`
	Quantity   int    `json:"quantity" db:"quantity"`
	Size       string `json:"size" db:"size"`
	Variant    string `json:"variant" db:"variant"`
	UserId     int    `json:"user_id" db:"user_id"`
	ProductId  int    `json:"product_id" db:"product_id"`
	Created_at time.Time
	Updated_at time.Time
}

type CartByUserID struct {
	Id           int       `json:"cart_id" db:"id"`
	Quantity     int       `json:"quantity" db:"quantity"`
	Size         string    `json:"size" db:"size"`
	Variant      string    `json:"variant" db:"variant"`
	UserId       int       `json:"user_id" db:"user_id"`
	ProductName  string    `json:"name_product" db:"name_product"`
	Price        int       `json:"price" db:"base_price"`
	Created_at   time.Time
	Updated_at   time.Time
}

type GetCartRequest struct {
    UserId int `json:"user_id"`
}