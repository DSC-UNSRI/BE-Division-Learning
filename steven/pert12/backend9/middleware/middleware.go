package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		TokenLookup:  "cookie:token",
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			userToken := c.Locals("user").(*jwt.Token)
			claims := userToken.Claims.(jwt.MapClaims)

			c.Locals("id", claims["id"])

			return c.Next()
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(401).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

func IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken := c.Locals("user").(*jwt.Token)
		claims := userToken.Claims.(jwt.MapClaims)
		role := claims["role"]

		c.Locals("role", role)

		if role != "admin" {
			return c.Status(403).JSON(fiber.Map{"error": "Forbidden: Division access denied"})
		}

		return c.Next()
	}
}