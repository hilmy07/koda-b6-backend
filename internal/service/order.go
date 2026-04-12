package service

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (c *OrderService) CreateOrder(req models.Order, userId int) error {
	return c.repo.CreateOrderItem(req, userId)
}