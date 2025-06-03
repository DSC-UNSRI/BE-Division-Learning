package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"Tugas-Pertemuan-7/database"
	"Tugas-Pertemuan-7/models"
	"Tugas-Pertemuan-7/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterDirector(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	password := r.FormValue("password")
	role := r.FormValue("role")

	if name == "" || password == "" {
		http.Error(w, "Name and password are required fields", http.StatusBadRequest)
		return
	}

	if role == "" {
		role = "user"
	} else if role != "admin" && role != "user" {
		http.Error(w, "Invalid role. Role must be 'admin' or 'user'", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM directors WHERE name = ? AND deleted_at IS NULL)", name).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error checking director existence", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Director with this name already exists", http.StatusConflict)
		return
	}

	res, err := database.DB.Exec("INSERT INTO directors (name, password_hash, role) VALUES (?, ?, ?)", name, hashedPassword, role)
	if err != nil {
		http.Error(w, "Failed to register director", http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to get last insert ID", http.StatusInternalServerError)
		return
	}

	director := models.Director{
		ID:   int(id),
		Name: name,
		Role: role,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Director registered successfully",
		"director": director,
	})
}

func LoginDirector(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	password := r.FormValue("password")

	if name == "" || password == "" {
		http.Error(w, "Name and password are required", http.StatusBadRequest)
		return
	}

	var director models.Director
	err = database.DB.QueryRow("SELECT id, name, password_hash, role FROM directors WHERE name = ? AND deleted_at IS NULL", name).
		Scan(&director.ID, &director.Name, &director.PasswordHash, &director.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal server error during login", http.StatusInternalServerError)
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(director.PasswordHash), []byte(password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := utils.GenerateToken(32)

	_, err = database.DB.Exec("UPDATE directors SET token = ? WHERE id = ? AND deleted_at IS NULL", token, director.ID)
	if err != nil {
		http.Error(w, "Failed to set token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"token":   token,
		"role":    director.Role,
	})
}