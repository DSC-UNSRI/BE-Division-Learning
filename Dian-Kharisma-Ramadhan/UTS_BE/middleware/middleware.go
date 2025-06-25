package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"UTS_BE/database"
	"strings"
	"strconv"
)

type contextKey string

const (
	UserIDKey contextKey = "userID"
	RoleKey   contextKey = "role"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1]

		var userIDStr string
		var role string

		err := database.DB.QueryRow("SELECT id, role FROM users WHERE token = ? AND deleted_at IS NULL", token).Scan(&userIDStr, &role)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Unauthorized - invalid token", http.StatusUnauthorized)
			} else {
				http.Error(w, "Database error", http.StatusInternalServerError)
			}
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Unauthorized - invalid user ID format", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, RoleKey, role)

		next(w, r.WithContext(ctx))
	}
}

func WithAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		AuthMiddleware(handler).ServeHTTP(w, r)
	}
}

func PremiumOnlyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleCtx := r.Context().Value(RoleKey)
		if roleCtx == nil || roleCtx.(string) != "premium" {
			http.Error(w, "Access denied: Premium only", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}