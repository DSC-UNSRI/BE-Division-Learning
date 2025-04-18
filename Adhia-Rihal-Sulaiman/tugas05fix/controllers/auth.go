package controllers

import (
	"be_pert5/database"
	"be_pert5/models"
	"encoding/json"
	"net/http"
	"database/sql"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	var chef models.Chef
	err = database.DB.QueryRow(`
		SELECT id, name, speciality, experience, username, password FROM chefs 
		WHERE username = ? AND deleted_at IS NULL
	`, username).Scan(&chef.ID, &chef.Name, &chef.Speciality, &chef.Experience, &chef.Username, &chef.Password)
	
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if chef.Password != password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"user":    "Welcome back, " + chef.Name + ".",
	})
}
