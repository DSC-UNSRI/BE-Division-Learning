package controllers

import (
	"backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetEvents(c *fiber.Ctx) error {
	var events []models.Event
	config.DB.Find(&events)

	return c.JSON(fiber.Map{"events": events})
}

func CreateEvent(c *fiber.Ctx) error {
	event := new(models.Event)
	if err := c.BodyParser(event); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	config.DB.Create(&event)

	return c.JSON(fiber.Map{"message": "Event created successfully", "event": event})
}

func UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if config.DB.First(&event, id).Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Event not found"})
	}

	if err := c.BodyParser(&event); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	config.DB.Save(&event)

	return c.JSON(fiber.Map{"message": "Event updated successfully", "event": event})
}

func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if config.DB.First(&event, id).Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Event not found"})
	}

	config.DB.Delete(&event)

	return c.JSON(fiber.Map{"message": "Event deleted successfully"})
}
