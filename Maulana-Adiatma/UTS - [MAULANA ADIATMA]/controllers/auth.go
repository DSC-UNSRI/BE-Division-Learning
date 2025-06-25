package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"utsquora/database"
	"utsquora/middlewares"
	"utsquora/models"
	"utsquora/utils"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	input.Email = strings.ToLower(input.Email)

	var exists bool
	err = database.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)",
		input.Email,
	).Scan(&exists)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	_, err = database.DB.Exec(`
    INSERT INTO users (username, email, password, role)
    VALUES (?, ?, ?, 'free')`,
		input.Username, input.Email, string(hashedPassword))

	if err != nil {
		log.Println("Insert error:", err)
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	input.Email = strings.ToLower(input.Email)

	var user models.User
	var dbToken sql.NullString

	err = database.DB.QueryRow(`
		SELECT id, username, email, password, token FROM users WHERE email = ?`,
		input.Email,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &dbToken)

	if err != nil {
		http.Error(w, "Email not found", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}

	TokenLogin := utils.GenerateToken(32)
	_, err = database.DB.Exec("UPDATE users SET token = ? WHERE id = ?", TokenLogin, user.ID)
	if err != nil {
		http.Error(w, "Failed to update token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Login successful",
		"username": user.Username,
		"email":    user.Email,
		"token":    TokenLogin,
	})
}

func RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token := utils.GenerateToken(32)
	expiry := time.Now().Add(15 * time.Minute)

	_, err := database.DB.Exec(`
		UPDATE users SET reset_token = ?, reset_token_expiry = ? WHERE email = ?`,
		token, expiry, strings.ToLower(input.Email))

	if err != nil {
		http.Error(w, "Failed to generate reset token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Reset token generated",
		"token":   token,
	})
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	log.Println("Input token:", input.Token)

	var userID int
	var expiry time.Time
	err := database.DB.QueryRow(`
	SELECT id, reset_token_expiry FROM users WHERE reset_token = ?`,
		input.Token).Scan(&userID, &expiry)

	if err != nil {
		log.Println("DB lookup error:", err)
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}

	log.Println("Token expiry:", expiry)

	if time.Now().After(expiry) {
		http.Error(w, "Token expired", http.StatusUnauthorized)
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)

	_, err = database.DB.Exec(`
		UPDATE users SET password = ?, reset_token = NULL, reset_token_expiry = NULL
		WHERE id = ?`, string(hashed), userID)

	if err != nil {
		http.Error(w, "Failed to reset password", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password successfully reset",
	})
}

func UpgradeToPremium(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUser).(models.User)

	if user.Role == "premium" {
		http.Error(w, "Already premium", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("UPDATE users SET role = 'premium' WHERE id = ?", user.ID)
	if err != nil {
		http.Error(w, "Failed to upgrade", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Account upgraded to premium",
	})
}
