package main

import (
	"oauth-2_redis/config"
	"oauth-2_redis/database"
	"oauth-2_redis/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.ENVLoad()
	database.Init()
	// database.Migrate()
	database.Redis()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api := app.Group("/api")

	app.Static("/assets", "./assets")
	routes.Routes(api)

	app.Listen(":3000")
}
