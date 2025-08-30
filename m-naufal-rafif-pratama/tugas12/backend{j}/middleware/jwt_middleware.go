package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Cookies("token")
		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		if role != "" && claims["role"] != role {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
		}

		c.Locals("user_id", claims["user_id"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}
