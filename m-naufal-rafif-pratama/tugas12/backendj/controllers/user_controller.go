package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"tugas12/database"
	"tugas12/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetMe(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func UpdateProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, ok := c.Locals("user_id").(float64)
	if !ok || fmt.Sprint(uint(userID)) != id {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: You can only update your own profile"})
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	if name := c.FormValue("name"); name != "" {
		user.Name = name
	}
	if password := c.FormValue("password"); password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash new password"})
		}
		user.Password = string(hashedPassword)
	}

	file, err := c.FormFile("profile_picture")
	if err == nil && file != nil {
		if err := os.MkdirAll("./uploads", 0755); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create uploads directory"})
		}

		filename := fmt.Sprintf("profile_%s%s", id, filepath.Ext(file.Filename))
		savePath := filepath.Join("./uploads", filename)

		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
		}

		user.ProfilePicture = "/uploads/" + filename
	}

	database.DB.Save(&user)
	return c.JSON(fiber.Map{"message": "Profile updated", "user": user})
}