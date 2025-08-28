package routes

import (
	"backend2/controllers"
	"backend2/middlewares"

	"github.com/gofiber/fiber/v2"
)

func eventRoutes(api fiber.Router) {
	api.Get("/event", controllers.GetEvent)
	api.Post("/event", controllers.PostEvent, middlewares.AdminMiddleware(), middlewares.AuthMiddleware())
	api.Patch("/event/:id", controllers.UpdateEvent, middlewares.AdminMiddleware(), middlewares.AuthMiddleware())
	api.Delete("/event/:id", controllers.DeleteEvent, middlewares.AdminMiddleware(), middlewares.AuthMiddleware())
}
