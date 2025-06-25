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
	secret := r.FormValue("secret_code")

	if username == "" || password == ""  || secret == "" {
		http.Error(w, "Missing field", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	token := utils.GenerateToken(32)

	result, err := database.DB.Exec(`
		INSERT INTO users (username, password, token, secret_code)
		VALUES (?, ?, ?, ?, ?)`,
		username, hash, token, secret,
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
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	var u models.User
	err := database.DB.QueryRow(`SELECT id, username, password, FROM users WHERE username = ? AND deleted_at IS NULL`, username).
		Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	newToken := utils.GenerateToken(32)
	_, _ = database.DB.Exec(`UPDATE users SET token = ? WHERE id = ?`, newToken, u.ID)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"token":   newToken,
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

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")

	var userID int
	err := database.DB.QueryRow(`SELECT id FROM users WHERE username = ?`, username).Scan(&userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	resetToken := utils.GenerateToken(32)
	expiry := time.Now().Add(10 * time.Minute)

	_, err = database.DB.Exec(`INSERT INTO forgot_password_tokens (user_id, token, expired_at) VALUES (?, ?, ?)`,
		userID, resetToken, expiry)
	if err != nil {
		http.Error(w, "Failed to create reset token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message":     "Reset token generated",
		"reset_token": resetToken,
	})
}
