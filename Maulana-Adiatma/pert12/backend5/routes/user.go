package routes

import (
	"pert12/controllers"
	"pert12/middleware"

	"github.com/gofiber/fiber/v2"
)

func routeUser(api fiber.Router) {
    user := api.Group("/user")
    user.Get("/me", middleware.JWTToken(), controllers.GetMe)
    user.Put("/:id", middleware.JWTToken(), controllers.UpdateProfile)
}