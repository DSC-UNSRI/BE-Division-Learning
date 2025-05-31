package controllers

import (
	"be_pert7/database"
	"be_pert7/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// GetChefs retrieves all non-deleted chefs.
func GetChefs(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT chef_id, chef_name, speciality, experience, username, role, deleted_at FROM chefs WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, "Failed to retrieve chefs: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	chefs := []models.Chef{}
	for rows.Next() {
		chef := models.Chef{}
		var deletedAt sql.NullTime
		if err := rows.Scan(&chef.ChefID, &chef.Name, &chef.Speciality, &chef.Experience, &chef.Username, &chef.Role, &deletedAt); err != nil {
			http.Error(w, "Failed to scan chef data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if deletedAt.Valid {
			chef.DeletedAt = &deletedAt.Time
		}
		chefs = append(chefs, chef)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"chefs": chefs})
}

// GetChefByID retrieves a single non-deleted chef by their ID.
func GetChefByID(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "Chef ID is required", http.StatusBadRequest)
		return
	}

	var chef models.Chef
	var deletedAt sql.NullTime
	err := database.DB.QueryRow("SELECT chef_id, chef_name, speciality, experience, username, role, deleted_at FROM chefs WHERE chef_id = ? AND deleted_at IS NULL", id).Scan(
		&chef.ChefID, &chef.Name, &chef.Speciality, &chef.Experience, &chef.Username, &chef.Role, &deletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Chef not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if deletedAt.Valid {
		chef.DeletedAt = &deletedAt.Time
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"chef": chef,
	})
}

// CreateChef creates a new chef entry.
func CreateChef(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	chef := models.Chef{
		Name:       r.FormValue("chef_name"),
		Speciality: r.FormValue("speciality"),
		Username:   r.FormValue("username"),
	}

	password := r.FormValue("password")
	if password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	chef.Password = string(hashedPassword)

	experienceStr := r.FormValue("experience")
	if experienceStr == "" {
		http.Error(w, "Experience is required", http.StatusBadRequest)
		return
	}
	if chef.Experience, err = strconv.Atoi(experienceStr); err != nil {
		http.Error(w, "Invalid experience format", http.StatusBadRequest)
		return
	}

	role := r.FormValue("role")
	if role == "" {
		chef.Role = "rookie" // Default role
	} else if role != "head" && role != "rookie" {
		http.Error(w, "Invalid role. Role must be 'head' or 'rookie'", http.StatusBadRequest)
		return
	} else {
		chef.Role = role
	}

	if chef.Name == "" || chef.Username == "" {
		http.Error(w, "Chef name and username are required", http.StatusBadRequest)
		return
	}

	var usernameExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM chefs WHERE username = ?)", chef.Username).Scan(&usernameExists)
	if err != nil {
		http.Error(w, "Database error while checking for existing username: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if usernameExists {
		http.Error(w, "Chef with this username already exists", http.StatusConflict)
		return
	}

	res, err := database.DB.Exec("INSERT INTO chefs (chef_name, speciality, experience, username, password, token, role) VALUES (?, ?, ?, ?, ?, ?, ?)",
		chef.Name, chef.Speciality, chef.Experience, chef.Username, chef.Password, chef.Token, chef.Role)
	if err != nil {
		http.Error(w, "Failed to create chef: "+err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	chef.ChefID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Chef successfully created",
		"chef":    chef,
	})
}

