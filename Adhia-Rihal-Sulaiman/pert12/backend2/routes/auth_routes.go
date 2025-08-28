package routes

import (
	"backend2/controllers"
	"backend2/middlewares"

	"github.com/gofiber/fiber/v2"
)

func userRoutes(api fiber.Router) {
	user := api.Group("/user")
	user.Post("/signup", controllers.SignUp)
	user.Post("/login", controllers.Login)
	user.Post("/logout", middlewares.Auth, controllers.Logout)
}
