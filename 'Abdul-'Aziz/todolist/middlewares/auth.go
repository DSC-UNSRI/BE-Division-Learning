package middleware

import (
	"context"
	"net/http"
	"strings"

	"tugas/todolist/helper"
	"tugas/todolist/lib"
)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")

			if tokenString == "" {
				helper.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
                return
			}

			tokenString = strings.TrimPrefix(tokenString, "Bearer ")

			claims, err := lib.VerifyJWT(tokenString)

			    if err != nil {
                    helper.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
                    return
                }

                ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
                next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}