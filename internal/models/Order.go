package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID uuid.UUID `gorm:"type:uuid"`
	Items  []Item    `gorm:"foreignKey:OrderID"`
}
