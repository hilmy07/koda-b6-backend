package service

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) CreateCart(req models.Cart) error {
	return s.repo.CreateCartItem(req)
}

