package services

import (
	"goshop/internal/models"
	"goshop/internal/repositories"

	"github.com/google/uuid"
)

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetProductByID(productID uuid.UUID) (*models.Product, error) {
	return s.repo.GetByID(productID)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(productID uuid.UUID) error {
	return s.repo.Delete(productID)
}
