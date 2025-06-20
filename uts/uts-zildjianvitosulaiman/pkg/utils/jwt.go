package utils

import (
	"os"
	"time"
	"uts-zildjianvitosulaiman/domain"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID int             `json:"user_id"`
	Tier   domain.UserTier `json:"tier"`
	jwt.RegisteredClaims
}

func GenerateJWT(user *domain.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &JWTClaims{
		UserID: user.ID,
		Tier:   user.Tier,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
