package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/artichys/uts-raki/repository" 
	"github.com/artichys/uts-raki/utils"     
)

func AuthMiddleware(sessionRepo *repository.SessionRepository, next http.Handler) http.Handler {
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

		session, err := sessionRepo.GetSessionByToken(tokenString)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid or expired token.")
				return
			}
			utils.ErrorResponse(w, http.StatusInternalServerError, "Database error checking token: "+err.Error())
			return
		}
		if session == nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid token.")
			return
		}

		if time.Now().After(session.ExpiresAt) {
			_ = sessionRepo.DeleteSessionByToken(session.Token)
			utils.ErrorResponse(w, http.StatusUnauthorized, "Token has expired.")
			return
		}

		ctx := context.WithValue(r.Context(), "userID", session.UserID)
		ctx = context.WithValue(ctx, "userType", session.UserType)

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