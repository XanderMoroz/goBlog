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

	result := db.Preload("Categories").Preload("User").Find(&articles)

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

// Извлекаем статью по ID
func GetArticleByIDFromDB(articleID string) models.Article {

	db := database.DB
	var article models.Article

	// Извлекаем статью вместе с автором и категориями
	result := db.Preload("Categories").Preload("User").First(&article, "ID = ?", articleID)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve article: " + result.Error.Error())
	}

	if article.ID == 0 {
		// handle error
		panic("failed to retrieve article: " + result.Error.Error())
	} else {
		// handle success
		log.Println("Cтатья — успешно извлечена:")
		log.Printf("	ID: <%d>\n", article.ID)
		log.Printf("	Название: <%s>\n", article.Title)
		log.Printf("	Содержание: <%s>\n", article.Content)
		log.Printf("	Автор: <%s>\n", article.User.Name)
		log.Printf("	Категории: <%v>\n", article.Categories)
	}

	return article
}

// Создаем новую статью
func CreateArticleInDB(newArticle models.Article) models.Article {

	db := database.DB

	// ... Create a new article record...
	result := db.Create(&newArticle)
	if result.Error != nil {
		panic("failed to create article: " + result.Error.Error())
	} else {
		// ... Handle successful creation ...
		log.Println("Новая статья — успешно создана:")
		log.Printf("	ID: <%d>\n", newArticle.ID)
		log.Printf("	Название: <%s>\n", newArticle.Title)
		log.Printf("	Текст: <%s>\n", newArticle.Content)
		log.Printf("	Автор: <%s>\n", newArticle.User.Name)
	}
	return newArticle
}
