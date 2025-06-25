package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s - %s", r.Method, r.URL.Path, start.Format(time.RFC1123))
		next(w, r)
	}
}