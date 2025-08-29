package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"tugas12/database"
	"tugas12/models"

	"github.com/gofiber/fiber/v2"
)

func GetEvents(c *fiber.Ctx) error {
	var events []models.Event
	database.DB.Find(&events)
	return c.JSON(fiber.Map{"events": events})
}

func GetEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Event not found"})
	}
	return c.JSON(event)
}

func CreateEvent(c *fiber.Ctx) error {
	type CreateEventRequest struct {
		Location string `json:"location"`
	}

	req := new(CreateEventRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	event := &models.Event{
		Location: req.Location,
		Cover:    "/uploads/cover-default.png",
	}
	database.DB.Create(&event)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Event created", "event": event})
}

func UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Event not found"})
	}

	if location := c.FormValue("location"); location != "" {
		event.Location = location
	}
	if start := c.FormValue("start"); start != "" {
		event.Start = start
	}

	file, err := c.FormFile("cover")
	if err == nil && file != nil {
		if err := os.MkdirAll("./uploads", 0755); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create uploads directory"})
		}
		
		filename := fmt.Sprintf("event_%s%s", id, filepath.Ext(file.Filename))
		savePath := filepath.Join("./uploads", filename)

		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
		}
		
		event.Cover = "/uploads/" + filename
	}

	database.DB.Save(&event)
	return c.JSON(fiber.Map{"message": "Event updated", "event": event})
}

func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Event{}, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Event deleted"})
}