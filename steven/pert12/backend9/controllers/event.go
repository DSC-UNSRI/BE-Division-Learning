package controllers

import (
	"backend/database"
	"backend/models"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllEvents(c *fiber.Ctx) error {
	var events []models.Event
	if err := database.DB.Find(&events).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(events)
}

func CreateEvent(c *fiber.Ctx) error {
	userRole := c.Locals("role")

	if userRole != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "only admin can create event",
		})
	}

	var event models.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if event.Cover == "" {
		event.Cover = "https://i.pravatar.cc/150"
	}

	if err := database.DB.Create(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Event berhasil dibuat",
	})
}

func UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
    var event models.Event

    if err := database.DB.First(&event, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Event not found"})
    }

    location := c.FormValue("location")
	start := c.FormValue("start")
    file, err := c.FormFile("cover")

    if location != "" {
        event.Location = location
    }

    if start != "" {
        parsedTime, err := time.Parse("2006-01-02T15:04:05", start)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid start format"})
        }
        event.Start = parsedTime
    }

    if err == nil && file != nil {
		if file.Size > 1*1024*1024 {
			return c.Status(400).JSON(fiber.Map{"error": "Ukuran foto maksimal 1MB"})
		}

		ext := filepath.Ext(file.Filename)
		allowedExt := map[string]bool{
			".png":  true,
			".jpg":  true,
			".jpeg": true,
		}
		if !allowedExt[ext] {
			return c.Status(400).JSON(fiber.Map{"error": "Foto berformat harus PNG, JPG, atau JPEG"})
		}

		if event.Cover != "" && !strings.Contains(event.Cover, "pravatar.cc") {
			oldPath := "." + strings.TrimPrefix(event.Cover, "http://127.0.0.1:3000")
			_ = os.Remove(oldPath)
		}
		
        filename := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
        savePath := "./assets/" + filename
        if err := c.SaveFile(file, savePath); err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Failed to save photo"})
        }
        event.Cover = "http://127.0.0.1:3000/assets/" + filename
    }

    if err := database.DB.Save(&event).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
	
	return c.Status(201).JSON(fiber.Map{
		"message": "Event berhasil diupdate",
	})
}

func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
    var event models.Event

    if err := database.DB.First(&event, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Event not found"})
    }

    if event.Cover != "" && !strings.Contains(event.Cover, "pravatar.cc") {
        oldPath := "." + strings.TrimPrefix(event.Cover, "http://127.0.0.1:3000")
        _ = os.Remove(oldPath)
    }

    if err := database.DB.Delete(&event).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{
        "message": "Event berhasil dihapus",
    })
}