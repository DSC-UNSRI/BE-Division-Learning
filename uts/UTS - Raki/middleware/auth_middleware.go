package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/artichys/uts-raki/utils" 
)


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid or expired token: "+err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		ctx = context.WithValue(ctx, "username", claims.Username)
		ctx = context.WithValue(ctx, "userType", claims.UserType)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthorizeMiddleware(requiredUserType string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userType := r.Context().Value("userType")
			if userType == nil {
				utils.ErrorResponse(w, http.StatusInternalServerError, "User type not found in context. AuthMiddleware missing?")
				return
			}

			if userType.(string) != requiredUserType {
				utils.ErrorResponse(w, http.StatusForbidden, "Access denied. Requires "+requiredUserType+" account.")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}