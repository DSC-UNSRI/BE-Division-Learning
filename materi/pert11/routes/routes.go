package routes

import (
	"pert11/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(api fiber.Router) {
	api.Get("/member", controllers.GetAllMembers)
	api.Get("/division", controllers.GetAllDivisions)
	api.Get("/project", controllers.GetAllProjects)
	api.Get("/event", controllers.GetAllEvents)
	api.Post("/member", controllers.CreateMember)
	api.Post("/logout", controllers.Logout)
}
