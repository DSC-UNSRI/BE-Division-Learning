package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(startTime)
		log.Printf("LOGGER: %s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, duration)
	})
}