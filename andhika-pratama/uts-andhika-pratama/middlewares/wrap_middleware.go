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

func WithOwnsQuestionAuth(handler func(http.ResponseWriter, *http.Request, string), id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wrappedControllerHandler := http.HandlerFunc(func(innerW http.ResponseWriter, innerR *http.Request) {
			handler(innerW, innerR, id)
		})

		middlewareChain := OwnsQuestionMiddleware(
			wrappedControllerHandler,
			id,
		)

		utils.ApplyMiddlewares(middlewareChain, AuthMiddleware).ServeHTTP(w, r)
	}
}

func WithOwnsAnswerAuth(handler func(http.ResponseWriter, *http.Request, string), id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wrappedControllerHandler := http.HandlerFunc(func(innerW http.ResponseWriter, innerR *http.Request) {
			handler(innerW, innerR, id)
		})

		middlewareChain := OwnsAnswerMiddleware(
			wrappedControllerHandler,
			id,
		)

		utils.ApplyMiddlewares(middlewareChain, AuthMiddleware).ServeHTTP(w, r)
	}
}