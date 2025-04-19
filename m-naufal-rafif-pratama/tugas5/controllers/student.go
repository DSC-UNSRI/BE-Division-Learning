package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas5/database"
	"tugas5/models"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM students")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Pass, &student.Major, &student.Year, &student.OrgID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request, id string) {
	var student models.Student
	err := database.DB.QueryRow("SELECT id, name, email, password, major, year, org_id FROM students WHERE id = ?", id).Scan(
		&student.ID, &student.Name, &student.Email, &student.Pass, &student.Major, &student.Year, &student.OrgID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Student not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !student.IsValid() {
		http.Error(w, "Invalid student data", http.StatusBadRequest)
		return
	}

	result, err := database.DB.Exec("INSERT INTO students (name, email, password, major, year, org_id) VALUES (?, ?, ?, ?, ?, ?)",
		student.Name, student.Email, student.Pass, student.Major, student.Year, student.OrgID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	student.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request, id string) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("UPDATE students SET name = ?, email = ?, password = ?, major = ?, year = ?, org_id = ? WHERE id = ?",
		student.Name, student.Email, student.Pass, student.Major, student.Year, student.OrgID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request, id string) {
	result, err := database.DB.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}