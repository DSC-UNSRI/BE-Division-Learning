package main

import (
	"pert12/config"
	"pert12/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ENVLoad()
	database.DBLoad()

	app := fiber.New()

	app.Listen(":8080")
}