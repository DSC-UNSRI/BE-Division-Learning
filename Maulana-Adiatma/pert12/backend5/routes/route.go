package routes

import "github.com/gofiber/fiber/v2"

func MainRoutes(app *fiber.App) {
	api := app.Group("/api")
	routeEvent(api)
	routeAuth(api)
	routeUser(api)
}
