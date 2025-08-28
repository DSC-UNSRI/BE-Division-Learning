package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func EventRoutes(api fiber.Router){
	event := api.Group("/event")
	event.Get("/", controllers.GetAllEvents)
	event.Post("/", middleware.Protected(), middleware.IsAdmin(), controllers.CreateEvent)
	event.Patch("/:id", middleware.Protected(),middleware.IsAdmin(), controllers.UpdateEvent)
	event.Delete("/:id", middleware.Protected(),middleware.IsAdmin(), controllers.DeleteEvent)
}