package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	// ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"uniqueIndex"`
	Password  []byte    `json:"-"` // contain the hashed password.
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

// LoginRequest
// @Description Тело запроса для аутентификации пользователя
type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// SignUpUserRequest
// @Description Тело запроса для регистрации пользователя
type SignUpUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// CreateUserRequest
// @Description Тело запроса для создания пользователя
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// CreateUserRequest
// @Description Тело ответа после оздания пользователя
type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

// CreateUserRequest
// @Description Тело запроса для обновления пользователя
type UpdateUserBody struct {
	Name  string `json:"firstName"`
	Email string `json:"email"`
}
