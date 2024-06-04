package utils

import (
	"log"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
)

// Извлекаем все категории и авторов
func GetCategoriesFromDB() []models.Category {

	db := database.DB
	var categories []models.Category // article slice

	result := db.Preload("Articles").Find(&categories)

	// 	var languages []Language
	//   err := db.Model(&Language{}).Preload("Users").Find(&languages).Error

	if result.Error != nil {
		// handle error
		panic("failed to retrieve articles: " + result.Error.Error())
	}

	log.Println("Список статей — успешно извлечен:")
	for _, article := range categories {
		log.Printf("Category ID: <%d>, Title: <%s>\n", article.ID, article.Title)
	}
	return categories
}
