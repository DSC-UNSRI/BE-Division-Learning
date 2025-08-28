package routes

import (
	"backend8/config"
	"backend8/controllers"
	"backend8/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, cfg config.Config) {
	auth := controllers.AuthController{Cfg: cfg}
	event := controllers.EventController{}
	user := controllers.UserController{}

	r := app.Group("/api")
	r.Post("/register", auth.Register)
	r.Post("/login", auth.Login)
	r.Post("/logout", auth.Logout)

	r.Get("/event", event.List)
	r.Post("/event", middlewares.AuthRequired(cfg), middlewares.AdminOnly(), event.Create)
	r.Patch("/event/:id", middlewares.AuthRequired(cfg), middlewares.AdminOnly(), event.Update)
	r.Delete("/event/:id", middlewares.AuthRequired(cfg), middlewares.AdminOnly(), event.Delete)

	r.Get("/me", middlewares.AuthRequired(cfg), auth.Me)
	r.Patch("/profile/:id", middlewares.AuthRequired(cfg), user.UpdateProfile)
}
