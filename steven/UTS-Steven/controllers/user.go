package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"uts-gdg/database"
	"uts-gdg/models"
	"uts-gdg/utils"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	role := r.FormValue("role")

	var user models.User
	err := database.DB.QueryRow("SELECT email FROM users WHERE email = ? AND deleted_at IS NULL", email).
		Scan(&user.Email)
	if err == nil {
		http.Error(w, "Email used", http.StatusUnauthorized)
		return
	}

	if(email == "" || name == "" || password == "" || role == ""){
		http.Error(w, "email, name, password can not be empty", http.StatusInternalServerError)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO users (email, name, password, role) VALUES (?, ?, ?, ?)", email, name, hash, role)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	var user models.User
	err := database.DB.QueryRow("SELECT id, email, password, token FROM users WHERE email = ? AND deleted_at IS NULL", email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Token)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := utils.GenerateToken(32)
	expire := time.Now().Add(24 * time.Hour)

	_, err = database.DB.Exec("UPDATE users SET token = ?, token_expire = ? WHERE email = ? AND deleted_at IS NULL", token, expire,  email)
	if err != nil {
		http.Error(w, "Failed to set token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func ForgotPassword (w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := strings.TrimSpace(r.FormValue("email"))

	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ? AND deleted_at IS NULL)", email).Scan(&exists)
	if err != nil || !exists {
		http.Error(w, "Email not found", http.StatusNotFound)
		return
	}

	token := utils.GenerateToken(32)

	expire := time.Now().Add(15 * time.Second)
	_, err = database.DB.Exec("UPDATE users SET reset_token = ?, reset_token_expire = ? WHERE email = ? AND deleted_at IS NULL", token, expire, email)
	if err != nil {
		http.Error(w, "Failed to save reset token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Reset token generated. Use this token to reset password.",
		"token":   token,
	})
}

func ResetPassword (w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.FormValue("token")
	newPassword := r.FormValue("password")

	if token == "" || newPassword == "" {
		http.Error(w, "Token and new password are required", http.StatusBadRequest)
		return
	}

	var userID int
	var expireTime time.Time

	err := database.DB.QueryRow(
		"SELECT id, reset_token_expire FROM users WHERE reset_token = ? AND deleted_at IS NULL", token,
	).Scan(&userID, &expireTime)

	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	if time.Now().After(expireTime) {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec(
		"UPDATE users SET password = ?, reset_token = NULL, reset_token_expire = NULL WHERE id = ?",
		hash, userID,
	)
	if err != nil {
		http.Error(w, "Failed to reset password", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password has been reset successfully",
	})
}