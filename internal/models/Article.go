package models

import (
	// "errors"
	// "html"
	// "strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Структура статьи
type Article struct {
	gorm.Model
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"` // Уникальный идентификатор
	User      User      `json:"author"`                               // Автор статьи
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"author_id"`  // Уникальный идентификатор автора статьи
	Title     string    `gorm:"size:255;not null" json:"title"`       // Название статьи
	Content   string    `gorm:"size:255;not null;" json:"content"`    // Содержание статьи
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

// CreateArticleRequest
// @Description Тело запроса для создания пользователя
type CreateArticleRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// UserResponse
// @Description Тело ответа после cоздания пользователя
type ArticleResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}
