package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("/api") 

	AuthRoutes(api)
	EventRoutes(api)
	UserRoutes(api)
}