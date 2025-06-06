package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte

func init() {
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		log.Println("Warning: JWT_SECRET_KEY not set, using default key")
		key = "default_secret_key_for_development"
	}
	jwtKey = []byte(key)
}

type Claims struct {
	StudentID int    `json:"student_id"`
	Email     string `json:"email"`
	OrgID     int    `json:"org_id"`
	jwt.RegisteredClaims
}

func GenerateToken(studentID int, email string, orgID int) (string, error) {
	if len(jwtKey) == 0 {
		return "", errors.New("JWT secret key not initialized")
	}

	claims := &Claims{
		StudentID: studentID,
		Email:     email,
		OrgID:     orgID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", err
	}

	log.Printf("Token generated successfully for student ID: %d, OrgID: %d", studentID, orgID)
	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	if len(jwtKey) == 0 {
		return nil, errors.New("JWT secret key not initialized")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return nil, err
	}

	if !token.Valid {
		log.Println("Token is invalid")
		return nil, errors.New("invalid token")
	}

	log.Printf("Token validated successfully. Claims: StudentID=%d, Email=%s, OrgID=%d",
		claims.StudentID, claims.Email, claims.OrgID)
	return claims, nil
} 