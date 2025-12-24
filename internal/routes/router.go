package routes

import (
	"goshop/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(authService *services.AuthService, orderService *services.OrderService) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	authHandler := NewAuthHandler(authService)
	orderHandler := NewOrderHandler(orderService)

	router.Post("/register", authHandler.Register)
	router.Post("/login", authHandler.Login)

	router.Group(func(r chi.Router) {
		r.Use(AuthMiddleware(authService.GetSecretKey()))
		r.Post("/orders", orderHandler.CreateOrder)
		r.Get("/orders", orderHandler.GetOrders)
		r.Post("/cancel", orderHandler.CancelOrder)
	})

	return router
}
