package main

import (
	"pert12/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ENVLoad()
	
	app := fiber.New()

	app.Listen(":8080")
}