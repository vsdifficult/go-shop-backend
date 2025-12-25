package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string
	Price       int64     `gorm:"not null"` // копейки
	Stock       int       `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
