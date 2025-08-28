package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	jwt "github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() fiber.Handler {
    return jwtware.New(jwtware.Config{
        SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
        TokenLookup: "cookie:token",

        SuccessHandler: func(c *fiber.Ctx) error {
            user := c.Locals("user").(*jwt.Token)
            claims := user.Claims.(jwt.MapClaims)

            if idFloat, ok := claims["user_id"].(float64); ok {
                c.Locals("user_id", int(idFloat))
            }
			  if role, ok := claims["role"].(string); ok {
				c.Locals("role", role)
			}

            return c.Next()
        },

        ErrorHandler: func(c *fiber.Ctx, err error) error {
            return c.Status(401).JSON(fiber.Map{
                "message": "Unauthorized",
            })
        },
    })
}