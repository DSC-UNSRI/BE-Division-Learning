// controllers/event.go
package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"backend/config"
	"backend/models"
	"backend/utils"
)

// GetEvents endpoint
func GetEvents(c *fiber.Ctx) error {
	var events []models.Event
	// Menggunakan Preload untuk memuat data User yang terkait
	result := config.DB.Preload("User").Find(&events)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get events"})
	}
	return c.JSON(fiber.Map{"events": events})
}

// PostEvent endpoint
func PostEvent(c *fiber.Ctx) error {
	var event models.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	// Mengambil userID dari locals, mengonversinya dari float64 ke uint
	userIDFloat64 := c.Locals("userID").(float64)
	event.UserID = uint(userIDFloat64)

	config.DB.Create(&event)
	return c.JSON(fiber.Map{"message": "Event created successfully"})
}

// UpdateEvent endpoint
// UpdateEvent endpoint
func UpdateEvent(c *fiber.Ctx) error {
	// Dapatkan ID dari parameter URL dan periksa error
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid event ID"})
	}

	var event models.Event
	// Cari event berdasarkan ID
	if result := config.DB.First(&event, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Event not found"})
	}

	// Tangani unggahan file
	file, err := c.FormFile("cover")
	if err == nil {
		filePath, err := utils.SaveFile(file, "./assets/images/covers")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to upload file"})
		}
		event.Cover = filePath
	}

	// Tangani pembaruan lokasi dan start
	if location := c.FormValue("location"); location != "" {
		event.Location = location
	}
	if start := c.FormValue("start"); start != "" {
		event.Start = start
	}

	// Simpan perubahan ke database
	config.DB.Save(&event)
	return c.JSON(fiber.Map{"message": "Event updated successfully"})
}

func DeleteEvent(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	config.DB.Delete(&models.Event{}, id)
	return c.JSON(fiber.Map{"message": "Event deleted successfully"})
}