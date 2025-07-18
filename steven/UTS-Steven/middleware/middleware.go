package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"time"
	"uts-gdg/database"
	"uts-gdg/utils"
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
		var userID int
		var role string
		var tokenExpire sql.NullTime
		
		err := database.DB.QueryRow(
			"SELECT EXISTS(SELECT 1 FROM users WHERE token = ? AND deleted_at IS NULL), id, role, token_expire FROM users WHERE token = ? AND deleted_at IS NULL",
			token, token,
		).Scan(&exists, &userID, &role, &tokenExpire)
		if err != nil || !exists {
			http.Error(w, "Unauthorized - invalid token", http.StatusUnauthorized)
			return
		}

		if tokenExpire.Valid && time.Now().After(tokenExpire.Time) {
			http.Error(w, "Unauthorized - token expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), utils.UserIDKey, userID)
		ctx = context.WithValue(ctx, utils.RoleKey, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PremiumMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value(utils.RoleKey)
		if role == nil || role.(string) != "premium" {
			http.Error(w, "Forbidden - premium access only", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}