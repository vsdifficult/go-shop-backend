package main

import (
	"fmt"
	"goshop/internal/config"
	"goshop/internal/models"
	"goshop/internal/repositories"
	"goshop/internal/routes"
	"goshop/internal/services"
	db "goshop/internal/storage/postgres"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting goshop")
	log.Debug("logger debug mode enabled")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.ConnectionString.Host,
		cfg.ConnectionString.Port,
		cfg.ConnectionString.DB_User,
		cfg.ConnectionString.DB_Password,
		cfg.ConnectionString.DB_Name,
	)
	db, err := db.New(dsn)
	if err != nil {
		log.Error(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Order{}, &models.Item{})
	log.Info("db migrated")
	userRepo := repositories.NewUserRepository(db.DB)
	orderRepo := repositories.NewOrderRepository(db.DB)
	authService := services.NewAuthService(log, userRepo, cfg.SecretKey, cfg.HTTPServer.Timeout)
	orderService := services.NewOrderService(log, orderRepo, userRepo)

	router := routes.NewRouter(authService, orderService)

	log.Info("starting server", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