// UpdateChef updates an existing chef.
func UpdateChef(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "Chef ID is required", http.StatusBadRequest)
		return
	}

	chef := models.Chef{}
	var deletedAt sql.NullTime
	err := database.DB.QueryRow("SELECT chef_id, chef_name, speciality, experience, username, password, token, role, deleted_at FROM chefs WHERE chef_id = ? AND deleted_at IS NULL", id).
		Scan(&chef.ChefID, &chef.Name, &chef.Speciality, &chef.Experience, &chef.Username, &chef.Password, &chef.Token, &chef.Role, &deletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Chef not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if deletedAt.Valid {
		chef.DeletedAt = &deletedAt.Time
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	updateFields := []string{}
	updateValues := []interface{}{}

	if name := r.FormValue("chef_name"); name != "" {
		chef.Name = name
		updateFields = append(updateFields, "chef_name = ?")
		updateValues = append(updateValues, name)
	}
	if speciality := r.FormValue("speciality"); speciality != "" {
		chef.Speciality = speciality
		updateFields = append(updateFields, "speciality = ?")
		updateValues = append(updateValues, speciality)
	}
	if experienceStr := r.FormValue("experience"); experienceStr != "" {
		experience, err := strconv.Atoi(experienceStr)
		if err != nil {
			http.Error(w, "Invalid experience format", http.StatusBadRequest)
			return
		}
		chef.Experience = experience
		updateFields = append(updateFields, "experience = ?")
		updateValues = append(updateValues, experience)
	}
	if username := r.FormValue("username"); username != "" {
		chef.Username = username
		updateFields = append(updateFields, "username = ?")
		updateValues = append(updateValues, username)
	}
	if password := r.FormValue("password"); password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		chef.Password = string(hashedPassword)
		updateFields = append(updateFields, "password = ?")
		updateValues = append(updateValues, hashedPassword)
	}
	if role := r.FormValue("role"); role != "" {
		if role != "head" && role != "rookie" {
			http.Error(w, "Invalid role. Role must be 'head' or 'rookie'", http.StatusBadRequest)
			return
		}
		chef.Role = role
		updateFields = append(updateFields, "role = ?")
		updateValues = append(updateValues, role)
	}

	if len(updateFields) == 0 {
		http.Error(w, "No fields provided for update", http.StatusBadRequest)
		return
	}

	query := "UPDATE chefs SET " + strings.Join(updateFields, ", ") + " WHERE chef_id = ? AND deleted_at IS NULL"
	updateValues = append(updateValues, id)

	if _, err := database.DB.Exec(query, updateValues...); err != nil {
		http.Error(w, "Failed to update chef: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Chef updated successfully",
		"chef":    chef,
	})
}

// DeleteChef soft deletes a chef.
func DeleteChef(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "Chef ID is required", http.StatusBadRequest)
		return
	}

	// Check if chef is assigned to any active menus
	var hasMenus bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM menus WHERE chef_id = ? AND deleted_at IS NULL)", id).Scan(&hasMenus)
	if err != nil {
		http.Error(w, "Database error checking chef's menus: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if hasMenus {
		http.Error(w, "Cannot delete chef: assigned to active menus", http.StatusBadRequest)
		return
	}

	// Check if the chef exists and is not already deleted
	var chefExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM chefs WHERE chef_id = ? AND deleted_at IS NULL)", id).Scan(&chefExists)
	if err != nil {
		http.Error(w, "Database error checking chef existence: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !chefExists {
		http.Error(w, "Chef not found or already deleted", http.StatusNotFound)
		return
	}

	// Perform soft delete
	if _, err := database.DB.Exec("UPDATE chefs SET token = '', deleted_at = ? WHERE chef_id = ?", time.Now(), id); err != nil {
		http.Error(w, "Failed to delete chef: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Chef deleted successfully",
		"chef_id": id,
	})
}

// GetChefsBySpeciality retrieves all non-deleted chefs based on their speciality.
func GetChefsBySpeciality(w http.ResponseWriter, r *http.Request, speciality string) {
	if speciality == "" {
		http.Error(w, "Speciality is required", http.StatusBadRequest)
		return
	}

	rows, err := database.DB.Query(`
		SELECT chef_id, chef_name, speciality, experience, username, role, deleted_at
		FROM chefs
		WHERE speciality = ? AND deleted_at IS NULL`, speciality)

	if err != nil {
		http.Error(w, "Database error while retrieving chefs by speciality: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	chefs := []models.Chef{}
	for rows.Next() {
		chef := models.Chef{}
		var deletedAt sql.NullTime
		if err := rows.Scan(&chef.ChefID, &chef.Name, &chef.Speciality, &chef.Experience, &chef.Username, &chef.Role, &deletedAt); err != nil {
			http.Error(w, "Failed to scan chef data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if deletedAt.Valid {
			chef.DeletedAt = &deletedAt.Time
		}
		chefs = append(chefs, chef)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating through chef rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(chefs) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":    "No chefs found for this speciality",
			"speciality": speciality,
			"chefs":      []models.Chef{},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"speciality": speciality,
		"chefs":      chefs,
	})
}