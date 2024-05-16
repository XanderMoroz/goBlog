package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/swagger"

	_ "github.com/XanderMoroz/goBlog/docs"
)

// @title Good News on Go
// @version 1.0
// @description Сервис с новостными статьями и блогами.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:3000

func main() {

	// Start a new fiber app
	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	app.Get("/article", getAllArticles)
	app.Get("/article/:id", getArticleByID)
	app.Post("/article", createArticle)
	app.Delete("/article/:id", deleteArticle)

	// Start Server and Listen on PORT 3000
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /nice [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

// @Summary		get all items
// @Description Get all items in the article list
// @Tags 		Articles
// @ID			get-all-articles
// @Produce		json
// @Success		200		{object}	article
// @Router		/article [get]
func getAllArticles(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	return c.JSON(articleList)
}

// @Summary		get a article item by ID
// @Description Get a article item by ID
// @Tags 		Articles
// @ID			get-article-by-id
// @Produce		json
// @Param		id	path		string	true	"article ID"
// @Success		200	{object}	article
// @Failure		404	{object}	message
// @Router		/article/{id} [get]
func getArticleByID(c *fiber.Ctx) error {
	ID := c.Params("id")

	// loop through articleList and return item with matching ID
	for _, todo := range articleList {
		if todo.ID == ID {
			c.Status(http.StatusOK)
			return c.JSON(todo)
		}
	}

	// return error message if article is not found
	r := message{"article not found"}
	c.Status(http.StatusNotFound)
	return c.JSON(r)
}

// @Summary		add a new item
// @Description Add a new item to the article list
// @ID			create-article
// @Tags 		Articles
// @Produce		json
// @Param		data	body		article	true	"article data"
// @Success		200		{object}	article
// @Failure		400		{object}	message
// @Router		/article [post]
func createArticle(c *fiber.Ctx) error {
	var newArticle article

	// bind the received JSON data to newArticle
	if err := c.BodyParser(&newArticle); err != nil {
		r := message{"an error occurred while creating article"}
		c.Status(http.StatusBadRequest)
		return c.JSON(r)
	}

	// add the new article item to articleList
	articleList = append(articleList, newArticle)
	c.Status(http.StatusCreated)
	return c.JSON(newArticle)
}

// @Summary		delete a article item by ID
// @Description Delete a article item by ID
// @ID			delete-article-by-id
// @Tags 		Articles
// @Produce		json
// @Param		id	path		string	true	"article ID"
// @Success		200	{object}	article
// @Failure		404	{object}	message
// @Router		/article/{id} [delete]
func deleteArticle(c *fiber.Ctx) error {
	ID := c.Params("id")

	// loop through articleList and delete item with matching ID
	for index, article := range articleList {
		if article.ID == ID {
			articleList = append(articleList[:index], articleList[index+1:]...)
			r := message{"successfully deleted todo"}
			c.Status(http.StatusOK)
			return c.JSON(r)
		}
	}

	// return error message if article is not found
	r := message{"article not found"}
	c.Status(http.StatusNotFound)
	return c.JSON(r)
}

// article represents data about a task in the article list
type article struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

// message represents request response with a message
type message struct {
	Message string `json:"message"`
}

// todo slice to seed article list data
var articleList = []article{
	{"1", "Learn Go"},
	{"2", "Build an API with Go"},
	{"3", "Document the API with swag"},
}
