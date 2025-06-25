package middlewares

import (
	"uts_adhia/utils"

	"net/http"
)

func RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value(utils.UserRoleKey)
		if role == nil || role.(string) != "admin" {
			http.Error(w, "Forbidden - Admin only", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}