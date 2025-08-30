package routes

import (
	"tugas12/controllers"
	"tugas12/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", controllers.Signup)
	api.Post("/login", controllers.Login)
	api.Post("/logout", controllers.Logout)

	api.Get("/me", middleware.JWTProtected(""), controllers.GetMe)
	api.Patch("/profile/:id", middleware.JWTProtected(""), controllers.UpdateProfile)

	api.Get("/event", controllers.GetEvents)
	api.Get("/event/:id", controllers.GetEvent)
	api.Post("/event", middleware.JWTProtected("admin"), controllers.CreateEvent)
	api.Patch("/event/:id", middleware.JWTProtected("admin"), controllers.UpdateEvent)
	api.Delete("/event/:id", middleware.JWTProtected("admin"), controllers.DeleteEvent)
}
