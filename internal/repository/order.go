package repository

import (
	"backend/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

func (c *OrderRepository) CreateOrderItem(req models.Order, userId int) error {
	now := time.Now()

	_, err := c.db.Exec(context.Background(), `
		INSERT INTO orders 
		(cart_id, total, status, fullname, phone, email, address, delivery, created_at, updated_at) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`, req.Cart_id, req.Total, req.Status, req.Fullname, req.Phone, req.Email, req.Address, req.Delivery, now, now)

	return err
}