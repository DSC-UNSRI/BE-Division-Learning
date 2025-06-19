package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/config"
)

type ContextKey string
const SpeakerIDKey ContextKey = "speakerID"

func AuthSpeakerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]

		var speakerID int
		err := config.DB.QueryRow("SELECT id FROM speakers WHERE auth_key = ?", token).Scan(&speakerID)
		if err != nil {
			http.Error(w, "Invalid authentication token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), SpeakerIDKey, speakerID)
		next(w, r.WithContext(ctx))
	}
}
