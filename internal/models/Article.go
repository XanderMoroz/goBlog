package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Структура статьи
type Article struct {
	gorm.Model
	ID        uint64    `gorm:"primary_key;auto_increment" json:"ID"` // Уникальный идентификатор
	User      User      `json:"author"`                               // Автор статьи
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"author_id"`  // Уникальный идентификатор автора статьи
	Title     string    `gorm:"size:255;not null" json:"title"`       // Название статьи
	Content   string    `gorm:"size:255;not null;" json:"content"`    // Содержание статьи
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	// Категории к которым принадлежит статья
	Categories []*Category `gorm:"many2many:article_categories;"`
	Comments   []Comment   `gorm:"foreignKey:ArticleID"`
}

// CreateArticleRequest
// @Description Тело запроса для создания статьи
type CreateArticleRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// ArticleResponse
// @Description Тело ответа после cоздания статьи
type ArticleResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

// UpdateArticleBody
// @Description Тело запроса для обновления статьи
type UpdateArticleBody struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
