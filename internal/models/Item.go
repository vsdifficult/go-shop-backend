package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	OrderID   uuid.UUID `gorm:"type:uuid"`
	ProductID uuid.UUID `gorm:"type:uuid"`
	Quantity  int       `gorm:"not null"`
	Price     float64   `gorm:"not null"`
}
