package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"Tugas-Pertemuan-7/database"
)

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	providedAuthKey := r.FormValue("auth")

	if username == "" || providedAuthKey == "" {
		http.Error(w, "Username and auth key are required", http.StatusBadRequest)
		return
	}

	var storedAuthKey string
	err = database.DB.QueryRow("SELECT auth_key FROM users WHERE username = ? AND deleted_at IS NULL", username).Scan(&storedAuthKey)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal server error during login", http.StatusInternalServerError)
		}
		return
	}

	if providedAuthKey == storedAuthKey {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Login successful",
		})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}