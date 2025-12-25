package models

import "github.com/google/uuid"

type Item struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID   uuid.UUID `gorm:"type:uuid;index"`
	ProductID uuid.UUID `gorm:"type:uuid"`
	Quantity  int       `gorm:"not null"`
	Price     int64     `gorm:"not null"` // копейки
}
