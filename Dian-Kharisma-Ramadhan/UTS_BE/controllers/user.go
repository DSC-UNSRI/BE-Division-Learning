package controllers

import (
	"UTS_BE/database"
	"UTS_BE/models"	
	"UTS_BE/utils"
	"UTS_BE/middleware"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	secretCode := r.FormValue("secret_code")
	secretHint := r.FormValue("secret_hint")

	if username == "" || password == "" || secretCode == "" || secretHint == "" {
		http.Error(w, "Missing field", http.StatusBadRequest)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	secretCodeHash, err := bcrypt.GenerateFromPassword([]byte(secretCode), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash secret code", http.StatusInternalServerError)
		return
	}

	token := utils.GenerateToken(32)

	result, err := database.DB.Exec(
		"INSERT INTO users (username, password, token, secret_code, secret_hint) VALUES (?, ?, ?, ?, ?)",
		username, passwordHash, token, secretCodeHash,secretHint,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to register: %v", err), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		id = 0
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Register success",
		"id":      id,
		"token":   token,
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	var u models.User
	err := database.DB.QueryRow(
		"SELECT id, username, password FROM users WHERE username = ? AND deleted_at IS NULL", username,
	).Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	newToken := utils.GenerateToken(32)
	expiredAt := time.Now().Add(10 * time.Minute)

	_, err = database.DB.Exec("UPDATE users SET token = ?, token_expired_at = ? WHERE id = ?", newToken, expiredAt, u.ID)
	if err != nil {
		http.Error(w, "Failed to update token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Login successful",
		"token":     newToken,
		"expiresAt": expiredAt.Format(time.RFC3339),
	})
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)

	_, err := database.DB.Exec("UPDATE users SET token = NULL, token_expired_at = NULL WHERE id = ?", userID)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)
	role := r.Context().Value(middleware.RoleKey).(string)

	var username string
	err := database.DB.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"username": username,
		"role":     role,
	})
}