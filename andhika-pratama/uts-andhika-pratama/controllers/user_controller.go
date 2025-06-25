package controllers

import (
	"uts/database"
	"uts/models"

	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	user_role := r.FormValue("role")
	user_type := r.FormValue("type")
	question := r.FormValue("question")
	answer := r.FormValue("answer")

	if  username == "" || password == "" || user_role == "" || user_type == "" {
		http.Error(w, "Missing required fields: username, password, role, type" , http.StatusBadRequest)
		return
	}

	if question == "" || answer == "" {
		http.Error(w, "Please input your security question and answer for password recovery", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	if user_role != "user" && user_role != "admin" {
		http.Error(w, "Invalid role. Role must be 'user' or 'admin'", http.StatusBadRequest)
		return
	}

	if user_type != "free" && user_type != "premium" {
		http.Error(w, "Invalid type. Type must be 'free' or 'premium'", http.StatusBadRequest)
		return
	}

	user := models.User {
		Username: 	  username,
		Password:     string(hashedPassword),
		Role:         user_role,
		Type:         user_type,
	}

	var userExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", user.Username).Scan(&userExists)

	if err != nil {
		http.Error(w, "Database error while checking for existing user", http.StatusInternalServerError)
		return
	}

	if userExists {
		http.Error(w, "User with this username already exists", http.StatusConflict)
		return
	}

	res, err := database.DB.Exec("INSERT INTO users(user_id, username, password, role, type) VALUES (?, ?, ?, ?, ?)",
		user.UserID, user.Username, user.Password, user.Role, user.Type)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "User created, but failed to retrieve user ID", http.StatusInternalServerError)
		return
	}

	user.UserID = int(userID)
	hashedAnswer, err := bcrypt.GenerateFromPassword([]byte(answer), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Failed to hash security answer", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO challenges (user_id, question, answer) VALUES (?, ?, ?)",
		userID, question, hashedAnswer,
	)

	if err != nil {
		http.Error(w, "Failed to store security question and answer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "User successfully created",
		"user": user,
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT user_id, username, role, type FROM users WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.UserID, &user.Username, &user.Role, &user.Type)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users": users,
	})
}

func GetUserByID(w http.ResponseWriter, r *http.Request, userID string) {
	if userID == "" {
		http.Error(w, "Please input user_id in the url", http.StatusBadRequest)
		return
	}

	user := models.User{}
	err := database.DB.QueryRow("SELECT user_id, username, role, type FROM users WHERE user_id = ? AND deleted_at IS NULL", userID).
		Scan(&user.UserID, &user.Username, &user.Role, &user.Type)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user": user,
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request, userID string) {
	if userID == "" {
		http.Error(w, "Please input user_id in the url", http.StatusBadRequest)
		return
	}

	user := models.User{}
	err := database.DB.QueryRow("SELECT user_id, username, password, role, type FROM users WHERE user_id = ? AND deleted_at IS NULL", userID).
		Scan(&user.UserID, &user.Username, &user.Password, &user.Role, &user.Type)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	user_role := r.FormValue("role")
	user_type := r.FormValue("type")

	updateFields := []string{}
	updateValues := []interface{}{}

	if username != "" && username != user.Username {
		var userExists bool
		err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ? AND user_id != ?)", username, userID).Scan(&userExists)

		if err != nil {
			http.Error(w, "Database error while checking for existing username", http.StatusInternalServerError)
			return
		}

		if userExists {
			http.Error(w, "User with this username already exists", http.StatusConflict)
			return
		}
		user.Username = username 
		updateFields = append(updateFields, "username = ?")
		updateValues = append(updateValues, username)
	}

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		user.Password = string(hashedPassword)
		updateFields = append(updateFields, "password = ?")
		updateValues = append(updateValues, hashedPassword)
	}

	if user_role != "" {
		if user_role != "user" && user_role != "admin" {
			http.Error(w, "Invalid role. Role must be 'user' or 'admin'", http.StatusBadRequest)
			return
		}
		user.Role = user_role
		updateFields = append(updateFields, "role = ?")
		updateValues = append(updateValues, user_role)
	}

	if user_type != "" {
		if user_type != "free" && user_type != "premium" {
			http.Error(w, "Invalid type. Type must be 'free' or 'premium'", http.StatusBadRequest)
			return
		}
		user.Type = user_type
		updateFields = append(updateFields, "type = ?")
		updateValues = append(updateValues, user_type)
	}

	if len(updateFields) == 0 {
		http.Error(w, "No fields provided for update", http.StatusBadRequest)
		return
	}

	query := "UPDATE users SET " + strings.Join(updateFields, ", ") + " WHERE user_id = ? AND deleted_at IS NULL"
	updateValues = append(updateValues, userID)

	_, err = database.DB.Exec(query, updateValues...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "User updated successfully",
		"user": user,
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request, userID string) {
	if userID == "" {
		http.Error(w, "Please input user_id in the url", http.StatusBadRequest)
		return
	}

	var userExists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id = ? AND deleted_at IS NULL)", userID).Scan(&userExists)

	if err != nil {
		http.Error(w, "Database error while checking user existence", http.StatusInternalServerError)
		return
	}

	if !userExists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("DELETE FROM tokens WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, "Failed to remove user's tokens", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("UPDATE users SET deleted_at = NOW() WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}