package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/controllers"

	// "github.com/XanderMoroz/goBlog/internal/middlewares"

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

// @host 127.0.0.1:8080

func main() {

	log.Println("Подключаюсь к базе данных")
	database.Connect()

	// Start a new fiber app
	log.Println("Инициализируем приложение Fiber")
	app := fiber.New(fiber.Config{
		ServerHeader:  "Good News on Go",
		AppName:       "Good News on Go v0.0.1",
		CaseSensitive: true,
		StrictRouting: false,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())

	// utils.GetCategoryByNameFromDB()
	// utils.GetArticleByIDFromDB()

	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.Status(200).JSON(fiber.Map{
			"message": "API APP is UP!",
			"docs":    "http://127.0.0.1:8080/swagger/index.html",
		})
		return err
	})

	// Auth routes
	app.Post("/api/v1/register", controllers.Register)
	app.Post("/api/v1/login", controllers.Login)
	app.Get("/api/v1/current_user", controllers.GetCurrentUser)
	app.Get("/api/v1/logout", controllers.Logout)

	// Article routes
	app.Get("/articles", controllers.GetAllArticles)
	app.Post("/articles", controllers.CreateMyArticle)
	app.Get("/articles/:id", controllers.GetArticleById)
	app.Put("/articles/:id", controllers.UpdateMyArticleById)
	app.Delete("/articles/:id", controllers.DeleteMyArticleById)

	// Category routes
	app.Post("/categories", controllers.CreateNewCategory)
	app.Get("/categories", controllers.GetAllCategories)
	app.Post("/categories/add_article", controllers.AddArticleToCategory)
	app.Post("/categories/remove_article", controllers.DeleteArticleFromCategory)

	// Start Server and Listen on PORT 8080
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
