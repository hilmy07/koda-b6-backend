package models

import "time"

type Product struct {
	Id           int     `json:"id"`
	Name_product string  `json:"name_product"`
	Description  string  `json:"description"`
	Base_price   float64 `json:"base_price"`
	Stock        float64 `json:"stock"`
	Created_at   time.Time
	Updated_at   time.Time
}

type ProductList struct {
	Id          int     `json:"id"`
	NameProduct string  `json:"name_product"`
	Description string  `json:"description"`
	BasePrice   float64 `json:"base_price"`
	Image       string  `json:"image"`
	Rating      float64 `json:"rating"`
}

type PaginatedProducts struct {
	Page     int           `json:"page"`
	Limit    int           `json:"limit"`
	Total    int           `json:"total"`    // total product di DB
	Products []ProductList `json:"products"` // daftar produk per page
}

type ProductDetail struct {
	ID          int      `json:"id"`
	NameProduct string   `json:"name_product"`
	BasePrice   float64  `json:"base_price"`
	Images      []string `json:"images"`
	Rating      float64 `json:"rating"`
	ReviewCount int      `json:"review_count"`
	Sizes       []string `json:"sizes"`
	Variants    []string `json:"variants"`
}

type ProductVariant struct {
	Id        int  `json:"id"`
	ProductID int `json:"product_id"`
	VariantName string `json:"variant_name"`
	AddPrice  float64 `json:"add_price"`
}

type ProductSize struct {
	Id        int `json:"id"`
	ProductID int `json:"product_id"`
	SizeName  string `json:"size_name"`
	AddPrice  float64 `json:"add_price"`
}

type ProductImage struct {
	Id        int `json:"id"`
	ProductID int `json:"product_id"`
	Path      string  `json:"path"`
}

type ProductReview struct {
	Id        int `json:"id"`
	Fullname  string `json:"fullname"`
	Message   string `json:"message"`
	Rating    int `json:"rating"`
	// CreatedAt time.Time 
}
