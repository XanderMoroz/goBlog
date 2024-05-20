package models

import (
	// "errors"
	// "html"
	// "strings"
	"time"
	// "github.com/jinzhu/gorm"
)

// Структура статьи
type Article struct {
	ID      uint64 `gorm:"primary_key;auto_increment" json:"id"` // Уникальный идентификатор
	Title   string `gorm:"size:255;not null" json:"title"`       // Название статьи
	Content string `gorm:"size:255;not null;" json:"content"`    // Содержание статьи
	// Author    User      `json:"author"`                                      // Автор статьи
	// AuthorID  uint32    `gorm:"not null" json:"author_id"`                   // Уникальный идентификатор автора статьи
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // Дата создания
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"` // Дата изменения
}

// CreatePostRequest
// @Description Request about creating Post
type CreatePostRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
