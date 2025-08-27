package routes

import (
	"be_pertemuan10/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/user", controllers.GetUser)
	app.Get("/user/:id", controllers.GetUserByID)
	app.Post("/user", controllers.CreateUser)
	app.Patch("/user/:id", controllers.PatchUserByID)
	app.Delete("/user/:id", controllers.DeleteUserByID)
	;

}