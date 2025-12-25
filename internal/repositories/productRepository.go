package repositories

import (
	"goshop/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(productID uuid.UUID) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Update(product *models.Product) error
	Delete(productID uuid.UUID) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) error {
	product.ID = uuid.New()
	return r.db.Create(product).Error
}

func (r *productRepository) GetByID(productID uuid.UUID) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, "id = ?", productID).Error
	return &product, err
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(productID uuid.UUID) error {
	return r.db.Delete(&models.Product{}, "id = ?", productID).Error
}
