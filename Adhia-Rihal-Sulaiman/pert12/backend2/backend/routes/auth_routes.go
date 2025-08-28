package routes

import (
	"backend2/controllers"

	"github.com/gofiber/fiber/v2"
)

func authRoutes(api fiber.Router) {
	api.Post("/register", controllers.SignUp)
	api.Post("/login", controllers.Login)
	api.Post("/logout", controllers.Logout)
}
