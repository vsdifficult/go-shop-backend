package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey"`
	OrderID   uuid.UUID       `gorm:"type:uuid;not null;index"`
	BuyerID   uuid.UUID       `gorm:"type:uuid;not null;index"`
	Cost      decimal.Decimal `gorm:"type:numeric(12,2);not null"`
	CreatedAt time.Time
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New()
	return nil
}
