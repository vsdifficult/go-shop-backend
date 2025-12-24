package repositories

import (
	"goshop/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	user.ID = uuid.New()

	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User

	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}
