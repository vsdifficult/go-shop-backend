package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Name      string    `gorm:"type:string"`
	ProductID uuid.UUID `gorm:"type:uuid"`
	Quantity  int       `gorm:"not null"`
	Price     float64   `gorm:"not null"`
}

type OrderCreateDto struct {
	Name      string    `json:"name"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
}

type OrderDto struct {
	Name      string    `json:"name"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
}
