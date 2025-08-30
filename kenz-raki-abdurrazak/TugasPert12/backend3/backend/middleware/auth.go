package middleware

import (
	"context"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"backend/handlers"
)

var jwtKey = []byte("supersecretkey")

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: Token not provided", http.StatusUnauthorized)
			return
		}
		
		tokenString := strings.Split(authHeader, " ")[1]

		claims := &handlers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}