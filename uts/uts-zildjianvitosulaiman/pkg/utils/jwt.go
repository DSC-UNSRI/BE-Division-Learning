package utils

import (
	"errors"
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

func ValidateJWT(tokenString string) (*JWTClaims, error) {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	return claims, nil
}
