package controllers

import (
	"pertemuan05/database"
	"pertemuan05/models"
	"pertemuan05/utils"

	"encoding/json"
	"net/http"
)

func GetCourses(w http.ResponseWriter, r *http.Request) {
	ctxLecturerID := r.Context().Value(utils.LecturerIDKey).(string)

	rows, err := database.DB.Query("SELECT course_id, course_name, lecturer_id, semester, credit FROM courses WHERE lecturer_id = ? AND deleted_at IS NULL", ctxLecturerID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	courses := []models.Course{}
	for rows.Next() {
		course := models.Course{}
		rows.Scan(&course.CourseID, &course.CourseName, &course.LecturerID, &course.Semester, &course.Credit)
		courses = append(courses, course)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"courses": courses,
	})
}

func GetCourseByID(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	course := models.Course{}
	err := database.DB.QueryRow("SELECT course_id, course_name, lecturer_id, semester, credit FROM courses WHERE course_id = ? AND deleted_at IS NULL", id).
		Scan(&course.CourseID, &course.CourseName, &course.LecturerID, &course.Semester, &course.Credit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	ctxLecturerID := r.Context().Value(utils.LecturerIDKey).(string)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	course := models.Course{
		CourseID:   r.FormValue("course_id"),
		CourseName: r.FormValue("course_name"),
		LecturerID: ctxLecturerID,
		Semester:   utils.Atoi(r.FormValue("semester")),
		Credit:     utils.Atoi(r.FormValue("credit")),
	}

	if course.CourseID == "" || course.CourseName == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM lecturers WHERE lecturer_id = ? AND deleted_at IS NULL)", course.LecturerID).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error while validating lecturer", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Lecturer not found", http.StatusBadRequest)
		return
	}

	var courseExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM courses WHERE course_id = ? AND deleted_at IS NULL)", course.CourseID).Scan(&courseExists)
	if err != nil {
		http.Error(w, "Database error while checking for existing course", http.StatusInternalServerError)
		return
	}
	if courseExists {
		http.Error(w, "Course with the same course_id already exists", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(`
		INSERT INTO courses (course_id, course_name, lecturer_id, semester, credit)
		VALUES (?, ?, ?, ?, ?)`,
		course.CourseID, course.CourseName, course.LecturerID, course.Semester, course.Credit)

	if err != nil {
		http.Error(w, "Failed to create course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Course created successfully",
		"course":  course,
	})
}

func UpdateCourse(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	ctxLecturerID := r.Context().Value(utils.LecturerIDKey).(string)

	course := models.Course{}
	err := database.DB.QueryRow("SELECT course_id, course_name, lecturer_id, semester, credit FROM courses WHERE course_id = ? AND lecturer_id = ? AND deleted_at IS NULL", id, ctxLecturerID).
		Scan(&course.CourseID, &course.CourseName, &course.LecturerID, &course.Semester, &course.Credit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	if name := r.FormValue("course_name"); name != "" {
		course.CourseName = name
	}
	if sem := r.FormValue("semester"); sem != "" {
		course.Semester = utils.Atoi(sem)
	}
	if credit := r.FormValue("credit"); credit != "" {
		course.Credit = utils.Atoi(credit)
	}

	_, err = database.DB.Exec(`
		UPDATE courses
		SET course_name = ?, semester = ?, credit = ?
		WHERE course_id = ? AND lecturer_id = ? AND deleted_at IS NULL`,
		course.CourseName, course.Semester, course.Credit, id, ctxLecturerID)

	if err != nil {
		http.Error(w, "Failed to update course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Course updated successfully",
		"course":  course,
	})
}

func DeleteCourse(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	ctxLecturerID := r.Context().Value(utils.LecturerIDKey).(string)

	_, err := database.DB.Exec("UPDATE courses SET deleted_at = NOW() WHERE course_id = ? AND lecturer_id = ?", id, ctxLecturerID)
	if err != nil {
		http.Error(w, "Failed to delete course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Course deleted successfully",
		"id":      id,
	})
}

func GetCoursesByLecturer(w http.ResponseWriter, r *http.Request, lecturerID string) {
	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM lecturers WHERE lecturer_id = ? AND deleted_at IS NULL)", lecturerID).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error while validating lecturer", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Lecturer not found", http.StatusNotFound)
		return
	}

	rows, err := database.DB.Query(`
		SELECT course_id, course_name, lecturer_id, semester, credit 
		FROM courses 
		WHERE lecturer_id = ? AND deleted_at IS NULL`, lecturerID)
	if err != nil {
		http.Error(w, "Database error while retrieving courses", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	courses := []models.Course{}
	for rows.Next() {
		course := models.Course{}
		err := rows.Scan(&course.CourseID, &course.CourseName, &course.LecturerID, &course.Semester, &course.Credit)
		if err != nil {
			http.Error(w, "Failed to scan course data", http.StatusInternalServerError)
			return
		}
		courses = append(courses, course)
	}

	if len(courses) == 0 {
		http.Error(w, "Lecturer has no assigned courses", http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"courses":     courses,
	})
}
