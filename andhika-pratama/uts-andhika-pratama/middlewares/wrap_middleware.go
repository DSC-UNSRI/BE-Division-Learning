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

func WithPremiumAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ApplyMiddlewares(handler, AuthMiddleware, TypeMiddleware).ServeHTTP(w, r)
	}
}

func WithAdminAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ApplyMiddlewares(handler, AuthMiddleware, RoleMiddleware).ServeHTTP(w, r)
	}
}