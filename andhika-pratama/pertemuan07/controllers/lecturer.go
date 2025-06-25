package controllers

import (
	"pertemuan05/database"
	"pertemuan05/models"

	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GetLecturers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT lecturer_id, lecturer_name, password, role, deleted_at FROM lecturers WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	lecturers := []models.Lecturer{}
	for rows.Next() {
		lecturer := models.Lecturer{}
		err := rows.Scan(&lecturer.LecturerID, &lecturer.LecturerName, &lecturer.Password, &lecturer.Role, &lecturer.DeletedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		lecturers = append(lecturers, lecturer)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"lecturers": lecturers,
	})
}

func GetLecturerByID(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	lecturer := models.Lecturer{}
	err := database.DB.QueryRow("SELECT lecturer_id, lecturer_name, password, role, deleted_at FROM lecturers WHERE lecturer_id = ? AND deleted_at IS NULL", id).
		Scan(&lecturer.LecturerID, &lecturer.LecturerName, &lecturer.Password, &lecturer.Role, &lecturer.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Lecturer not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"lecturer": lecturer,
	})
}

func CreateLecturer(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	lecturerID := r.FormValue("lecturer_id")
	lecturerName := r.FormValue("lecturer_name")
	password := r.FormValue("password")
	role := r.FormValue("role")

	if lecturerID == "" || lecturerName == "" || password == "" {
		http.Error(w, "All fields (lecturer_id, lecturer_name, password) are required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	if role == "" {
		role = "new"
	} else if role != "old" && role != "new" {
		http.Error(w, "Invalid role. Role must be 'old' or 'new'", http.StatusBadRequest)
		return
	}

	lecturer := models.Lecturer{
		LecturerID:   lecturerID,
		LecturerName: lecturerName,
		Password:     string(hashedPassword),
		Role:         role,
	}

	var lecturerExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM lecturers WHERE lecturer_id = ?)", lecturer.LecturerID).Scan(&lecturerExists)
	if err != nil {
		http.Error(w, "Database error while checking for existing lecturer", http.StatusInternalServerError)
		return
	}
	if lecturerExists {
		http.Error(w, "Lecturer with this ID already exists", http.StatusConflict)
		return
	}

	_, err = database.DB.Exec("INSERT INTO lecturers (lecturer_id, lecturer_name, password, token, role) VALUES (?, ?, ?, ?, ?)",
		lecturer.LecturerID, lecturer.LecturerName, lecturer.Password, lecturer.Token, lecturer.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Lecturer successfully created",
		"lecturer": lecturer,
	})
}

func UpdateLecturer(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	lecturer := models.Lecturer{}
	err := database.DB.QueryRow("SELECT lecturer_id, lecturer_name, password, token, role, deleted_at FROM lecturers WHERE lecturer_id = ? AND deleted_at IS NULL", id).
		Scan(&lecturer.LecturerID, &lecturer.LecturerName, &lecturer.Password, &lecturer.Token, &lecturer.Role, &lecturer.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Lecturer not found", http.StatusNotFound)
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

	name := r.FormValue("lecturer_name")
	password := r.FormValue("password")
	role := r.FormValue("role")

	updateFields := []string{}
	updateValues := []interface{}{}

	if name != "" {
		lecturer.LecturerName = name
		updateFields = append(updateFields, "lecturer_name = ?")
		updateValues = append(updateValues, name)
	}
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		lecturer.Password = string(hashedPassword)
		updateFields = append(updateFields, "password = ?")
		updateValues = append(updateValues, hashedPassword)
	}
	if role != "" {
		if role != "old" && role != "new" {
			http.Error(w, "Invalid role. Role must be 'old' or 'new'", http.StatusBadRequest)
			return
		}
		lecturer.Role = role
		updateFields = append(updateFields, "role = ?")
		updateValues = append(updateValues, role)
	}

	if len(updateFields) == 0 {
		http.Error(w, "No fields provided for update", http.StatusBadRequest)
		return
	}

	query := "UPDATE lecturers SET " + strings.Join(updateFields, ", ") + " WHERE lecturer_id = ? AND deleted_at IS NULL"
	updateValues = append(updateValues, id)

	_, err = database.DB.Exec(query, updateValues...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Lecturer updated successfully",
		"lecturer": lecturer,
	})
}

func DeleteLecturer(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM courses WHERE lecturer_id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "Cannot delete lecturer: assigned to a course", http.StatusBadRequest)
		return
	}

	var lecturerExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM lecturers WHERE lecturer_id = ? AND deleted_at IS NULL)", id).Scan(&lecturerExists)
	if err != nil {
		http.Error(w, "Database error while checking lecturer existence", http.StatusInternalServerError)
		return
	}
	if !lecturerExists {
		http.Error(w, "Lecturer not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE lecturers SET token = '', deleted_at = NOW() WHERE lecturer_id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete lecturer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Lecturer deleted successfully",
		"id":      id,
	})
}

func GetLecturersByCity(w http.ResponseWriter, r *http.Request, city string) {
	rows, err := database.DB.Query(`
		SELECT l.lecturer_id, l.lecturer_name, l.password, l.token, l.role, l.deleted_at
		FROM lecturers l
		JOIN addresses a ON l.lecturer_id = a.lecturer_id
		WHERE a.city = ? AND a.deleted_at IS NULL AND l.deleted_at IS NULL`, city)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	lecturers := []models.Lecturer{}
	for rows.Next() {
		lecturer := models.Lecturer{}
		if err := rows.Scan(&lecturer.LecturerID, &lecturer.LecturerName, &lecturer.Password, &lecturer.Token, &lecturer.Role, &lecturer.DeletedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		lecturers = append(lecturers, lecturer)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(lecturers) == 0 {
		http.Error(w, "There are no lecturers in this city", http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"lecturers": lecturers,
	})
}