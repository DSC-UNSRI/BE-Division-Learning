package middlewares

import (
	"uts_adhia/utils"

	"net/http"
)

func TypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value(utils.UserTypeKey)
		if role == nil || role.(string) != "premium" {
			http.Error(w, "Forbidden - Premium users only", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}