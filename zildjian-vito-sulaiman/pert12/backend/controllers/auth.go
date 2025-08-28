package controllers

import (
	"backend/config"
	"backend/models"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	// Generate a random profile picture URL
	rand.Seed(time.Now().UnixNano())
	randomID := rand.Intn(1000) // Random ID for picsum.photos
	user.ProfilePicture = fmt.Sprintf("https://picsum.photos/200/200?random=%d", randomID)

	config.DB.Create(&user)

	return c.JSON(fiber.Map{"message": "User created successfully", "user": user})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	var user models.User
	config.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 || user.Password != data["password"] {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	sess, err := c.Locals("session").(*session.Store).Get(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}

	sess.Set("userID", user.ID)
	sess.Save()

	return c.JSON(fiber.Map{"message": "Login successful", "user": fiber.Map{"id": user.ID, "role": user.Role}})
}

func Logout(c *fiber.Ctx) error {
	sess, err := c.Locals("session").(*session.Store).Get(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}

	sess.Destroy()

	return c.JSON(fiber.Map{"message": "Logout successful"})
}