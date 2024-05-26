package routes

import (
	"github.com/XanderMoroz/goBlog/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetUpRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
