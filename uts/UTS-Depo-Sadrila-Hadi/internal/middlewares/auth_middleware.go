package middlewares

import (
	"context"
	"forum-app/internal/models"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userContextKey = contextKey("userClaims")

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func RoleMiddleware(requiredRole string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(userContextKey).(*models.Claims)
			if !ok {
				http.Error(w, "User claims not found in context", http.StatusInternalServerError)
				return
			}

			if claims.UserType != requiredRole {
				http.Error(w, "Forbidden: You don't have access to this resource", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		}
	}
}

func GetClaimsFromContext(r *http.Request) (*models.Claims, bool) {
	claims, ok := r.Context().Value(userContextKey).(*models.Claims)
	return claims, ok
}