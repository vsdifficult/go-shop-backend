package services

import (
	"goshop/internal/models"
	"goshop/internal/repositories"

	"github.com/google/uuid"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id uuid.UUID) (*models.UserDto, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &models.UserDto{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Verified: user.Verified,
	}, nil
}
