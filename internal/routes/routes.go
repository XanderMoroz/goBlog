package routes

import (
	"github.com/XanderMoroz/goBlog/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetUpRoutes sets up all the routes for the application
func SetupAuthRoutes(app *fiber.App) {
	app.Post("/api/v1/register", controllers.Register)
	app.Post("/api/v1/login", controllers.Login)
	app.Get("/api/v1/current_user", controllers.GetCurrentUser)
}
