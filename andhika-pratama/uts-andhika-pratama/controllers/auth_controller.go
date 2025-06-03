package controllers

import (
	"database/sql"
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

	_, err = database.DB.Exec("INSERT INTO tokens (user_id, token_value, expires_at) VALUES (?, ?, ?)",
		user.UserID, newToken, expirationDate,
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

func Logout(w http.ResponseWriter, r *http.Request) {
	tokenValue, ok := r.Context().Value(utils.TokenValueKey).(string) 
	if !ok {
		http.Error(w, "Internal server error: Token not found in context", http.StatusInternalServerError)
		return
	}

	_, err := database.DB.Exec("DELETE FROM tokens WHERE token_value = ?", tokenValue)
	if err != nil {
		http.Error(w, "Failed to logout due to internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out successfully"))
}

func InitiatePasswordReset(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() 
	username := r.FormValue("username")

	if username == "" {
		http.Error(w, "Missing required field: username", http.StatusBadRequest)
		return
	}

	var userID int
	err := database.DB.QueryRow("SELECT user_id FROM users WHERE username = ? AND deleted_at IS NULL", username).Scan(&userID)

	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Database error during password reset initiation", http.StatusInternalServerError)
		return
	}

	var challengeQuestion string

	err = database.DB.QueryRow("SELECT question FROM challenges WHERE user_id = ?", userID).Scan(&challengeQuestion)

	if err == sql.ErrNoRows {
		http.Error(w, "No security question found for this user", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Database error retrieving security question", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"question": challengeQuestion,
		"message":  "Please answer your security question",
	})
}

func PasswordReset(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	answer := r.FormValue("answer")  
	newPassword := r.FormValue("new_password")     

	if username == "" || answer == "" {
		http.Error(w, "Missing required fields: username, answer, new_password", http.StatusBadRequest)
		return
	}

	var userID int
	var storedHashedAnswer string 

	err := database.DB.QueryRow(
		"SELECT u.user_id, c.answer FROM users u JOIN challenges c ON u.user_id = c.user_id WHERE u.username = ? AND u.deleted_at IS NULL",
		username,
	).Scan(&userID, &storedHashedAnswer)

	if err == sql.ErrNoRows {
		http.Error(w, "Password reset failed: Invalid username", http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(w, "Internal server error during security answer verification: ", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedHashedAnswer), []byte(answer)); err != nil {
		http.Error(w, "Password reset failed: Invalid security answer", http.StatusUnauthorized)
		return
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Failed to hash new password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("UPDATE users SET password = ? WHERE user_id = ?", hashedNewPassword, userID)
	if err != nil {
		http.Error(w, "Failed to update password", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Password reset successfully"))
}