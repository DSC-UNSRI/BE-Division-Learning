package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Incoming Request: Method=%s, Path=%s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Request Handled: Method=%s, Path=%s, Duration=%s", r.Method, r.URL.Path, time.Since(start))
	})
}