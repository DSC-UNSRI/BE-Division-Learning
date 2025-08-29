package controllers

import (
	"os"
	"time"
	"tugas12/database"
	"tugas12/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	role := c.FormValue("role")

	if role == "" {
		role = "user"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user := &models.User{
		Name:           name,
		Email:          email,
		Password:       string(hashedPassword),
		Role:           role,
		ProfilePicture: "/uploads/default-profile.png",
	}

	database.DB.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Signup success", "user": user})
}

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to sign token"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
		Expires:  time.Now().Add(72 * time.Hour),
	})

	return c.JSON(fiber.Map{
		"message": "Login success",
		"user":    fiber.Map{"id": user.ID, "role": user.Role},
	})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{"message": "Logged out"})
}