package routes

import (
	"pert12/controllers"

	"github.com/gofiber/fiber/v2"
)

func routeEvent(api fiber.Router) {
	event := api.Group("/event")
	event.Post("/", controllers.PostEvent)
	event.Get("/", controllers.GetEvent)
	event.Put("/:id", controllers.UpdateEvent)
	event.Delete("/:id", controllers.DeleteEvent)
}