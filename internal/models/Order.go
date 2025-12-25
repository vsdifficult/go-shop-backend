package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusDraft    OrderStatus = "DRAFT"
	OrderStatusPaid     OrderStatus = "PAID"
	OrderStatusCanceled OrderStatus = "CANCELED"
)

type Order struct {
	ID        uuid.UUID   `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID   `gorm:"type:uuid;not null;index"`
	Status    OrderStatus `gorm:"type:varchar(20);not null"`
	Items     []Item      `gorm:"foreignKey:OrderID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
