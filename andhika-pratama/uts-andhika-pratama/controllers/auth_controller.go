package controllers

import (
	"uts/database"

	"encoding/json"
	"net/http"

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