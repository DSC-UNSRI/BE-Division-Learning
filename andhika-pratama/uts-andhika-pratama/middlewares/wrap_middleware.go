package middlewares

import (
	"uts/utils"

	"net/http"
)

func WithAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ApplyMiddlewares(handler, AuthMiddleware).ServeHTTP(w, r)
	}
}