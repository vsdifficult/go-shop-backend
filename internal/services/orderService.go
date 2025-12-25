package services

import (
	"errors"
	"fmt"
	"log/slog"

	"goshop/internal/models"
	"goshop/internal/repositories"

	"github.com/google/uuid"
)

type OrderService struct {
	log       *slog.Logger
	orderRepo repositories.OrderRepository
	userRepo  repositories.UserRepository
}

func NewOrderService(
	log *slog.Logger,
	orderRepo repositories.OrderRepository,
	userRepo repositories.UserRepository,
) *OrderService {
	return &OrderService{
		log:       log,
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (s *OrderService) CreateOrder(userID uuid.UUID) (*models.Order, error) {
	const op = "OrderService.CreateOrder"
	log := s.log.With(slog.String("op", op))

	if _, err := s.userRepo.GetByID(userID); err != nil {
		log.Error("user not found", slog.Any("err", err))
		return nil, fmt.Errorf("%s: user not found", op)
	}

	order := &models.Order{
		UserID: userID,
	}

	if err := s.orderRepo.Create(order); err != nil {
		log.Error("failed to create order", slog.Any("err", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("order created", slog.String("orderID", order.ID.String()))
	return order, nil
}

func (s *OrderService) AddProduct(
	orderID uuid.UUID,
	productID uuid.UUID,
	quantity int,
	price int64,
) error {
	const op = "OrderService.AddProduct"

	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	if err := s.orderRepo.AddProduct(orderID, productID, quantity, price); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *OrderService) GetOrder(orderID uuid.UUID) (*models.Order, error) {
	const op = "OrderService.GetOrder"

	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return order, nil
}

func (s *OrderService) GetOrders(userID uuid.UUID) ([]models.Order, error) {
	const op = "OrderService.GetOrders"

	orders, err := s.orderRepo.GetOrdersByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return orders, nil
}

func (s *OrderService) CancelOrder(orderID uuid.UUID) error {
	const op = "OrderService.CancelOrder"

	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if order.Status != models.OrderStatusDraft {
		return errors.New("only draft order can be canceled")
	}

	order.Status = models.OrderStatusCanceled
	return nil
}

