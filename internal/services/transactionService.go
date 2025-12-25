package services

import (
	"goshop/internal/models"
	"goshop/internal/repositories"

	"github.com/google/uuid"
)

type TransactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(tx *models.Transaction) error {
	return s.repo.Create(tx)
}

func (s *TransactionService) GetTransactionByID(id uuid.UUID) (*models.Transaction, error) {
	return s.repo.GetByID(id)
}

func (s *TransactionService) GetTransactionsByBuyer(buyerID uuid.UUID) ([]models.Transaction, error) {
	return s.repo.GetByBuyerID(buyerID)
}
