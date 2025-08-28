package controllers

import (
	"fmt"
	"path/filepath"
	"time"

	"backend8/database"
	"backend8/models"

	"github.com/gofiber/fiber/v2"
)

type EventController struct{}

func (EventController) List(c *fiber.Ctx) error {
	var events []models.Event
	database.DB.Table("events").Find(&events)
	for i := range events {
		if events[i].Cover == "" {
			events[i].Cover = "http://127.0.0.1:3000/assets/defaults/event-cover.png"
		}
	}
	return c.JSON(fiber.Map{"events": events})
}

func (EventController) Create(c *fiber.Ctx) error {
	type payload struct{ Location string `json:"location"` }
	var p payload
	if err := c.BodyParser(&p); err != nil || p.Location == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid payload"})
	}
	event := models.Event{
		Location: p.Location,
		Start:    time.Now().Format(time.RFC3339),
		Cover:    "http://127.0.0.1:3000/assets/defaults/event-cover.png",
	}
	database.DB.Table("events").Create(&event)
	return c.JSON(fiber.Map{"message": "created"})
}

func (EventController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event
	if err := database.DB.Table("events").First(&event, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not found"})
	}
	location := c.FormValue("location")
	start := c.FormValue("start")
	if location != "" {
		event.Location = location
	}
	if start != "" {
		event.Start = start
	}
	file, err := c.FormFile("cover")
	if err == nil && file != nil {
		name := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
		path := filepath.Join("uploads", name)
		if err := c.SaveFile(file, path); err == nil {
			event.Cover = "http://127.0.0.1:3000/uploads/" + name
		}
	}
	database.DB.Table("events").Save(&event)
	return c.JSON(fiber.Map{"message": "updated"})
}

func (EventController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Table("events").Delete(&models.Event{}, id)
	return c.JSON(fiber.Map{"message": "deleted"})
}
