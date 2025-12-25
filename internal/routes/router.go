package routes

import (
	"goshop/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(authService *services.AuthService, orderService *services.OrderService, productService *services.ProductService, transactionService *services.TransactionService, userService *services.UserService, buyerService *services.BuyerService) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	authHandler := NewAuthHandler(authService)
	orderHandler := NewOrderHandler(orderService, buyerService)
	productHandler := NewProductHandler(productService)
	transactionHandler := NewTransactionHandler(transactionService)
	userHandler := NewUserHandler(userService)

	router.Post("/register", authHandler.Register)
	router.Post("/login", authHandler.Login)

	router.Group(func(r chi.Router) {
		r.Use(AuthMiddleware(authService.GetSecretKey()))
		r.Post("/orders", orderHandler.CreateOrder)
		r.Get("/orders", orderHandler.GetOrders)
		r.Post("/cancel", orderHandler.CancelOrder)
		r.Post("/orders/pay", orderHandler.PayOrder)

		r.Get("/products", productHandler.GetProducts)
		r.Get("/product", productHandler.GetProduct)
		r.Post("/products", productHandler.CreateProduct)
		r.Put("/products", productHandler.UpdateProduct)
		r.Delete("/product", productHandler.DeleteProduct)

		r.Get("/transactions", transactionHandler.GetTransactions)

		r.Get("/user", userHandler.GetUser)
	})

	return router
}
