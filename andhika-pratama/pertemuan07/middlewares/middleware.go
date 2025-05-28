package middleware

import (
	"pertemuan05/database"
	"pertemuan05/utils"

	"context"
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

		var exists bool
		var lecturerID string
		var lecturerRole string

		err := database.DB.QueryRow(
			"SELECT EXISTS(SELECT 1 FROM lecturers WHERE token = ? AND deleted_at IS NULL), lecturer_id, role FROM lecturers WHERE token = ? AND deleted_at IS NULL",
			token, token,
		).Scan(&exists, &lecturerID, &lecturerRole)
		if err != nil || !exists {
			http.Error(w, "Unauthorized - invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), utils.LecturerIDKey, lecturerID)
		ctx = context.WithValue(ctx, utils.RoleKey, lecturerRole)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value(utils.RoleKey)
		if role == nil || role.(string) != "old" {
			http.Error(w, "Forbidden - old lecturers access only", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}