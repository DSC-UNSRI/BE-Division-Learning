package routes

import (
	"pert12/controllers"
	"pert12/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	api.Get("/me", middlewares.AuthMiddleware(), controllers.GetMe)
	// api.Patch("/profile/:id", middlewares.AuthMiddleware, controllers.UpdateProfile)
}