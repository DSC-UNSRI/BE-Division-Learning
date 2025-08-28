package main

import (
	"backend/api"
	"backend/config"
	"backend/models"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{}, &models.Event{})

	store := session.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("session", store)
		return c.Next()
	})

	api.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not specified in .env
	}

	log.Fatal(app.Listen(":" + port))
}
