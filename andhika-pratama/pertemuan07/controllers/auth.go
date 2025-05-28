package controllers

import (
	"pertemuan05/database"
	"pertemuan05/models"

	"encoding/json"
	"net/http"
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

	name := r.FormValue("lecturer_name")
	password := r.FormValue("password")

	if name == "" || password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	var lecturer models.Lecturer
	err = database.DB.QueryRow(`
		SELECT lecturer_id, lecturer_name, password FROM lecturers 
		WHERE lecturer_name = ? AND password = ? AND deleted_at IS NULL
	`, name, password).Scan(&lecturer.LecturerID, &lecturer.LecturerName, &lecturer.Password)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful.",
		"user":    "Welcome back, " + lecturer.LecturerName + ".",
	})
}
