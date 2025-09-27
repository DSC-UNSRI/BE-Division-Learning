package utils

import (
	"context"
	"fmt"
	"oauth-2_redis/database"
	"oauth-2_redis/models"
	"time"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(c *fiber.Ctx, userID int, name, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   userID,
		"name": name,
		"role": role,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
	}

	signed, err := GenerateJWT(claims, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", err
	}

	if err := SetTokenCookie(c, signed, "access_token", 15*time.Minute); err != nil {
		return "", err
	}

	return signed, nil
}

func GenerateRefreshToken(c *fiber.Ctx, userID int, name, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   userID,
		"name": name,
		"role": role,
		"exp":  time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	signed, err := GenerateJWT(claims, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", err
	}

	if err := SetTokenCookie(c, signed, "refresh_token", 7*24*time.Hour); err != nil {
		return "", err
	}

	return signed, nil
}

func GenerateJWT(claims jwt.MapClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func Encrypt(plaintext string) (string, error) {
	keyBytes, _ := base64.StdEncoding.DecodeString(os.Getenv("ENCRYPT_KEY"))

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(encrypted string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(os.Getenv("ENCRYPT_KEY"))
	if err != nil {
		return "", err
	}
	if len(keyBytes) != 32 {
		return "", errors.New("ENCRYPT_KEY harus 32 byte")
	}

	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext terlalu pendek")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func SaveRefreshToken(ctx *fiber.Ctx, token models.Token) error {
	rdbCtx := context.Background()

	key := fmt.Sprintf("refresh:%d", token.UserID)

	encryptedToken, err := Encrypt(token.RefreshToken)
	if err != nil {
		return err
	}

	encryptedParent := ""
	if token.ParentToken != "" {
		encryptedParent, err = Encrypt(token.ParentToken)
		if err != nil {
			return err
		}
	}

	fields := map[string]interface{}{
		"refresh_token": encryptedToken,
		"parent_token":  encryptedParent,
		"exp":           token.Exp,
	}

	if err := database.Rdb.HSet(rdbCtx, key, fields).Err(); err != nil {
		return err
	}

	ttl := time.Until(time.Unix(token.Exp, 0))
	if ttl <= 0 {
		ttl = 7 * 24 * time.Hour
	}
	if err := database.Rdb.Expire(rdbCtx, key, ttl).Err(); err != nil {
		return err
	}

	return nil
}

func SetTokenCookie(c *fiber.Ctx, token string, name string, expiration time.Duration) error {
	encrypted, err := Encrypt(token)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    encrypted,
		Expires:  time.Now().Add(expiration),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return nil
}
