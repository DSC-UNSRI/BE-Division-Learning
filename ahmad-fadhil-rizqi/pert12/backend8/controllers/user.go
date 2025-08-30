package controllers

import (
	"fmt"
	"path/filepath"
	"time"

	"backend8/database"
	"backend8/models"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (UserController) UpdateProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := database.DB.Table("users").First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not found"})
	}
	name := c.FormValue("name")
	if name != "" {
		user.Name = name
	}
	file, err := c.FormFile("profile_picture")
	if err == nil && file != nil {
		name := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
		path := filepath.Join("uploads", name)
		if err := c.SaveFile(file, path); err == nil {
			user.ProfilePicture = "http://127.0.0.1:3000/uploads/" + name
		}
	}
	if err := database.DB.Table("users").Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to update"})
	}
	return c.JSON(fiber.Map{"message": "updated"})
}

