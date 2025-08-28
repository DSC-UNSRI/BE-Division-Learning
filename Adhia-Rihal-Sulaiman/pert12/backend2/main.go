package main

import (
	"backend2/config"
	"backend2/database"
	"backend2/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.ENVLoad()
	database.DBLoad()
	database.DBMigrate()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://localhost:5173/",
		AllowCredentials: true,
	}))
	app.Static("/assets", "./assets")

	routes.RoutesList(app)
	app.Listen(":8080")
}
