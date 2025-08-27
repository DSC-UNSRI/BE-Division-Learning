package controllers

import (
	"pert12/database"
	"pert12/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	var existingUser models.User

	name := c.FormValue("name")
	if name == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Name is required",
		})
	}

	email := c.FormValue("email")
	if email == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email is required",
		})
	}
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error;  err == nil {
		return c.Status(409).JSON(fiber.Map{
			"message": "Email is already taken",
		})
	}

	password := c.FormValue("password")
	if password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password is required",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	role := c.FormValue("role")
	if role == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Role is required",
		})
	}
	if role != "user" && role != "admin" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Role must be user or admin",
		})
	}

	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"user": newUser,
	})
}
