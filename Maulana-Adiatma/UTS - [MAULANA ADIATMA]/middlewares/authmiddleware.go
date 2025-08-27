package middleware

import (
	"context"
	"net/http"
	"strings"

	"utsquora/database"
	"utsquora/models"
)

type contextKey string

const ContextUser contextKey = "user"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: missing Bearer token", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		var user models.User
		err := database.DB.QueryRow(
			"SELECT id, username, email, role FROM users WHERE token = ?",
			token,
		).Scan(&user.ID, &user.Username, &user.Email, &user.Role)

		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextUser, user)
		next(w, r.WithContext(ctx))
	}
}

func PremiumOnly(next http.HandlerFunc) http.HandlerFunc { 
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(ContextUser).(models.User)
		if !ok || user.Role != "premium" {
			http.Error(w, "Access restricted to premium users", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}