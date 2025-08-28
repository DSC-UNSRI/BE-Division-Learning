package routes

import (
	"backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("/api", middleware.Protected()) 

	EventRoutes(api)
	UserRoutes(api)
}