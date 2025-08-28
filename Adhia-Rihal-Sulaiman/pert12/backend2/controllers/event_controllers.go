package controllers

import (
	"backend2/database"
	"backend2/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetEvent(c *fiber.Ctx) error {
	var events []models.Event
	if err := database.DB.Find(&events).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to fetch events"})
	}
	return c.Status(200).JSON(fiber.Map{"events": events})
}

func PostEvent(c *fiber.Ctx) error {
	location := c.FormValue("location")
	start := c.FormValue("start")

	if location == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Location is required"})
	}
	if start == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Start time is required"})
	}

	startTime, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid start time format"})
	}

	var coverPath string
	file, err := c.FormFile("cover")
	if err == nil {
		coverPath = fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), file.Filename)
		if err := c.SaveFile(file, coverPath); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "Failed to save cover"})
		}
	}

	event := models.Event{
		Location: location,
		Start:    startTime,
		Cover:    coverPath,
	}

	if err := database.DB.Create(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to create event"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Event created successfully"})
}

func UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Event not found"})
	}

	location := c.FormValue("location")
	start := c.FormValue("start")

	if location != "" {
		event.Location = location
	}

	if start != "" {
		startTime, err := time.Parse(time.RFC3339, start)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "Invalid start time format"})
		}
		event.Start = startTime
	}

	file, err := c.FormFile("cover")
	if err == nil {
		coverPath := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), file.Filename)
		if err := c.SaveFile(file, coverPath); err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "Failed to save new cover"})
		}
		event.Cover = coverPath
	}

	if err := database.DB.Save(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to update event"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Event updated successfully"})
}

func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Event not found"})
	}

	if err := database.DB.Delete(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to delete event"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Event deleted successfully"})
}
