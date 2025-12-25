package repositories

import (
	"goshop/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *models.Transaction) error
	GetByID(id uuid.UUID) (*models.Transaction, error)
	GetByBuyerID(buyerID uuid.UUID) ([]models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(tx *models.Transaction) error {
	return r.db.Create(tx).Error
}

func (r *transactionRepository) GetByID(id uuid.UUID) (*models.Transaction, error) {
	var tx models.Transaction
	err := r.db.First(&tx, "id = ?", id).Error
	return &tx, err
}

func (r *transactionRepository) GetByBuyerID(buyerID uuid.UUID) ([]models.Transaction, error) {
	var txs []models.Transaction
	err := r.db.Find(&txs, "buyer_id = ?", buyerID).Error
	return txs, err
}
