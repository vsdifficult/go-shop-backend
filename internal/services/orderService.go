package services

import (
	"fmt"
	"goshop/internal/models"
	"log/slog"

	"github.com/google/uuid"
)

type OrderService struct {
	log       *slog.Logger
	orderRepo OrderRepository
	userRepo  UserRepository
}

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrders(userID uuid.UUID) ([]models.Order, error)
	GetOrder(orderID uuid.UUID) (models.Order, error)
	IsOrderExists(orderID uuid.UUID) (bool, error)
}

func NewOrderService(log *slog.Logger, orderRepo OrderRepository, userRepo UserRepository) *OrderService {
	return &OrderService{
		log:       log,
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (s *OrderService) CreateOrder(userID uuid.UUID, items []models.Item) (*models.Order, error) {
	const op = "OrderService.CreateOrder"
	log := s.log.With(slog.String("op", op))

	log.Info("creating order", slog.String("userID", userID.String()))

	order := &models.Order{
		UserID: userID,
		Items:  items,
	}

	if err := s.orderRepo.CreateOrder(order); err != nil {
		log.Error("failed to create order", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("order created successfully", slog.String("orderID", order.ID.String()))

	return order, nil
}

func (s *OrderService) GetOrders(userID uuid.UUID) ([]models.Order, error) {
	const op = "OrderService.GetOrders"
	log := s.log.With(slog.String("op", op))

	log.Info("getting orders", slog.String("userID", userID.String()))

	orders, err := s.orderRepo.GetOrders(userID)
	if err != nil {
		log.Error("failed to get orders", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("orders retrieved successfully", slog.Int("count", len(orders)))

	return orders, nil
}

func (s *OrderService) CancelOrder(orderID uuid.UUID) (bool, error) {
	const op = "OrderService.CancelOrder"
	log := s.log.With(slog.String("op", op))

	_, err := s.orderRepo.IsOrderExists(orderID)
	if err != nil {
		log.Error("failed to get order", slog.String("error", err.Error()))
		return false, fmt.Errorf("%s: %w", op, err)
	}
	return true, nil
}
