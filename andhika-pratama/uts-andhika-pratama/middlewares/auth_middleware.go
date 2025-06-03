package middlewares

import (
	"uts/database"
	"uts/utils"

	"context"
	"database/sql"
	"net/http"
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

		var userID string
		var userRole string
		var userType string

		err := database.DB.QueryRow(
			`SELECT u.user_id, u.role, u.type 
			FROM users u
            JOIN tokens t ON u.user_id = t.user_id
            WHERE t.token_value = ? AND u.deleted_at IS NULL`, 
			token,
		).Scan(&userID, &userRole, &userType)

		if err == sql.ErrNoRows {
			http.Error(w, "Unauthorized - invalid token", http.StatusUnauthorized)
			return
		}

		if err != nil {
			http.Error(w, "Unauthorized - token validation error", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), utils.UserIDKey, userID)
		ctx = context.WithValue(ctx, utils.RoleKey, userRole)
		ctx = context.WithValue(ctx, utils.TypeKey, userType) 
		ctx = context.WithValue(ctx, utils.TokenValueKey, token) 

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}