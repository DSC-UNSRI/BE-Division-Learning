package main

import (
	"be_pertemuan10/config"
	"be_pertemuan10/database"
	"be_pertemuan10/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ENVLoad()
	database.Init()
	database.Migrate()
	app := fiber.New()

	routes.UserRoutes(app)

	app.Listen(":3000")
}
