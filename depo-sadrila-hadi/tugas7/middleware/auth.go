package middleware

import (
	"context"
	"net/http"
	"tugas7/models"
	"tugas7/utils"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := utils.AuthenticateUserFromRequest(r)
		if err != nil {
			authErrMessage := "Authentication failed"
			statusCode := http.StatusUnauthorized

			errMsg := err.Error()
			if errMsg == "authorization header required" ||
				strings.Contains(errMsg, "invalid authorization format") ||
				strings.Contains(errMsg, "invalid base64 encoding") ||
				strings.Contains(errMsg, "invalid credentials format") {
				authErrMessage = errMsg
			} else if strings.Contains(errMsg, "internal server error") {
				statusCode = http.StatusInternalServerError
				authErrMessage = "Internal server error"
			}

			utils.RespondWithError(w, statusCode, authErrMessage)
			return
		}

		ctx := context.WithValue(r.Context(), CurrentUserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetCurrentUserFromContext(ctx context.Context) (models.Mahasiswa, bool) {
	user, ok := ctx.Value(CurrentUserKey).(models.Mahasiswa)
	return user, ok
}