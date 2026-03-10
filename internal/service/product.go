package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"time"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(req models.Product) error {

	req.Created_at = time.Now()

	return s.repo.CreateProduct(req)
}

func (s *ProductService) GetProductList() ([]models.ProductList, error) {
	return s.repo.GetProductList()
}
