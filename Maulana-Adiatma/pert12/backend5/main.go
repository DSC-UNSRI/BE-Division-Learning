package main

import (
	"pert12/config"
	"pert12/database"
	"pert12/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.ENVinit()
	database.DBInit()
	database.DBMigrate()
	app := fiber.New()
	routes.MainRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Static("/assets", "./assets")
	app.Listen(":8080")
}