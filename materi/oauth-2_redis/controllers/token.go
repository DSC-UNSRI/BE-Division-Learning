package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"oauth-2_redis/database"
	"oauth-2_redis/models"
	"oauth-2_redis/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RefreshToken(c *fiber.Ctx) error {
	rdbCtx := context.Background()

	encRefresh := c.Cookies("refresh_token")
	if encRefresh == "" {
		return c.Status(401).JSON(fiber.Map{"message": "refresh token missing"})
	}

	refreshStr, err := utils.Decrypt(encRefresh)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "invalid refresh token1"})
	}

	token, err := jwt.Parse(refreshStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"message": "invalid refresh token2"})
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := int(claims["id"].(float64))

	userKey := fmt.Sprintf("refresh:%d", userID)
	hash, err := database.Rdb.HGetAll(rdbCtx, userKey).Result()
	if err != nil || len(hash) == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "refresh token not found"})
	}

	hashRefreshtoken := hash["refresh_token"]
	refreshDecrypted, err := utils.Decrypt(hashRefreshtoken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "invalid refresh token3"})
	}
	hashParentToken := hash["parent_token"]
	parentDecrypted := ""
	if hashParentToken != "" {
		parentDecrypted, err = utils.Decrypt(hashParentToken)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"message": "invalid refresh token4"})
		}
	}
	expInt, _ := strconv.ParseInt(hash["exp"], 10, 64)

	if refreshStr == parentDecrypted && refreshStr != refreshDecrypted {
		database.Rdb.Del(rdbCtx, userKey)
		return c.Status(401).JSON(fiber.Map{"message": "refresh token mismatch, please login again"})
	}

	if time.Now().Unix() > expInt {
		database.Rdb.Del(rdbCtx, userKey)
		return c.Status(401).JSON(fiber.Map{"message": "refresh token expired"})
	}

	_, err = utils.GenerateAccessToken(c, userID, claims["name"].(string), claims["role"].(string))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "failed generate access token"})
	}

	signedRefresh, err := utils.GenerateRefreshToken(c, userID, claims["name"].(string), claims["role"].(string))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal buat refresh token"})
	}

	refreshToken := models.Token{
		RefreshToken: signedRefresh,
		ParentToken:  refreshStr,
		UserID:       userID,
		Exp:          time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	if err := utils.SaveRefreshToken(c, refreshToken); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Redis error"})
	}

	return c.JSON(fiber.Map{
		"message":      "refresh success",
	})
}
