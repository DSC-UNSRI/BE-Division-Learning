package lib

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = os.Getenv("ACCESS_TOKEN_SECRET_KEY")
var secretKey = []byte(secret)

func GenerateJWT(userID int, userEmail string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email": userEmail,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
	
		return "", err
	}

	return signedToken, nil
}


func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
       token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
           if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
               return nil, fmt.Errorf("invalid signing method")
           }
           return secretKey, nil
       })

       if err != nil {
           return nil, err
       }

       if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
           return claims, nil
       }

       return nil, fmt.Errorf("invalid token")
   }