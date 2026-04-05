package repository

import (
	"backend/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CartRepository struct {
	db *pgxpool.Pool
}

func NewCartRepository(db *pgxpool.Pool) *CartRepository {
	return &CartRepository{db: db}
}

func (c *CartRepository) CreateCartItem(req models.Cart) error {

	now := time.Now()

	_, err := c.db.Query(context.Background(), `INSERT INTO carts (quantity, size, variant, user_id, product_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT (user_id, product_id, size, variant) DO UPDATE SET quantity = "carts".quantity + EXCLUDED.quantity`, req.Quantity, req.Size, req.Variant, req.UserId, req.ProductId, now, now)

	return err
}

func (c *CartRepository) GetCartList() ([]models.Cart, error) {

	rows, err := c.db.Query(context.Background(), `SELECT * FROM carts`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	carts, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Cart])

	if err != nil {
		return nil, err
	}

	return carts, nil
}