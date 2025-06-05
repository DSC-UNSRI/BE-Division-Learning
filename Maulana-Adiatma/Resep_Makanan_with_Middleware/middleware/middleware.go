package middleware

import (
	"context"
	"net/http"
	"strings"

	"resepku/database"
	"resepku/models"
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

		var user models.Negara
		err := database.DB.QueryRow(
			"SELECT id, email_users, role_users FROM data_negara WHERE token_users = ?",
			token,
		).Scan(&user.ID, &user.Email, &user.Role)

		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextUser, user)
		next(w, r.WithContext(ctx))
	}
}

func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(ContextUser).(models.Negara)
		if !ok || user.Role != "admin" {
			http.Error(w, "Forbidden: admin only", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
