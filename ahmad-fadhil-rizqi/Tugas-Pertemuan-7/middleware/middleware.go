package middleware

import (
	"Tugas-Pertemuan-7/database"
	"Tugas-Pertemuan-7/utils"
	"context"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Unauthorized: No token provided.", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: Invalid token format.", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		var directorID int
		var role string
		err := database.DB.QueryRow(
			"SELECT id, role FROM directors WHERE token = ? AND deleted_at IS NULL",
			token,
		).Scan(&directorID, &role)

		if err != nil {
			http.Error(w, "Unauthorized: Invalid or expired token.", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), utils.DirectorIDKey, directorID)
		ctx = context.WithValue(ctx, utils.RoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(utils.RoleKey).(string)
		if !ok || role != "admin" {
			http.Error(w, "Forbidden: Admin access required.", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr) //
		next.ServeHTTP(w, r)
	})
}