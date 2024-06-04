package models

import (
	"time"

	"gorm.io/gorm"
)

// Структура статьи
type Category struct {
	gorm.Model
	ID        uint64    `gorm:"primary_key;auto_increment" json:"ID"` // Уникальный идентификатор
	Title     string    `gorm:"size:255;not null" json:"title"`       // Название категории
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	// Список статей связанных с категорией
	Articles []*Article `gorm:"many2many:article_categories;"`
}

// CreateCategoryBody
// @Description Тело запроса для создания статьи
type CreateCategoryBody struct {
	Title string `json:"title" validate:"required"`
}

// CategoryResponse
// @Description Тело ответа после извлечения статьи
type CategoryResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

// CreateCategoryBody
// @Description Тело запроса для создания статьи
type AddArticleToCategoryBody struct {
	ArticleID    string `json:"article_id" validate:"required"`
	CategoryName string `json:"category_name" validate:"required"`
}
