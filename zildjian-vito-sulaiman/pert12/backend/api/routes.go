package api

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Post("/logout", controllers.Logout)

	event := api.Group("/event")
	event.Get("/", controllers.GetEvents)
	event.Post("/", controllers.CreateEvent)
	event.Patch("/:id", controllers.UpdateEvent)
	event.Delete("/:id", controllers.DeleteEvent)

	user := api.Group("/user")
	user.Get("/me", controllers.GetMe)
	user.Patch("/profile/:id", controllers.UpdateProfile)
}
