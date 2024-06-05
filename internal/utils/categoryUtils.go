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

	result := db.Find(&categories)

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

// Извлекаем категорию по названию
func GetCategoryByNameFromDB(name string) models.Category {

	db := database.DB
	var category models.Category // article slice

	// Retrieve the record you want to update
	// result := db.First(&article, "ID = ?", id)
	result := db.Preload("Articles").First(&category, "Title = ?", name)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve category: " + result.Error.Error())
	}

	if category.ID == 0 {
		// handle error
		panic("failed to retrieve category: " + result.Error.Error())
	}

	log.Println("Категория — успешно извлечена:")
	log.Printf("Category ID: <%d>, Title: <%s>\n", category.ID, category.Title)

	return category
}

// // Извлекаем категорию по названию
// func AddArticleToCategoryInDB() models.Category {

// 	name := "Business"
// 	db := database.DB
// 	var category models.Category // category slice

// 	// Retrieve the record you want to update
// 	// result := db.First(&article, "ID = ?", id)
// 	result := db.Model(&category).Association("Languages").Append(&Language{Name: "DE"})

// 	if result.Error != nil {
// 		// handle error
// 		panic("failed to retrieve category: " + result.Error.Error())
// 	}

// 	if category.ID == 0 {
// 		// handle error
// 		panic("failed to retrieve category: " + result.Error.Error())
// 	}

// 	log.Println("Категория — успешно извлечена:")
// 	log.Printf("Category ID: <%d>, Title: <%s>\n", category.ID, category.Title)

// 	return category
// }
