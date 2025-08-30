package routes

import (
	"pert12/controllers"
	"pert12/middleware"

	"github.com/gofiber/fiber/v2"
)

func routeEvent(api fiber.Router) {
	event := api.Group("/event")
	event.Post("/", middleware.JWTToken(), middleware.AdminOnly(), controllers.PostEvent)
	event.Get("/", controllers.GetEvent)
	event.Put("/:id", middleware.JWTToken(), middleware.AdminOnly(), controllers.UpdateEvent)
	event.Delete("/:id",  middleware.JWTToken(), middleware.AdminOnly(), controllers.DeleteEvent)
}
