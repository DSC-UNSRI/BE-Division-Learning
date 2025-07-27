package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		TokenLookup:  "cookie:token",
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(401).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

func IsDivision1() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken := c.Locals("user").(*jwt.Token)
		claims := userToken.Claims.(jwt.MapClaims)
		divisonID := claims["division_id"]

		if divisonID != 1 {
			return c.Status(403).JSON(fiber.Map{"error": "Forbidden: Division access denied"})
		}

		return c.Next()
	}
}
