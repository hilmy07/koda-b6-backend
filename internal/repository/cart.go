package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CartRepository struct {
	db *pgxpool.Pool
}

func NewCartRepository(db *pgxpool.Pool) *CartRepository {
	return &CartRepository{db: db}
}

func (c *CartRepository) CreateCartItem(req models.Cart) error {

	_, err := c.db.Query(context.Background(), `INSERT INTO carts (quantity, size, variant, user_id, product_id) VALUES ($1,$2,$3,$4,$5) ON CONFLICT (user_id, product_id, size, variant) DO UPDATE SET quantity = "carts".quantity + EXCLUDED.quantity`, req.Quantity, req.Size, req.Variant, req.UserId, req.ProductId)

	return err
}
