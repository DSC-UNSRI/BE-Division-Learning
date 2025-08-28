package routes

import "github.com/gofiber/fiber/v2"

func RoutesList(app *fiber.App) {
	api := app.Group("/api")
	authRoutes(api)
	eventRoutes(api)
	userRoutes(api)
}