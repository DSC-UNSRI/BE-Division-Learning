package routes

import (
	"pert12/controllers"

	"github.com/gofiber/fiber/v2"
)

func routeAuth(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/signup", controllers.SignUp)
	auth.Post("/login", controllers.Login)
	auth.Post("/logout", controllers.Logout)
}
