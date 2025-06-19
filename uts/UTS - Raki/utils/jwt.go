package utils

import (
    "log"
    "os"
    "time"

    "github.com/dgrijalva/jwt-go"
)

type Claims struct {
    UserID   int    `json:"user_id"`
    Username string `json:"username"`
    UserType string `json:"user_type"`
    jwt.StandardClaims
}

func GenerateToken(userID int, username, userType string) (string, error) {
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Println("JWT_SECRET environment variable is not set!")
        return "", nil
    }

    expirationTime := time.Now().Add(24 * time.Hour) 
    claims := &Claims{
        UserID:   userID,
        Username: username,
        UserType: userType,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        log.Println("JWT_SECRET environment variable is not set!")
        return nil, nil 
    }

    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(jwtSecret), nil
    })

    if err != nil {
        return nil, err
    }
    if !token.Valid {
        return nil, jwt.ErrSignatureInvalid
    }
    return claims, nil
}