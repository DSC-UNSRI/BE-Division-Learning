package middleware

import (
	"context"
	"database/sql"
	"uts/database"
	"uts/models"
	"net/http"
	"strings"
	"time"
)

type UserContextKey string

const userKey = UserContextKey("user")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenValue := parts[1]

		var token models.Token
		err := database.DB.QueryRow("SELECT user_id, expires_at FROM tokens WHERE value = ?", tokenValue).Scan(&token.UserID, &token.ExpiresAt)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Database error on token validation", http.StatusInternalServerError)
			return
		}

		if time.Now().After(token.ExpiresAt) {
			_, err := database.DB.Exec("DELETE FROM tokens WHERE value = ?", tokenValue)
			if err != nil {
				http.Error(w, "Failed to clear expired token", http.StatusInternalServerError)
				return
			}
			http.Error(w, "Token has expired", http.StatusUnauthorized)
			return
		}

		var user models.User
		err = database.DB.QueryRow("SELECT id, name, email, role, created_at FROM users WHERE id = ?", token.UserID).Scan(
			&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "User associated with token not found", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Database error on user retrieval", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), userKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserFromContext(ctx context.Context) (models.User, bool) {
	user, ok := ctx.Value(userKey).(models.User)
	return user, ok
} 