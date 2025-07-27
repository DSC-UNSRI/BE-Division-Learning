package main

import (
	"pert11/config"
	"pert11/database"
	"pert11/middleware"
	"pert11/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.ENVLoad()
	database.Init()
	database.Migrate()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://127.0.0.1:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api := app.Group("/api", middleware.Protected())

	routes.Login(app)
	routes.Routes(api)

	app.Listen(":3000")
}
