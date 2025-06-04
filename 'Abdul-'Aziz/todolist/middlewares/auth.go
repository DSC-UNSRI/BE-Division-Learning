package middleware

import (
	"context"
	"net/http"
	"strings"

	"tugas/todolist/lib"
)

func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")

			if tokenString == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
			}

			tokenString = strings.TrimPrefix(tokenString, "Bearer ")

			claims, err := lib.VerifyJWT(tokenString)

			    if err != nil {
                    http.Error(w, "Unauthorized", http.StatusUnauthorized)
                    return
                }

                ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
                next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}