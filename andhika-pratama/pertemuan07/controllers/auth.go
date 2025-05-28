package controllers

import (
	"pertemuan05/database"
	"pertemuan05/models"
	"pertemuan05/utils"

	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)
func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lecturerID := r.FormValue("lecturer_id")
	lecturerName := r.FormValue("lecturer_name")
	password := r.FormValue("password")

	if lecturerID == "" || lecturerName == "" || password == "" {
		http.Error(w, "Missing required fields: lecturer_id, lecturer_name, password", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM lecturers WHERE lecturer_id = ? AND deleted_at IS NULL)", lecturerID).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error checking lecturer ID existence", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Lecturer with this ID already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO lecturers (lecturer_id, lecturer_name, password) VALUES (?, ?, ?)",
		lecturerID, lecturerName, hashedPassword,
	)
	if err != nil {
		http.Error(w, "Failed to register lecturer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Lecturer registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lecturerID := r.FormValue("lecturer_id")
	password := r.FormValue("password")

	if lecturerID == "" ||  password == "" {
		http.Error(w, "Missing required fields: lecturer_id, password", http.StatusBadRequest)
		return
	}

	var lecturer models.Lecturer
	err := database.DB.QueryRow("SELECT lecturer_id, lecturer_name, password, token, role FROM lecturers WHERE lecturer_id = ? AND deleted_at IS NULL", lecturerID).
		Scan(&lecturer.LecturerID, &lecturer.LecturerName, &lecturer.Password, &lecturer.Token, &lecturer.Role)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(lecturer.Password), []byte(password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	newToken := utils.GenerateToken(32)

	_, err = database.DB.Exec("UPDATE lecturers SET token = ? WHERE lecturer_id = ? AND deleted_at IS NULL", newToken, lecturerID)
	if err != nil {
		http.Error(w, "Failed to update token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": newToken,
	})
}
