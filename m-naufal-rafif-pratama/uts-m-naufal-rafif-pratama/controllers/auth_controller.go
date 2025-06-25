package controllers

import (
	"database/sql"
	"encoding/json"
	"uts/database"
	"uts/models"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Role             string `json:"role"`
	SecurityQuestion string `json:"security_question"`
	SecurityAnswer   string `json:"security_answer"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" || req.SecurityQuestion == "" || req.SecurityAnswer == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	if req.Role != "premium" {
		req.Role = "free"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	hashedAnswer, err := bcrypt.GenerateFromPassword([]byte(req.SecurityAnswer), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash security answer", http.StatusInternalServerError)
		return
	}

	newUser := models.User{
		Name:             req.Name,
		Email:            req.Email,
		Password:         string(hashedPassword),
		Role:             req.Role,
		SecurityQuestion: req.SecurityQuestion,
		SecurityAnswer:   string(hashedAnswer),
		CreatedAt:        time.Now(),
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", req.Email).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Email already in use", http.StatusConflict)
		return
	}

	result, err := database.DB.Exec("INSERT INTO users (name, email, password, role, security_question, security_answer, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		newUser.Name, newUser.Email, newUser.Password, newUser.Role, newUser.SecurityQuestion, newUser.SecurityAnswer, newUser.CreatedAt)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	userID, _ := result.LastInsertId()
	newUser.ID = int(userID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	var foundUser models.User
	err := database.DB.QueryRow("SELECT id, password, role FROM users WHERE email = ?", req.Email).Scan(&foundUser.ID, &foundUser.Password, &foundUser.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenString := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)

	_, err = database.DB.Exec("INSERT INTO tokens (value, user_id, expires_at) VALUES (?, ?, ?)", tokenString, foundUser.ID, expiresAt)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var req ForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var securityQuestion string
	err := database.DB.QueryRow("SELECT security_question FROM users WHERE email = ?", req.Email).Scan(&securityQuestion)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"security_question": securityQuestion})
}

type ResetPasswordRequest struct {
	Email          string `json:"email"`
	SecurityAnswer string `json:"security_answer"`
	NewPassword    string `json:"new_password"`
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var storedAnswer string
	var userID int
	err := database.DB.QueryRow("SELECT id, security_answer FROM users WHERE email = ?", req.Email).Scan(&userID, &storedAnswer)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedAnswer), []byte(req.SecurityAnswer)); err != nil {
		http.Error(w, "Invalid security answer", http.StatusUnauthorized)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash new password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("UPDATE users SET password = ? WHERE id = ?", string(hashedPassword), userID)
	if err != nil {
		http.Error(w, "Failed to reset password", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password reset successfully"})
} 