package main

import (
	"backend/api"
	"backend/config"
	"backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	app := fiber.New()

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{}, &models.Event{})

	store := session.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("session", store)
		return c.Next()
	})

	api.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}