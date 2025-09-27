package main

import (
	"log"
	"tugas12/config"
	"tugas12/database"
	"tugas12/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()
	database.Migrate()

	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PATCH,PUT,DELETE,OPTIONS",
	}))

	app.Static("/uploads", "./uploads")

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}