package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Name         string    `gorm:"name"`
	Email        string    `gorm:"email"`
	PasswordHash string    `gorm:"password_hash"`
	Verified     bool      `gorm:"verified"`
	VerifiedCode string    `gorm:"verified_code"`
}

type UserCreateDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserVerifyDto struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type UserDto struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Verified bool      `json:"verified"`
}
