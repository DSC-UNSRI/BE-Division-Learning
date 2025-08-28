package controllers

import (
	"backend2/database"
	"backend2/models"
	"backend2/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UpdateProfileRequest struct {
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
}

func GetMe(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	var user models.User

	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(200).JSON(user)
}

func UpdateProfile(c *fiber.Ctx) error {
	tokenUserID := c.Locals("user_id").(int)
	paramID := c.Params("id")

	userID, err := strconv.Atoi(paramID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid user ID"})
	}

	if tokenUserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden - you can only update your own profile"})
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	var req UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	updates := make(map[string]interface{})

	if req.Name != "" {
		updates["name"] = req.Name
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
		}
		updates["password"] = string(hashedPassword)
	}

	if _, err := c.FormFile("profile_picture"); err == nil {
		filePath, err := utils.SaveFile(c, "profile_picture", "profile_pictures")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		updates["profile_picture"] = filePath
	}

	if len(updates) > 0 {
		if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update profile"})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Profile updated successfully",
	})
}
