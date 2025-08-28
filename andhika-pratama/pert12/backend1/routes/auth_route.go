package routes

import (
	"pert12/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router) {
	api.Post("/register", controllers.Signup)
	api.Post("/login", controllers.Login)
	// api.Post("/logout", controllers.Logout)
}