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

func (s *ProductService) DeleteProduct(id int) error {

	return s.repo.DeleteProduct(id)
}

func (s *ProductService) CreateProduct(req models.Product) error {

	req.Created_at = time.Now()

	return s.repo.CreateProduct(req)
}

func (s *ProductService) GetProductList() ([]models.ProductList, error) {
	return s.repo.GetProductList()
}

func (s *ProductService) GetProductDetail(productID int) (*models.ProductDetail, error) {
	return s.repo.GetProductDetail(productID)
}

func (s *ProductService) GetRecommendedProduct() ([]models.ProductList, error) {
	return s.repo.GetRecommendedProduct()
}

func (s *ProductService) GetProductReview() ([]models.ProductReview, error) {
	return s.repo.GetProductReview()
}