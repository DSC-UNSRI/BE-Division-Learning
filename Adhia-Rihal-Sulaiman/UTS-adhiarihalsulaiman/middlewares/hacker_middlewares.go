package middlewares

import (
	"net/http"
	"uts_adhia/utils"
)

func HackerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userType, okType := r.Context().Value(utils.UserTypeKey).(string)
		userRole, okRole := r.Context().Value(utils.UserRoleKey).(string)

		if !okType || !okRole {
			http.Error(w, "Unauthorized: User type or role not found in context", http.StatusInternalServerError)
			return
		}

		if userType == "premium" || userRole == "admin" {
			next.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Forbidden: Premium access required", http.StatusForbidden)
	})
}