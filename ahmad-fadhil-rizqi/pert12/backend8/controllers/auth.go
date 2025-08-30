package controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"backend8/database"
	"backend8/models"
	"backend8/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct{}

func (a AuthController) Register(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	role := c.FormValue("role")
	if role == "" {
		role = "user"
	}
	if name == "" || email == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid payload"})
	}
	var existing models.User
	if err := database.DB.Table("users").Where("email = ?", email).First(&existing).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "email already exists"})
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "internal error"})
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	avatarIdx := rng.Intn(70) + 1
	avatarURL := fmt.Sprintf("https://i.pravatar.cc/150?img=%d", avatarIdx)
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{
		Name:           name,
		Email:          email,
		Password:       string(hash),
		Role:           role,
		ProfilePicture: avatarURL,
	}
	if err := database.DB.Table("users").Create(&user).Error; err != nil {
		errMsg := err.Error()
		if strings.Contains(strings.ToLower(errMsg), "duplicate") || strings.Contains(strings.ToLower(errMsg), "unique") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "email already exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to register"})
	}
	return c.JSON(fiber.Map{
		"message": "registered",
		"user":    user,
	})
}

func (a AuthController) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	var user models.User
	if err := database.DB.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid credentials"})
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid credentials"})
	}
	token, err := utils.GenerateJWT(os.Getenv("JWT_SECRET"), user.ID, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed"})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
	})
	return c.JSON(fiber.Map{
		"message": "logged in",
		"user": fiber.Map{
			"id":   user.ID,
			"role": user.Role,
		},
	})
}

func (a AuthController) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: true,
		Path:     "/",
	})
	return c.JSON(fiber.Map{"message": "logged out"})
}

func (a AuthController) Me(c *fiber.Ctx) error {
	uid := c.Locals("user_id").(uint)
	var user models.User
	if err := database.DB.Table("users").First(&user, uid).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not found"})
	}
	if user.ProfilePicture == "" {
		idx := int(user.ID%70) + 1
		user.ProfilePicture = fmt.Sprintf("https://i.pravatar.cc/150?img=%d", idx)
	}
	return c.JSON(user)
}
