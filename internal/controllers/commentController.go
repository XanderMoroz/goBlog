package controllers

import (
	"log"
	// "net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
	"github.com/XanderMoroz/goBlog/internal/utils"
)

// @Summary        	create new comment to article
// @Description    	Creating new Comment to Article in DB with given request body
// @Tags           	Articles
// @Accept         	json
// @Produce        	json
// @Param			id								path		string							true	"Article ID"
// @Param          	request         				body        models.CreateCommentRequest    	true    "Введите текст комментария"
// @Success       	201              				{string}    string
// @Failure        	400              				{string}    string    						"Bad Request"
// @Router         	/article/{id}/add_comment 		[post]
func AddNewCommentToArticle(c *fiber.Ctx) error {

	db := database.DB

	body := new(models.CreateCommentRequest)

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

	// var user models.User
	log.Println("Извлекаем пользователя по ID...")
	user := utils.GetUserByIDFromDB(userID)

	// Read the param id
	articleID := c.Params("id")

	article := utils.GetArticleByIDFromDB(articleID)
	// result := db.Where("ID =?", userID).First(&user)

	// if result.Error != nil {
	// 	panic("failed to retrieve user: " + result.Error.Error())
	// } else {
	// 	log.Println("Пользователь — успешно извлечен:")
	// 	log.Printf("	ID: <%s>\n", user.ID)
	// 	log.Printf("	Имя: <%s>\n", user.Name)
	// 	log.Printf("	E-mail: <%s>\n", user.Email)
	// }

	newComment := models.Comment{
		Content:   body.Content,
		User:      user,
		UserID:    user.ID,
		Article:   article,
		ArticleID: article.ID,
	}

	// ... Создаем новый комментарий...
	result := db.Create(&newComment)
	if result.Error != nil {
		// ... В случае ошибки ...
		panic("failed to create article: " + result.Error.Error())
	} else {
		// ... В случае успеха ...
		log.Println("Новый комментарий — успешно создан:")
		log.Printf("	ID: <%d>\n", newComment.ID)
		log.Printf("	Текст: <%s>\n", newComment.Content)
		log.Printf("	Имя автора: <%s>\n", newComment.User.Name)
		log.Printf("	Статья ID: <%d>\n", newComment.Article.ID)
	}

	// Возвращаем статью
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Article Created",
		"data":    newComment,
	})
}
