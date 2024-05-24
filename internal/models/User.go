package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"uniqueIndex"`
	Country   string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	Age       int       `gorm:"not null;size:3"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

// CreateUserRequest
// @Description Тело запроса для создания пользователя
type CreateUserRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"LastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Country   string `json:"country" validate:"required"`
	Role      string `json:"role" validate:"required"`
	Age       int    `json:"age" validate:"required"`
}

// CreateUserRequest
// @Description Тело запроса для создания пользователя
type UserResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Country   string    `json:"country" validate:"required"`
	Role      string    `json:"role" validate:"required"`
	Age       int       `json:"age" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

// CreateUserRequest
// @Description Тело запроса для обновления пользователя
type UpdateUserBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	// Role      string `json:"role"`
	// Age       int    `json:"age"`
}
