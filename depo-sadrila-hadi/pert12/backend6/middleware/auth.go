package middleware

import (
	"context"
	"net/http"
	"nobar-backend/auth"
)

type contextKey string

const UserIDKey contextKey = "userID"
const UserRoleKey contextKey = "userRole"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, `{"message":"Unauthorized"}`, http.StatusUnauthorized)
				return
			}
			http.Error(w, `{"message":"Bad request"}`, http.StatusBadRequest)
			return
		}

		tokenStr := c.Value
		claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, `{"message":"Unauthorized"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UserRoleKey, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}