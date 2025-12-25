package services

import (
	"errors"
	"log/slog"

	"goshop/internal/models"
	"goshop/internal/repositories"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type BuyerService struct {
	log             *slog.Logger
	orderRepo       repositories.OrderRepository
	userRepo        repositories.UserRepository
	transactionRepo repositories.TransactionRepository
}

func NewBuyerService(
	log *slog.Logger,
	orderRepo repositories.OrderRepository,
	userRepo repositories.UserRepository,
	transactionRepo repositories.TransactionRepository,
) *BuyerService {
	return &BuyerService{
		log:             log,
		orderRepo:       orderRepo,
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *BuyerService) AddProductInOrder(
	orderID, productID uuid.UUID,
	quantity int,
	price int64,
) error {
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	return s.orderRepo.AddProduct(orderID, productID, quantity, price)
}

func (s *BuyerService) BuyOrder(
	orderID, buyerID uuid.UUID,
) (uuid.UUID, error) {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return uuid.Nil, err
	}

	if order.UserID != buyerID {
		return uuid.Nil, errors.New("order does not belong to user")
	}

	if order.Status != models.OrderStatusDraft {
		return uuid.Nil, errors.New("order cannot be paid")
	}

	var total int64
	for _, item := range order.Items {
		total += int64(item.Quantity) * item.Price
	}

	tx := &models.Transaction{
		OrderID: order.ID,
		BuyerID: buyerID,
		Cost:    decimal.NewFromInt(total).Div(decimal.NewFromInt(100)),
	}

	if err := s.transactionRepo.Create(tx); err != nil {
		return uuid.Nil, err
	}

	if err := s.orderRepo.MarkAsPaid(order.ID); err != nil {
		return uuid.Nil, err
	}

	return tx.ID, nil
}
