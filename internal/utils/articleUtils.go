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
	var article models.Article // article slice

	// Извлекаем статью вместе с автором и категориями
	result := db.Preload("Categories").Preload("User").First(&article, "ID = ?", articleID)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve article: " + result.Error.Error())
	}

	if article.ID == 0 {
		// handle error
		panic("failed to retrieve article: " + result.Error.Error())
	}

	log.Println("Cтатья — успешно извлечена:")
	log.Printf("	ID: <%d>\n", article.ID)
	log.Printf("	Название: <%s>\n", article.Title)
	log.Printf("	Содержание: <%s>\n", article.Content)
	log.Printf("	Автор: <%s>\n", article.User.Name)
	log.Printf("	Категории: <%v>\n", article.Categories)

	return article
}

// Создаем новую статью
func CreateArticleInDB() {

	db := database.DB

	var user models.User

	// Retrieve the record you want to update
	result := db.First(&user, "ID = ?", "31e5665a-6930-4fda-97e0-ed24d3a561a6")
	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}
	// iterate over the users slice and print the details of each user
	log.Println("Пользователь — успешно извлечен:")
	log.Printf("	ID: <%s>\n", user.ID)
	log.Printf("	Имя: <%s>\n", user.Name)
	log.Printf("	E-mail: <%s>\n", user.Email)

	newArticle := models.Article{
		Title:   "Jane",
		Content: "Doe",
		User:    user,
		UserID:  user.ID,
	}

	// ... Create a new user record...
	result = db.Create(&newArticle)
	if result.Error != nil {
		panic("failed to create article: " + result.Error.Error())
	}
	// ... Handle successful creation ...
	log.Println("Новая статья — успешно создана:")
	log.Printf("	ID: <%d>\n", newArticle.ID)
	log.Printf("	Название: <%s>\n", newArticle.Title)
	log.Printf("	Текст: <%s>\n", newArticle.Content)
	log.Printf("	Автор: <%s>\n", newArticle.User.Name)
}
