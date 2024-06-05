package controllers

import (
	"log"
	// "net/http"

	"github.com/gofiber/fiber/v2"

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
// @Router         /comment 			[post]
func CreateNewCommentToArticle(c *fiber.Ctx) error {

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

	// Извлекаем JWT токен из куки пользователя
	cookieWithJWT := c.Cookies("jwt")

	log.Println("Извлекаем ID пользователя из JWT токена")
	userID, err := utils.ParseUserIDFromJWTToken(cookieWithJWT)

	if err != nil {
		log.Println("Ошибка:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "Please Log in",
		})
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

	// ... Создаем новую статью...
	result = db.Create(&newArticle)
	if result.Error != nil {
		// ... В случае ошибки ...
		panic("failed to create article: " + result.Error.Error())
	} else {
		// ... В случае успеха ...
		log.Println("Новая статья — успешно создана:")
		log.Printf("	ID: <%d>\n", newArticle.ID)
		log.Printf("	Название: <%s>\n", newArticle.Title)
		log.Printf("	Текст: <%s>\n", newArticle.Content)
		log.Printf("	Автор: <%s>\n", newArticle.User.Name)
	}

	// Возвращаем статью
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Article Created",
		"data":    newArticle,
	})
}
