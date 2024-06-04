package utils

import (
	"log"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
)

// Извлекаем все статьи и авторов
func GetArticlesFromDB() []models.Article {

	db := database.DB
	var articles []models.Article // article slice

	result := db.Preload("User").Find(&articles)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve articles: " + result.Error.Error())
	}

	log.Println("Список статей — успешно извлечен:")
	for _, article := range articles {
		log.Printf("Article ID: %d, Title: %s, Content: %s Author: <%s>\n", article.ID, article.Title, article.Content, article.User.Name)
	}
	return articles
}
