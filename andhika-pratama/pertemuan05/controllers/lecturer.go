package controllers

import (
	"pertemuan05/database"
	"pertemuan05/models"

	"database/sql"
	"encoding/json"
	"net/http"
)

func GetLecturers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT lecturer_id, lecturer_name, password FROM lecturers WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	lecturers := []models.Lecturer{}
	for rows.Next() {
		lecturer := models.Lecturer{}
		rows.Scan(&lecturer.LecturerID, &lecturer.LecturerName, &lecturer.Password)
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
	err := database.DB.QueryRow("SELECT lecturer_id, lecturer_name, password FROM lecturers WHERE lecturer_id = ? AND deleted_at IS NULL", id).
		Scan(&lecturer.LecturerID, &lecturer.LecturerName, &lecturer.Password)

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

	lecturer := models.Lecturer{
		LecturerID:   r.FormValue("lecturer_id"),
		LecturerName: r.FormValue("lecturer_name"),
		Password:     r.FormValue("password"),
	}

	if lecturer.LecturerID == "" || lecturer.LecturerName == "" || lecturer.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("INSERT INTO lecturers (lecturer_id, lecturer_name, password) VALUES (?, ?, ?)",
		lecturer.LecturerID, lecturer.LecturerName, lecturer.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Lecturer successfully created",
		"lecturer": lecturer,
	})
}

func UpdateLecturer(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	lecturer := models.Lecturer{}
	err := database.DB.QueryRow("SELECT lecturer_id, lecturer_name, password FROM lecturers WHERE lecturer_id = ? AND deleted_at IS NULL", id).
		Scan(&lecturer.LecturerID, &lecturer.LecturerName, &lecturer.Password)
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
	if name != "" {
		lecturer.LecturerName = name
	}
	if password != "" {
		lecturer.Password = password
	}

	_, err = database.DB.Exec("UPDATE lecturers SET lecturer_name = ?, password = ? WHERE lecturer_id = ? AND deleted_at IS NULL",
		lecturer.LecturerName, lecturer.Password, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Lecturer updated successfully",
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

	_, err = database.DB.Exec("UPDATE lecturers SET deleted_at = NOW() WHERE lecturer_id = ?", id)
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