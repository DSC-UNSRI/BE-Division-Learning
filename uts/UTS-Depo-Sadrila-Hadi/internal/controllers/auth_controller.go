package controllers

import (
	"database/sql"
	"encoding/json"
	"forum-app/internal/database"
	"forum-app/internal/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterPayload struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	SecurityQuestion string `json:"security_question"`
	SecurityAnswer   string `json:"security_answer"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var payload RegisterPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("Request decoded in: %s", time.Since(start))

	if payload.Username == "" || payload.Password == "" || payload.SecurityQuestion == "" || payload.SecurityAnswer == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	hashStart := time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	hashedSecurityAnswer, err := bcrypt.GenerateFromPassword([]byte(payload.SecurityAnswer), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash security answer", http.StatusInternalServerError)
		return
	}
	log.Printf("Hashing completed in: %s", time.Since(hashStart))

	dbStart := time.Now()
	tx, err := database.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	result, err := tx.Exec("INSERT INTO users (username, password) VALUES (?, ?)", payload.Username, string(hashedPassword))
	if err != nil {
		tx.Rollback()
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}
	userID, _ := result.LastInsertId()

	_, err = tx.Exec("INSERT INTO security_answers (user_id, question_text, answer_hash) VALUES (?, ?, ?)",
		userID, payload.SecurityQuestion, string(hashedSecurityAnswer))
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to save security answer", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}
	log.Printf("Database transaction completed in: %s", time.Since(dbStart))

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
	log.Printf("Total registration process took: %s", time.Since(start))
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user models.User
	err := database.DB.QueryRow("SELECT id, username, password, user_type FROM users WHERE username = ?", payload.Username).Scan(&user.ID, &user.Username, &user.Password, &user.UserType)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		UserID:   user.ID,
		Username: user.Username,
		UserType: user.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	log.Printf("Membuat token dengan kunci: '%s'", string(jwtKey))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

type RequestQuestionPayload struct {
	Username string `json:"username"`
}

func RequestSecurityQuestion(w http.ResponseWriter, r *http.Request) {
	var payload RequestQuestionPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var questionText string
	err := database.DB.QueryRow(`
		SELECT sa.question_text 
		FROM security_answers sa
		JOIN users u ON sa.user_id = u.id
		WHERE u.username = ?`, payload.Username).Scan(&questionText)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "If user exists, security question will be retrieved", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve security question", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"security_question": questionText})
}

type ResetPasswordPayload struct {
	Username       string `json:"username"`
	SecurityAnswer string `json:"security_answer"`
	NewPassword    string `json:"new_password"`
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var payload ResetPasswordPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if payload.NewPassword == "" || payload.SecurityAnswer == "" || payload.Username == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	var hashedAnswer string
	var userID int64
	err := database.DB.QueryRow(`
		SELECT sa.answer_hash, u.id
		FROM security_answers sa
		JOIN users u ON sa.user_id = u.id
		WHERE u.username = ?`, payload.Username).Scan(&hashedAnswer, &userID)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedAnswer), []byte(payload.SecurityAnswer))
	if err != nil {
		http.Error(w, "Invalid security answer", http.StatusUnauthorized)
		return
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash new password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("UPDATE users SET password = ? WHERE id = ?", string(newHashedPassword), userID)
	if err != nil {
		http.Error(w, "Failed to update password", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Password has been reset successfully"})
}