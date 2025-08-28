package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		TokenLookup:  "cookie:token",
		ContextKey:   "jwt", // simpan token di ctx.Locals("jwt")
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			if token, ok := c.Locals("jwt").(*jwt.Token); ok {
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					if idFloat, ok := claims["user_id"].(float64); ok {
						c.Locals("user_id", int(idFloat))
					}
					if role, ok := claims["role"].(string); ok {
						c.Locals("role", role)
					}
				}
			}
			return c.Next()
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}
