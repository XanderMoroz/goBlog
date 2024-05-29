package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/XanderMoroz/goBlog/database"
	"github.com/XanderMoroz/goBlog/internal/controllers"
	"github.com/XanderMoroz/goBlog/internal/middlewares"

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

	// controllers.CreateArticleInDB()

	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP! Go to:\nhttp://127.0.0.1:8080/swagger/index.html")
		return err
	})

	// Auth routes
	app.Post("/api/v1/register", controllers.Register)
	app.Post("/api/v1/login", controllers.Login)
	app.Get("/api/v1/current_user", middlewares.JWTProtected(), controllers.GetCurrentUser)
	app.Get("/api/v1/logout", controllers.Logout)

	// User routes
	app.Get("/users", controllers.GetAllUsers)
	app.Post("/users", controllers.AddNewUser)
	app.Get("/users/:id", controllers.GetUserById)
	app.Put("/users/:id", controllers.UpdateUserById)
	app.Delete("/users/:id", controllers.DeleteUserById)

	// User routes
	app.Post("/articles", controllers.CreateMyArticle)

	// Start Server and Listen on PORT 8080
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
