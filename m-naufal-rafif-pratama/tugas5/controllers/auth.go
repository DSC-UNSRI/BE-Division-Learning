package controllers

import (
	"database/sql"
	"net/http"
	"tugas5/database"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Authorization required", http.StatusUnauthorized)
			return
		}

		var dbPassword string
		err := database.DB.QueryRow(
			"SELECT password FROM students WHERE email = ?", 
			email,
		).Scan(&dbPassword)

		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			} else {
				http.Error(w, "Database error", http.StatusInternalServerError)
			}
			return
		}

		if dbPassword != password {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}