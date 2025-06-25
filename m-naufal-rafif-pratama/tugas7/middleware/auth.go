package middleware

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"tugas7/utils"
)

type contextKey string

const (
	StudentIDKey contextKey = "studentID"
	EmailKey     contextKey = "email"
	OrgIDKey     contextKey = "orgID"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Auth Middleware: Starting authentication")
		
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Auth Middleware: No Authorization header found")
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		log.Printf("Auth Middleware: Authorization header: %s", authHeader)

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			log.Println("Auth Middleware: Invalid Authorization format")
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := bearerToken[1]
		log.Printf("Auth Middleware: Attempting to validate token: %s", tokenString)

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			log.Printf("Auth Middleware: Invalid token: %v", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		log.Printf("Auth Middleware: Token validated successfully")
		log.Printf("Auth Middleware: Claims - StudentID: %d, Email: %s, OrgID: %d", 
			claims.StudentID, claims.Email, claims.OrgID)

		r.Header.Set("X-Student-ID", strconv.Itoa(claims.StudentID))
		r.Header.Set("X-Student-Email", claims.Email)
		r.Header.Set("X-Student-OrgID", strconv.Itoa(claims.OrgID))

		next.ServeHTTP(w, r)
	}
} 