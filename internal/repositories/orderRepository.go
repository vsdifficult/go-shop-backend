package repositories

import (
	"goshop/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetByID(orderID uuid.UUID) (*models.Order, error)
	GetOrdersByUserID(userID uuid.UUID) ([]models.Order, error)
	AddProduct(orderID, productID uuid.UUID, quantity int, price int64) error
	MarkAsPaid(orderID uuid.UUID) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *models.Order) error {
	order.ID = uuid.New()
	order.Status = models.OrderStatusDraft
	return r.db.Create(order).Error
}

func (r *orderRepository) GetByID(orderID uuid.UUID) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Items").
		First(&order, "id = ?", orderID).Error
	return &order, err
}

func (r *orderRepository) GetOrdersByUserID(userID uuid.UUID) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Items").
		Where("user_id = ?", userID).
		Find(&orders).Error
	return orders, err
}

func (r *orderRepository) AddProduct(
	orderID, productID uuid.UUID,
	quantity int,
	price int64,
) error {
	item := models.Item{
		ID:        uuid.New(),
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
	return r.db.Create(&item).Error
}

func (r *orderRepository) MarkAsPaid(orderID uuid.UUID) error {
	return r.db.Model(&models.Order{}).
		Where("id = ?", orderID).
		Update("status", models.OrderStatusPaid).Error
}
