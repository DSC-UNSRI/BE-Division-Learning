package main

import (
	"backend/config"
	"backend/database"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ENVLoad()
	database.Init()
	database.Migrate()
	app := fiber.New()

	routes.Login(app)
	routes.Routes(app)

	app.Listen(":3306")
}