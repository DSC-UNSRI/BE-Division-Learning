package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
	"uts_adhia/database"
	"uts_adhia/models"

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
	userRole := r.FormValue("role")
	userType := r.FormValue("type")
	question := r.FormValue("question")
	answer := r.FormValue("answer")

	if username == "" || password == "" || userRole == "" || userType == "" {
		http.Error(w, "Missing required fields: username, password, role, type", http.StatusBadRequest)
		return
	}

	if question == "" || answer == "" {
		http.Error(w, "Missing required fields: security question and answer", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	if userRole != "user" && userRole != "admin" {
		http.Error(w, "Invalid role. Role must be 'user' or 'admin'", http.StatusBadRequest)
		return
	}

	if userType != "free" && userType != "premium" {
		http.Error(w, "Invalid type. Type must be 'free' or 'premium'", http.StatusBadRequest)
		return
	}

	var userExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ? AND deleted_at IS NULL)", username).Scan(&userExists)

	if err != nil {
		http.Error(w, "Database error while checking for existing user", http.StatusInternalServerError)
		return
	}

	if userExists {
		http.Error(w, "User with this username already exists", http.StatusConflict)
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	res, err := tx.Exec("INSERT INTO users(username, password, role, type) VALUES (?, ?, ?, ?)",
		username, hashedPassword, userRole, userType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUserID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "User created, but failed to retrieve user ID", http.StatusInternalServerError)
		return
	}

	hashedAnswer, err := bcrypt.GenerateFromPassword([]byte(answer), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash security answer", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("INSERT INTO forgets (user_id, question, answer_hash) VALUES (?, ?, ?)",
		newUserID, question, hashedAnswer,
	)

	if err != nil {
		http.Error(w, "Failed to store security question and answer for user", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{
		ID:       int(newUserID),
		Username: username,
		Role:     userRole,
		Type:     userType,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User successfully created",
		"user":    user,
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT user_id, username, role, type, created_at, updated_at, deleted_at FROM users WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Role, &user.Type, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
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
		http.Error(w, "User ID missing in URL", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
		return
	}

	user := models.User{}
	err = database.DB.QueryRow("SELECT user_id, username, role, type, created_at, updated_at, deleted_at FROM users WHERE user_id = ? AND deleted_at IS NULL", idInt).
		Scan(&user.ID, &user.Username, &user.Role, &user.Type, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

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
		http.Error(w, "User ID missing in URL", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
		return
	}

	currentUser := models.User{}
	err = database.DB.QueryRow("SELECT user_id, username, password, role, type FROM users WHERE user_id = ? AND deleted_at IS NULL", idInt).
		Scan(&currentUser.ID, &currentUser.Username, &currentUser.Password, &currentUser.Role, &currentUser.Type)

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
	userRole := r.FormValue("role")
	userType := r.FormValue("type")

	updateFields := []string{}
	updateValues := []interface{}{}

	if username != "" && username != currentUser.Username {
		var userExists bool
		err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ? AND user_id != ? AND deleted_at IS NULL)", username, idInt).Scan(&userExists)

		if err != nil {
			http.Error(w, "Database error while checking for existing username", http.StatusInternalServerError)
			return
		}

		if userExists {
			http.Error(w, "User with this username already exists", http.StatusConflict)
			return
		}
		updateFields = append(updateFields, "username = ?")
		updateValues = append(updateValues, username)
	}

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		updateFields = append(updateFields, "password = ?")
		updateValues = append(updateValues, hashedPassword)
	}

	if userRole != "" {
		if userRole != "user" && userRole != "admin" {
			http.Error(w, "Invalid role. Role must be 'user' or 'admin'", http.StatusBadRequest)
			return
		}
		updateFields = append(updateFields, "role = ?")
		updateValues = append(updateValues, userRole)
	}

	if userType != "" {
		if userType != "free" && userType != "premium" {
			http.Error(w, "Invalid type. Type must be 'free' or 'premium'", http.StatusBadRequest)
			return
		}
		updateFields = append(updateFields, "type = ?")
		updateValues = append(updateValues, userType)
	}

	if len(updateFields) == 0 {
		http.Error(w, "No fields provided for update", http.StatusBadRequest)
		return
	}

	query := "UPDATE users SET " + strings.Join(updateFields, ", ") + ", updated_at = ? WHERE user_id = ? AND deleted_at IS NULL"
	updateValues = append(updateValues, time.Now())
	updateValues = append(updateValues, idInt)

	_, err = database.DB.Exec(query, updateValues...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedUser := models.User{}
	err = database.DB.QueryRow("SELECT user_id, username, role, type, created_at, updated_at, deleted_at FROM users WHERE user_id = ? AND deleted_at IS NULL", idInt).
		Scan(&updatedUser.ID, &updatedUser.Username, &updatedUser.Role, &updatedUser.Type, &updatedUser.CreatedAt, &updatedUser.UpdatedAt, &updatedUser.DeletedAt)

	if err != nil {
		http.Error(w, "Failed to fetch updated user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request, userID string) {
	if userID == "" {
		http.Error(w, "User ID missing in URL", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
		return
	}

	var userExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id = ? AND deleted_at IS NULL)", idInt).Scan(&userExists)

	if err != nil {
		http.Error(w, "Database error while checking user existence", http.StatusInternalServerError)
		return
	}

	if !userExists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM tokens WHERE user_id = ?", idInt)
	if err != nil {
		http.Error(w, "Failed to remove user's tokens", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE user_id = ?", idInt)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}