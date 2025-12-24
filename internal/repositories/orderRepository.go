package repositories

import (
	"goshop/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order *models.Order) error {
	order.ID = uuid.New()
	return r.db.Create(order).Error
}

func (r *orderRepository) GetOrders(userID uuid.UUID) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Preload("Items").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetOrder(orderID uuid.UUID) (models.Order, error) {
	var order models.Order

	if err := r.db.Preload("Items").First(&order, "id = ?", orderID).Error; err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (r *orderRepository) IsOrderExists(orderID uuid.UUID) (bool, error) {
	var order models.Order

	if err := r.db.Preload("Items").First(&order, "id = ?", orderID).Error; err != nil {
		return false, err
	}
	return true, nil
}
