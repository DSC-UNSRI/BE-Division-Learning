package controllers

import (
	"uts/database"
	"uts/models"
	"uts/utils"

	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)
func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	question := r.FormValue("question")
	answer := r.FormValue("answer")

	if username == "" || password == "" {
		http.Error(w, "Missing required fields: username, password", http.StatusBadRequest)
		return
	} 
	
	if question == "" || answer == "" {
		http.Error(w, "Please input your security question and answer for password recovery", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ? AND deleted_at IS NULL)", username).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error checking username existence", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "Username with this username already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	result, err := database.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)",
		username, hashedPassword,
	)

	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	userID, err := result.LastInsertId()

	if err != nil {
		http.Error(w, "Failed to retrieve user ID after registration", http.StatusInternalServerError)
		return
	}

	hashedAnswer, err := bcrypt.GenerateFromPassword([]byte(answer), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash security answer", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO challenges (user_id, question, answer) VALUES (?, ?, ?)",
		userID, question, hashedAnswer,
	)

	if err != nil {
		http.Error(w, "Failed to register user's security question and answer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" ||  password == "" {
		http.Error(w, "Missing required fields: username, password", http.StatusBadRequest)
		return
	}

	var user models.User
	err := database.DB.QueryRow("SELECT user_id, username, password, role, type FROM users WHERE username = ? AND deleted_at IS NULL", username).
		Scan(&user.UserID, &user.Username, &user.Password, &user.Role, &user.Type)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	newToken := utils.GenerateToken(32)
	expirationDate := time.Now().Add(time.Hour)

	_, err = database.DB.Exec("INSERT INTO tokens (token_value, expires_at) VALUES (?, ?)",
		newToken, expirationDate,
	)

	if err != nil {
		http.Error(w, "Failed to create a token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": newToken,
	})
}