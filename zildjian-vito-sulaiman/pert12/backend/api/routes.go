package api

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Post("/logout", controllers.Logout)

	event := api.Group("/event", middleware.Authenticated)
	event.Get("/", controllers.GetEvents)
	event.Post("/", middleware.Admin, controllers.CreateEvent)
	event.Patch("/:id", middleware.Admin, controllers.UpdateEvent)
	event.Delete("/:id", middleware.Admin, controllers.DeleteEvent)

	user := api.Group("/user", middleware.Authenticated)
	user.Get("/me", controllers.GetMe)
	user.Patch("/profile/:id", middleware.ProfileUpdate, controllers.UpdateProfile)
}
