package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/config"
	"github.com/joho/godotenv"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		godotenv.Load()
		authKey := os.Getenv("AUTH_KEY")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]
		
		if token != authKey {
			var exists bool
			err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM speakers WHERE auth_key = ?)", token).Scan(&exists)
			
			if err != nil || !exists {
				http.Error(w, "Invalid authentication token", http.StatusUnauthorized)
				return
			}
		}

		next(w, r)
	}
}

func RequireSpeakerAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]
		
		var speakerID int
		err := config.DB.QueryRow("SELECT id FROM speakers WHERE auth_key = ?", token).Scan(&speakerID)
		
		if err != nil {
			http.Error(w, "Invalid speaker authentication", http.StatusUnauthorized)
			return
		}

		r.Header.Set("X-Speaker-ID", string(speakerID))
		next(w, r)
	}
}
