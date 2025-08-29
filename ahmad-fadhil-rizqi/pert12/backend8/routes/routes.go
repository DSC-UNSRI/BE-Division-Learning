package routes

import (
	"backend8/controllers"
	"backend8/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	auth := controllers.AuthController{}
	event := controllers.EventController{}
	user := controllers.UserController{}

	r := app.Group("/api")
	r.Post("/register", auth.Register)
	r.Post("/login", auth.Login)
	r.Post("/logout", auth.Logout)

	r.Get("/event", event.List)
	r.Post("/event", middlewares.AuthRequired(), middlewares.AdminOnly(), event.Create)
	r.Patch("/event/:id", middlewares.AuthRequired(), middlewares.AdminOnly(), event.Update)
	r.Delete("/event/:id", middlewares.AuthRequired(), middlewares.AdminOnly(), event.Delete)

	r.Get("/me", middlewares.AuthRequired(), auth.Me)
	r.Patch("/profile/:id", middlewares.AuthRequired(), user.UpdateProfile)
}
