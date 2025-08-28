package routes

import (
	"pert12/controllers"
	"pert12/middlewares"

	"github.com/gofiber/fiber/v2"
)

func EventRoutes(api fiber.Router) {
	api.Get("/event", controllers.GetEvents)
	api.Post("/event", middlewares.AuthMiddleware(), middlewares.AdminMiddleware, controllers.CreateEvent)
	// api.Patch("/event/:id", middlewares.AdminAuth, controllers.UpdateEvent)
	// api.Delete("/event/:id", middlewares.AdminAuth, controllers.DeleteEvent)
}