package middleware

import (
	"Task_5/database"
	"Task_5/utils"
	"context"
	"net/http"
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
		err := database.DB.QueryRow("SELECT id, role FROM nasabah WHERE token = ? AND deleted_at IS NULL", token).Scan(&userIDStr, &role)
		if err != nil {
			http.Error(w, "Unauthorized - invalid token", http.StatusUnauthorized)
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Unauthorized - invalid user ID", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, RoleKey, role)
		ctx = context.WithValue(ctx, "token", token)
		next(w, r.WithContext(ctx))
	}
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value(utils.RoleKey)
		if role == nil || role.(string) != "admin" {
			http.Error(w, "Forbidden - admin access only", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

