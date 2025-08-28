package main

import (
	"pert12/config"
	"pert12/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ENVinit()
	database.DBInit()
	app := fiber.New()

	app.Listen(":8080")
}