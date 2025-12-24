package services

import (
	"fmt"
	"goshop/internal/models"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	log          *slog.Logger
	userRepo     UserRepository
	secretKey    string
	tokenTimeout time.Duration
}

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

func NewAuthService(log *slog.Logger, userRepo UserRepository, secretKey string, tokenTimeout time.Duration) *AuthService {
	return &AuthService{
		log:          log,
		userRepo:     userRepo,
		secretKey:    secretKey,
		tokenTimeout: tokenTimeout,
	}
}

func (s *AuthService) Register(email, password string) (string, error) {
	const op = "AuthService.Register"
	log := s.log.With(slog.String("op", op))

	log.Info("registering user", slog.String("email", email))

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to hash password", slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	user := &models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		log.Error("failed to create user", slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("user registered successfully", slog.String("email", email))

	return s.generateToken(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	const op = "AuthService.Login"
	log := s.log.With(slog.String("op", op))

	log.Info("logging in user", slog.String("email", email))

	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		log.Error("failed to get user", slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		log.Error("invalid password", slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: invalid credentials", op)
	}

	log.Info("user logged in successfully", slog.String("email", email))
	return s.generateToken(user)
}

func (s *AuthService) GetSecretKey() string {
	return s.secretKey
}

func (s *AuthService) generateToken(user *models.User) (string, error) {
	const op = "AuthService.generateToken"
	log := s.log.With(slog.String("op", op))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(s.tokenTimeout).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Error("failed to sign token", slog.String("error", err.Error()))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return tokenString, nil
}
