package routes

import (
	"pert12/controllers"

	"github.com/gofiber/fiber/v2"
)

func routeUser(api fiber.Router) {
	user := api.Group("/user")
	user.Get("/:id", controllers.GetMeByID)
	user.Put("/:id", controllers.UpdateProfile)
}
