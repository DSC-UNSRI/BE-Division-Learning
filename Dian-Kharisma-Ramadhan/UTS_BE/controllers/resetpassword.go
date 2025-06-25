package controllers

import (
	"UTS_BE/database"
	"database/sql"
	"encoding/json"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	secretCode := r.FormValue("secret_code")
	newPassword := r.FormValue("new_password")

	if username == "" || secretCode == "" || newPassword == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	var storedSecret string
	err := database.DB.QueryRow("SELECT secret_code FROM users WHERE username = ? AND deleted_at IS NULL", username).
		Scan(&storedSecret)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if storedSecret != secretCode {
		http.Error(w, "Invalid secret code", http.StatusUnauthorized)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("UPDATE users SET password = ? WHERE username = ?", hashed, username)
	if err != nil {
		http.Error(w, "Failed to update password", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password reset successful",
	})
}
