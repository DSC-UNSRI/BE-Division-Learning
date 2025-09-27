package middleware

import (
	"oauth-2_redis/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		encToken := c.Cookies("access_token")
		if encToken == "" {
			return c.Status(401).JSON(fiber.Map{"message": "unauthorized"})
		}

		tokenStr, err := utils.Decrypt(encToken)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"message": err})
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"message": "invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"message": "invalid token"})
		}

		idFloat, ok := claims["id"].(float64)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"message": "invalid token"})
		}
		userID := int(idFloat)

		role, ok := claims["role"].(string)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"message": "invalid token"})
		}

		c.Locals("id", userID)
		c.Locals("role", role)

		return c.Next()
	}
}
