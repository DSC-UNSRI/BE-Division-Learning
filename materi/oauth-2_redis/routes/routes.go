package routes

import (
	"oauth-2_redis/controllers"
	"oauth-2_redis/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(api fiber.Router) {
	api.Post("/login", controllers.Login)
	api.Post("/refresh", controllers.RefreshToken)
	api.Post("/register", controllers.CreateUser)
	api.Get("/event", controllers.GetEvent)

	api.Use(middleware.Protected())
	api.Post("/logout", controllers.Logout)
	api.Post("/event", controllers.CreateEvent)
	api.Patch("/event/:id", controllers.UpdateEvent)
	api.Delete("/event/:id", controllers.DeleteEvent)
	api.Get("/me", controllers.Me)
	api.Patch("/profile/:id", controllers.UpdateProfile)
}
