package controllers

import (
	"UTS_BE/database"
	"UTS_BE/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	username := r.FormValue("username")
	secretCode := r.FormValue("secret_code")
	newPassword := r.FormValue("new_password")

	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	var storedSecret string
	var storedHint string

	err := database.DB.QueryRow("SELECT secret_code, secret_hint FROM users WHERE username = ? AND deleted_at IS NULL", username).
		Scan(&storedSecret, &storedHint)

	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if secretCode == "" {
		json.NewEncoder(w).Encode(map[string]string{
			"message":     "Please enter your secret code",
			"secret_hint": storedHint,
		})
		return
	}

	if newPassword == "" {
		http.Error(w, "New password is required", http.StatusBadRequest)
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

	newToken := utils.GenerateToken(32)
	_, err = database.DB.Exec("UPDATE users SET token = ? WHERE username = ?", newToken, username)
	if err != nil {
		http.Error(w, "Failed to update token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password reset successful",
		"token":   newToken,
	})
}

