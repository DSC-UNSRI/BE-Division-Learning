package routes

import (
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router) {
	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.SignUp)
	api.Post("/logout", controllers.Logout)
}