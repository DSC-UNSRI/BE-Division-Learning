package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"backend/config"
	"backend/models"
	"backend/routes"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Event{})
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	app.Static("/assets", "./assets")

	routes.SetupRoutes(app)

	app.Listen(":3000")
}