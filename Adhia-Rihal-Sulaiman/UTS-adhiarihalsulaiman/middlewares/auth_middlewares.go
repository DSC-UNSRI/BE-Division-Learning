package middlewares

import (
	"uts_adhia/database"
	"uts_adhia/utils"

	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

func AdminBypassMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole, ok := r.Context().Value(utils.UserRoleKey).(string)
		if ok && userRole == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

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

		var userID string
		var userRole string
		var userType string
		var expiresAt sql.NullTime

		err := database.DB.QueryRow(
			`SELECT u.user_id, u.role, u.type, t.expires_at
			FROM users u
            JOIN tokens t ON u.user_id = t.user_id
            WHERE t.token_value = ? AND u.deleted_at IS NULL`,
			token,
		).Scan(&userID, &userRole, &userType, &expiresAt)

		if err == sql.ErrNoRows {
			http.Error(w, "Unauthorized - invalid token", http.StatusUnauthorized)
			return
		}

		if err != nil {
			fmt.Printf("DEBUG: AuthMiddleware Database Query Error: %v\n", err)
			http.Error(w, "Unauthorized - token validation error", http.StatusUnauthorized)
			return
		}

		if expiresAt.Valid && time.Now().After(expiresAt.Time) {
			_, err := database.DB.Exec("DELETE FROM tokens WHERE token_value = ?", token)
			if err != nil {
				fmt.Printf("DEBUG: Error deleting expired token %s: %v\n", token, err)
			}
			http.Error(w, "Unauthorized - token has expired", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), utils.UserIDKey, userID)
		ctx = context.WithValue(ctx, utils.UserRoleKey, userRole)
		ctx = context.WithValue(ctx, utils.UserTypeKey, userType)
		ctx = context.WithValue(ctx, utils.TokenValueKey, token)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
