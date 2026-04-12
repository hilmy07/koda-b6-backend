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

func (c *CartService) CreateCart(req models.Cart) error {
	return c.repo.CreateCartItem(req)
}

func (c *CartService) GetCartList() ([]models.Cart, error) {
	return c.repo.GetCartList()
}

func (c *CartService) GetCartByUserId(userId int) ([]models.CartByUserID, error) {
	return c.repo.GetCartByUserId(userId)
}

func (s *CartService) DeleteCartItem(cartId int, userId int) error {
	return s.repo.DeleteCartItem(cartId, userId)
}
