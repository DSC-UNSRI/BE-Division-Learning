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

func OldMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value(utils.RoleKey)
		if role == nil || role.(string) != "old" {
			http.Error(w, "Forbidden - old lecturers access only", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CourseOwnershipMiddleware(next http.Handler, courseID string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxLecturerID := r.Context().Value(utils.LecturerIDKey)
		lecturerID, ok := ctxLecturerID.(string)
		if !ok || lecturerID == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if courseID == "" {
			http.Error(w, "course_id is null", http.StatusBadRequest)
			return
		}

		var courseOwnerID string
		err := database.DB.QueryRow("SELECT lecturer_id FROM courses WHERE course_id = ? AND deleted_at IS NULL", courseID).Scan(&courseOwnerID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if courseOwnerID != lecturerID {
			http.Error(w, "Forbidden - the course isn't yours", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}