package routes

import (
	"backend2/controllers"
	"backend2/middlewares"

	"github.com/gofiber/fiber/v2"
)

func userRoutes(api fiber.Router) {
	api.Get("/me", middlewares.AuthMiddleware(), controllers.GetMe)
	api.Patch("/profile/:id", middlewares.AuthMiddleware(), controllers.UpdateProfile)
}
