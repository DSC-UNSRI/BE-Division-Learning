package main

import (
	"backend8/config"
	"backend8/database"
	"backend8/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.ENVLoad()
	database.DBLoad()
	database.DBMigrate()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))

	app.Static("/uploads", "uploads")

	routes.Register(app)
	app.Listen(":3000")
}