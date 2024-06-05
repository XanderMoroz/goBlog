package controllers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/models"
	"github.com/XanderMoroz/goBlog/internal/utils"
)

// @Summary        create new category
// @Description    Creating Category in DB with given request body
// @Tags           Categories
// @Accept         json
// @Produce        json
// @Param          request         		body        models.CreateCategoryBody    true    "Введите название категории статьи"
// @Success        201              	{string}    string
// @Failure        400              	{string}    string    "Bad Request"
// @Router         /categories 			[post]
func CreateNewCategory(c *fiber.Ctx) error {

	db := database.DB

	body := new(models.CreateCategoryBody)

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

	newCategory := models.Category{
		Title: body.Title,
	}

	// ... Создаем новую категорию...
	result := db.Create(&newCategory)
	if result.Error != nil {
		// ... В случае ошибки ...
		panic("failed to create category: " + result.Error.Error())
	} else {
		// ... В случае успеха ...
		log.Println("Новая категория статья — успешно создана:")
		log.Printf("	ID: <%d>\n", newCategory.ID)
		log.Printf("	Название: <%s>\n", newCategory.Title)
	}

	// Возвращаем категорию
	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Category Created",
		"data":    newCategory,
	})
}

// @Summary		get all categories
// @Description Get all categories from db
// @Tags 		Categories
// @ID			get-all-categories
// @Produce		json
// @Success		200		{object}	[]models.CategoryResponse
// @Router		/categories [get]
func GetAllCategories(c *fiber.Ctx) error {

	categories := utils.GetCategoriesFromDB()

	if len(categories) == 0 {
		return c.Status(http.StatusNoContent).JSON(fiber.Map{
			"status":  "error",
			"message": "categories not found",
			"data":    nil,
		})
	}

	// c.Status(http.StatusOK)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Articles Found",
		"data":    categories,
	})
}

// @Summary        add article to category
// @Description    Adding Article to Category in DB with given request body
// @Tags           Categories
// @Accept         json
// @Produce        json
// @Param          request         			body        models.AddArticleToCategoryBody    true    "Введите ID статьи и название категории"
// @Success        201              		{string}    string
// @Failure        400              		{string}    string    "Bad Request"
// @Router         /categories/add_article 	[post]
func AddArticleToCategory(c *fiber.Ctx) error {

	db := database.DB

	body := new(models.AddArticleToCategoryBody)

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

	article := utils.GetArticleByIDFromDB(body.ArticleID)

	category := utils.GetCategoryByNameFromDB(body.CategoryName)

	db.Model(&category).Association("Articles").Append(&article)
	db.Model(&article).Association("Categories").Append(&category)

	// Возвращаем категорию
	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Article added to Category",
		// "data":    newCategory,
	})
}

// @Summary        delete article from category
// @Description    Deleting Article from Category in DB with given request body
// @Tags           Categories
// @Accept         json
// @Produce        json
// @Param          request         			body        models.AddArticleToCategoryBody    true    "Введите ID статьи и название категории"
// @Success        201              		{string}    string
// @Failure        400              		{string}    string    "Bad Request"
// @Router         /categories/remove_article 	[post]
func DeleteArticleFromCategory(c *fiber.Ctx) error {

	db := database.DB

	body := new(models.AddArticleToCategoryBody)

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

	article := utils.GetArticleByIDFromDB(body.ArticleID)

	category := utils.GetCategoryByNameFromDB(body.CategoryName)

	// result := db.Create(&newCategory)
	db.Model(&category).Association("Articles").Delete(&article)
	db.Model(&article).Association("Categories").Delete(&category)

	// Возвращаем категорию
	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Article delete from Category",
	})
}
