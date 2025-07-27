package routes

import (
	"pert11/controllers"

	"github.com/gofiber/fiber/v2"
)

func Login(app *fiber.App) {
	app.Post("/login", controllers.Login)
}
