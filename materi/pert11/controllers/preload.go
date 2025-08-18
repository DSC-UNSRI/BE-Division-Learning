package controllers

import (
	"os"
	"pert11/database"
	"pert11/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func GetAllDivisions(c *fiber.Ctx) error {
	var divisions []models.Division
	if err := database.DB.Preload("Member").Find(&divisions).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(divisions)
}

func GetAllEvents(c *fiber.Ctx) error {
	var events []models.Event
	if err := database.DB.Preload("Division").Find(&events).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(events)
}

func GetAllProjects(c *fiber.Ctx) error {
	var projects []models.Project
	if err := database.DB.Preload("Member").Find(&projects).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(projects)
}

func GetAllMembers(c *fiber.Ctx) error {
	var members []models.Member
	query := database.DB.Preload("Division")
	if gender := c.Query("gender"); gender != "" {
		query = query.Where("gender = ?", gender)
	}
	if err := query.Find(&members).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(members)
}

func CreateMember(c *fiber.Ctx) error {
	var member models.Member

	if err := c.BodyParser(&member); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(member.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to hash password"})
	}
	member.Password = string(hashedPassword)

	if err := database.DB.Create(&member).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	member.Password = ""

	return c.Status(201).JSON(member)
}

func Login(c *fiber.Ctx) error {
	var body struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.Member
	if err := database.DB.Where("name = ?", body.Name).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"division_id": user.DivisionID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not login"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    signed,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"message": "Login success",
		"user":    user,
	})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"message": "Logged out",
	})
}
