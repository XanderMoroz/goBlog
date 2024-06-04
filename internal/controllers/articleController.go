package controllers

import (
	"log"
	// "net/http"

	"github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
	"github.com/XanderMoroz/goBlog/internal/utils"
)

// @Summary        create new article
// @Description    Creating Article in DB with given request body
// @Tags           Articles
// @Accept         json
// @Produce        json
// @Param          request         	body        models.CreateArticleRequest    true    "Введите данные статьи"
// @Success        201              {string}    string
// @Failure        400              {string}    string    "Bad Request"
// @Router         /articles 			[post]
func CreateMyArticle(c *fiber.Ctx) error {
	db := database.DB

	body := new(models.CreateArticleRequest)

	// Извлекаем тело запроса
	err := c.BodyParser(body)
	if err != nil {
		// Обрабатываем ошибку
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Проверьте данные",
			"data":    err,
		})
	}
	log.Println("Запрос успешно обработан обработан")

	//Извлекаем JWT токен из куки пользователя
	cookieWithJWT := c.Cookies("jwt")

	log.Println("Извлекаем ID пользователя по из JWT токена")
	userID, err := utils.ParseUserIDFromJWTToken(cookieWithJWT)

	if err != nil {
		log.Println("Ошибка:", err)
	} else {
		log.Println("USER_ID из токена:", userID)
	}

	var user models.User
	log.Println("Извлекаем пользователя по ID...")
	result := db.Where("ID =?", userID).First(&user)

	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	} else {
		log.Println("Пользователь — успешно извлечен:")
		log.Printf("	ID: <%s>\n", user.ID)
		log.Printf("	Имя: <%s>\n", user.Name)
		log.Printf("	E-mail: <%s>\n", user.Email)
	}

	newArticle := models.Article{
		Title:   body.Title,
		Content: body.Content,
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

	// Return the created note
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Article Created",
		"data":    newArticle,
	})
}

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
