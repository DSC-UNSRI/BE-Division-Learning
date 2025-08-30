package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/controllers"
	"backend/middlewares"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/event", controllers.GetEvents)
	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.Signup)
	api.Post("/logout", controllers.Logout)

	protected := api.Group("/", middlewares.AuthMiddleware())
	
	protected.Get("/me", controllers.GetMe)
	protected.Patch("/profile/:id", controllers.UpdateProfile)

	protected.Post("/event", controllers.PostEvent)
	protected.Patch("/event/:id", middlewares.AdminMiddleware(), controllers.UpdateEvent)
	protected.Delete("/event/:id", middlewares.AdminMiddleware(), controllers.DeleteEvent)
}