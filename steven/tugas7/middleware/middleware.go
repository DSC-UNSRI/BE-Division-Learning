package middleware

import (
	"context"
	"net/http"
	"tugas5/database"
	"tugas5/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Unauthorized - no token", http.StatusUnauthorized)
			return
		}

		const prefix = "Bearer "
		if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
			http.Error(w, "Unauthorized - invalid token format", http.StatusUnauthorized)
			return
		}

		token := authHeader[len(prefix):]

		var exists bool
		var userID string
		var role string
		err := database.DB.QueryRow(
			"SELECT EXISTS(SELECT 1 FROM stores WHERE token = ? AND deleted_at IS NULL), id, role FROM stores WHERE token = ? AND deleted_at IS NULL",
			token, token,
		).Scan(&exists, &userID, &role)
		if err != nil || !exists {
			http.Error(w, "Unauthorized - invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), utils.StoreIDKey, userID)
		ctx = context.WithValue(ctx, utils.RoleKey, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
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