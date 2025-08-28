package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func EventRoutes(api fiber.Router){
	event := api.Group("/event")
	event.Get("/", controllers.GetAllEvents)
	event.Post("/", middleware.IsAdmin(), controllers.CreateEvent)
	event.Patch("/:id",middleware.IsAdmin(), controllers.UpdateEvent)
	event.Delete("/:id",middleware.IsAdmin(), controllers.DeleteEvent)
}