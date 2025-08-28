package routes

import (
	"backend2/controllers"

	"github.com/gofiber/fiber/v2"
)

func userRoutes(api fiber.Router) {
	user := api.Group("/user")
	user.Post("/signup", controllers.SignUp)
	user.Post("/login", controllers.Login)
	user.Post("/logout", controllers.Logout)
}
