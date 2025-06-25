package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"UTS_BE/database"
	"strings"
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

		var userID int
		var role string

		err := database.DB.QueryRow("SELECT id, role FROM users WHERE token = ? AND deleted_at IS NULL", token).
			Scan(&userID, &role)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Unauthorized - invalid token", http.StatusUnauthorized)
			} else {
				http.Error(w, "Database error", http.StatusInternalServerError)
			}
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
