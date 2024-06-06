package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Структура комментария
type Comment struct {
	gorm.Model
	ID uint64 `gorm:"primary_key;auto_increment" json:"ID"`
	// Содержание статьи
	Content   string    `gorm:"size:1024;not null;" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	// Автор комментария
	User   User      `json:"author"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"author_id"`
	// Прокоментированная статья
	Article   Article `json:"article"`
	ArticleID uint64  `gorm:"type:uint64;not null" json:"article_id"`
}

// CreateCommentRequest
// @Description Тело запроса для создания статьи
type CreateCommentRequest struct {
	Content string `json:"content" validate:"required"`
}
