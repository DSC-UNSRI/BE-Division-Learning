package routes

import (
	"backend/middleware"
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	api.Get("/me", middleware.Protected(), controllers.GetMe)
	api.Patch("/profile/:id",  middleware.Protected(), controllers.UpdateProfile)
}