package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"Tugas-Pertemuan-7/database"
	"Tugas-Pertemuan-7/models"
	"Tugas-Pertemuan-7/utils"

	"golang.org/x/crypto/bcrypt"
)

func GetDirectors(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, role FROM directors WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	directors := []models.Director{}
	for rows.Next() {
		director := models.Director{}
		if err := rows.Scan(&director.ID, &director.Name, &director.Role); err != nil {
			http.Error(w, "Failed to scan director data", http.StatusInternalServerError)
			return
		}
		directors = append(directors, director)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"directors": directors,
	})
}

func GetDirectorByID(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	director := models.Director{}
	err := database.DB.QueryRow("SELECT id, name, role FROM directors WHERE id = ? AND deleted_at IS NULL", id).Scan(&director.ID, &director.Name, &director.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Director not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(director)
}

func CreateDirector(w http.ResponseWriter, r *http.Request) {
	director := models.Director{}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	director.Name = r.FormValue("name")
	password := r.FormValue("password") 
	director.Role = r.FormValue("role")

	if director.Name == "" || password == "" { 
		http.Error(w, "Name and password are required fields", http.StatusBadRequest)
		return
	}

	if director.Role == "" {
		director.Role = "user"
	} else if director.Role != "admin" && director.Role != "user" {
		http.Error(w, "Invalid role. Role must be 'admin' or 'user'", http.StatusBadRequest)
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM directors WHERE name = ? AND deleted_at IS NULL)", director.Name).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error checking director existence", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Director with this name already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	director.PasswordHash = string(hashedPassword)

	res, err := database.DB.Exec("INSERT INTO directors (name, role, password_hash) VALUES (?, ?, ?)", director.Name, director.Role, director.PasswordHash) // Masukkan hash password
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
		return
	}
	director.ID = int(resID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Director successfully created",
		"director": director,
	})
}

func UpdateDirector(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	directorIDFromContext, ok := r.Context().Value(utils.DirectorIDKey).(int)
	if !ok {
		http.Error(w, "Internal server error: Director ID not found in context", http.StatusInternalServerError)
		return
	}

	parsedID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	directorRoleFromContext, ok := r.Context().Value(utils.RoleKey).(string)
	if !ok {
		http.Error(w, "Internal server error: Director role not found in context", http.StatusInternalServerError)
		return
	}

	if directorRoleFromContext != "admin" && directorIDFromContext != parsedID {
		http.Error(w, "Forbidden: You can only update your own director profile unless you are an admin", http.StatusForbidden)
		return
	}

	director := models.Director{}
	err = database.DB.QueryRow("SELECT id, name, role FROM directors WHERE id = ? AND deleted_at IS NULL", id).Scan(&director.ID, &director.Name, &director.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Director not found", http.StatusNotFound)
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

	name := r.FormValue("name")
	role := r.FormValue("role")
	password := r.FormValue("password")

	updateFields := make(map[string]interface{})
	if name != "" && name != director.Name {
		updateFields["name"] = name
		director.Name = name
	}
	if role != "" && role != director.Role {
		if role != "admin" && role != "user" {
			http.Error(w, "Invalid role. Role must be 'admin' or 'user'", http.StatusBadRequest)
			return
		}
		if directorRoleFromContext != "admin" {
			http.Error(w, "Forbidden: Only admins can change roles", http.StatusForbidden)
			return
		}
		updateFields["role"] = role
		director.Role = role
	}
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		updateFields["password_hash"] = hashedPassword
	}

	if len(updateFields) > 0 {
		query := "UPDATE directors SET "
		params := []interface{}{}
		first := true
		for key, val := range updateFields {
			if !first {
				query += ", "
			}
			query += key + " = ?"
			params = append(params, val)
			first = false
		}
		query += " WHERE id = ? AND deleted_at IS NULL"
		params = append(params, id)

		_, err = database.DB.Exec(query, params...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Director updated successfully",
		"director": director,
	})
}

func DeleteDirector(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	directorIDFromContext, ok := r.Context().Value(utils.DirectorIDKey).(int)
	if !ok {
		http.Error(w, "Internal server error: Director ID not found in context", http.StatusInternalServerError)
		return
	}

	parsedID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	directorRoleFromContext, ok := r.Context().Value(utils.RoleKey).(string)
	if !ok {
		http.Error(w, "Internal server error: Director role not found in context", http.StatusInternalServerError)
		return
	}

	if directorRoleFromContext != "admin" && directorIDFromContext != parsedID {
		http.Error(w, "Forbidden: You can only delete your own director profile unless you are an admin", http.StatusForbidden)
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM directors WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error checking director existence", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Director not found or already deleted", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE directors SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Director deleted successfully",
		"id":      id,
	})
}

func GetDirectorByFilmID(w http.ResponseWriter, r *http.Request, filmID string) {
	if filmID == "" {
		http.Error(w, "Film ID parameter is missing", http.StatusBadRequest)
		return
	}

	var directorID int
	err := database.DB.QueryRow("SELECT director_id FROM films WHERE id = ? AND deleted_at IS NULL", filmID).Scan(&directorID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Film not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	director := models.Director{}
	err = database.DB.QueryRow("SELECT id, name, role FROM directors WHERE id = ? AND deleted_at IS NULL", directorID).Scan(&director.ID, &director.Name, &director.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Director associated with this film not found or has been deleted", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(director)
}