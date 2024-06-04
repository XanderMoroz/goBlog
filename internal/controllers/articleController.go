package controllers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
	"github.com/XanderMoroz/goBlog/internal/utils"
)

// @Summary		get all articles
// @Description Get all articles from db
// @Tags 		Articles
// @ID			get-all-articles
// @Produce		json
// @Success		200		{object}	[]models.ArticleResponse
// @Router		/articles [get]
func GetAllArticles(c *fiber.Ctx) error {

	articles := utils.GetArticlesFromDB()

	if len(articles) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "articles not found",
			"data":    nil,
		})
	}

	c.Status(http.StatusOK)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Articles Found",
		"data":    articles,
	})
}

// @Summary		get an article by ID
// @Description Get an article by ID
// @Tags 		Articles
// @ID			get-article-by-id
// @Produce		json
// @Param		id				path		string					true	"Article ID"
// @Success		200				{object}	models.ArticleResponse
// @Failure		404				{object}	[]string
// @Router		/articles/{id} 	[get]
func GetArticleById(c *fiber.Ctx) error {

	db := database.DB
	var article models.Article

	// Read the param userUUID
	id := c.Params("id")

	// Retrieve the record you want to update
	// result := db.First(&article, "ID = ?", id)
	result := db.Preload("User").First(&article, "ID = ?", id)

	if result.Error != nil {
		panic("failed to retrieve article: " + result.Error.Error())
	}

	// If no such note present return an error
	if article.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Article not found",
			"data":    nil,
		})
	}

	log.Println("Пользователь — успешно извлечен:")
	log.Printf("	ID: <%d>\n", article.ID)
	log.Printf("	Имя: <%s>\n", article.Title)
	log.Printf("	E-mail: <%s>\n", article.Content)

	// Return the note with the Id
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Article Found",
		"data":    article,
	})
}

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

// @Summary			update article by ID
// @Description 	Update article by ID
// @ID				delete-article-by-id
// @Tags 			Articles
// @Produce			json
// @Param			id					path		int								true	"articleID"
// @Param           request         	body        models.UpdateArticleBody    	true    "Введите новые данные статьи"
// @Success			200	{object}	[]string
// @Failure			404	{object}	[]string
// @Router			/articles/{id} 	[put]
func UpdateMyArticleById(c *fiber.Ctx) error {

	db := database.DB

	// Извлекаем JWT токен из куки пользователя
	cookieWithJWT := c.Cookies("jwt")
	body := new(models.UpdateArticleBody)

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

	// Read the param articleID
	id := c.Params("id")

	var article models.Article
	// Извлекаем запись по ID
	result = db.First(&article, "ID = ?", id)

	if result.Error != nil {
		panic("failed to retrieve article: " + result.Error.Error())
	}

	// If no such note present return an error
	if article.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Article not found",
			"data":    nil,
		})
	}

	// Если пользователь не является автором статьи
	if article.UserID != user.ID {
		log.Println(article.UserID)
		log.Println(user.ID)
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "You can update just your oun articles",
			"data":    nil,
		})
	}

	// Извлекаем тело запроса
	err = c.BodyParser(body)
	if err != nil {
		// Обрабатываем ошибку
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Проверьте данные",
			"data":    err,
		})
	}
	log.Println("Тело запроса извлечено:")
	log.Printf("	Новое Название статьи: <%s>\n", body.Title)
	log.Printf("	Новый Текст статьи: <%s>\n", body.Content)

	article.Title = body.Title
	article.Content = body.Content

	// Сохраняем изменения в БД
	result = db.Save(&article)
	if result.Error != nil {
		panic("failed to update article: " + result.Error.Error())
	}

	log.Println("Статья — успешно обновлена:")
	log.Printf("	ID: <%d>\n", article.ID)
	log.Printf("	Название: <%s>\n", article.Title)
	log.Printf("	Текст: <%s>\n", article.Content)

	// Return success message
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Article Updated",
	})
}

// @Summary		delete a article by ID
// @Description Delete a article by ID
// @ID			delete-article-by-id
// @Tags 		Articles
// @Produce		json
// @Param		id				path		string		true	"articleID"
// @Success		200				{object}	[]string
// @Failure		404				{object}	[]string
// @Router		/articles/{id} 	[delete]
func DeleteMyArticleById(c *fiber.Ctx) error {

	db := database.DB

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

	// Извлекаем параметр articleID
	id := c.Params("id")

	var article models.Article

	// Извлекаем статью по ID
	result = db.First(&article, "ID = ?", id)

	if result.Error != nil {
		panic("failed to retrieve user: " + result.Error.Error())
	}

	// Если статья не найдена
	if article.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Article Not Found",
		})
	}

	// Если пользователь не является автором статьи
	if article.UserID != user.ID {
		log.Println(article.UserID)
		log.Println(user.ID)
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "You can update just your oun articles",
		})
	}

	// Удаляем статью по ID (с извлечением ошибки)
	err = db.Delete(&article, "ID = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete article",
		})
	}

	log.Println("Статья — успешно удалена:")
	log.Printf("	ID: <%d>\n", article.ID)

	// Return success message
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Article Deleted",
	})
}
