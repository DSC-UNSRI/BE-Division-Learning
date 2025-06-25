package controllers

import (
	"UTS_BE/database"
	"UTS_BE/middleware"
	"UTS_BE/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	userIDCtx := r.Context().Value(middleware.UserIDKey)
	roleCtx := r.Context().Value(middleware.RoleKey)

	if userIDCtx == nil || roleCtx == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := userIDCtx.(int)
	role := roleCtx.(string)

	// Hitung jumlah pertanyaan hari ini untuk user free
	if role == "free" {
		var count int
		err := database.DB.QueryRow(`
			SELECT COUNT(*) FROM questions 
			WHERE user_id = ? AND DATE(created_at) = CURDATE()
			AND deleted_at IS NULL
		`, userID).Scan(&count)
		if err != nil {
			http.Error(w, "Failed to check question quota", http.StatusInternalServerError)
			return
		}

		if count >= 3 {
			http.Error(w, "You have reached the daily limit of 3 questions", http.StatusForbidden)
			return
		}
	}

	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")
	highlightStr := r.FormValue("highlight")

	if title == "" || content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	highlight := false
	if role == "premium" && highlightStr == "true" {
		highlight = true
	}

	_, err := database.DB.Exec(
		"INSERT INTO questions (user_id, title, content, highlight) VALUES (?, ?, ?, ?)",
		userID, title, content, highlight,
	)
	if err != nil {
		http.Error(w, "Failed to create question", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Question created",
	})
}


func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, user_id, title, content, highlight, created_at FROM questions WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, "Failed to fetch questions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var q models.Question
		err := rows.Scan(&q.ID, &q.UserID, &q.Title, &q.Content, &q.Highlight, &q.CreatedAt)
		if err != nil {
			http.Error(w, "Error scanning result", http.StatusInternalServerError)
			return
		}
		questions = append(questions, q)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"questions": questions,
	})
}

func GetMyQuestions(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)

	rows, err := database.DB.Query(
		"SELECT id, user_id, title, content, created_at FROM questions WHERE user_id = ? AND deleted_at IS NULL", userID,
	)
	if err != nil {
		http.Error(w, "Failed to fetch questions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var q models.Question
		err := rows.Scan(&q.ID, &q.UserID, &q.Title, &q.Content, &q.CreatedAt)
		if err != nil {
			http.Error(w, "Error scanning result", http.StatusInternalServerError)
			return
		}
		questions = append(questions, q)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"my_questions": questions,
	})
}

func GetQuestionByID(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var q models.Question
	err = database.DB.QueryRow(
		"SELECT id, user_id, title, content, created_at FROM questions WHERE id = ? AND deleted_at IS NULL", id,
	).Scan(&q.ID, &q.UserID, &q.Title, &q.Content, &q.CreatedAt)
	if err != nil {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"question": q,
	})
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request, idStr string) {
	userIDCtx := r.Context().Value(middleware.UserIDKey)
	roleCtx := r.Context().Value(middleware.RoleKey)

	if userIDCtx == nil || roleCtx == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := userIDCtx.(int)
	role := roleCtx.(string)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	// Cek apakah user adalah pemilik question
	var ownerID int
	err = database.DB.QueryRow("SELECT user_id FROM questions WHERE id = ? AND deleted_at IS NULL", id).Scan(&ownerID)
	if err != nil {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}

	if ownerID != userID {
		http.Error(w, "Unauthorized: only the owner can update", http.StatusUnauthorized)
		return
	}

	// Ambil data dari form
	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")
	highlightStr := r.FormValue("highlight")

	if title == "" || content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	// Cek apakah user bisa update highlight
	highlight := false
	if role == "premium" && highlightStr == "true" {
		highlight = true
	}

	_, err = database.DB.Exec(`
		UPDATE questions 
		SET title = ?, content = ?, highlight = ? 
		WHERE id = ?`, title, content, highlight, id)
	if err != nil {
		http.Error(w, "Failed to update question", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Question updated successfully",
	})
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request, idStr string) {
	userID := r.Context().Value(middleware.UserIDKey).(int)
	// role := r.Context().Value(middleware.RoleKey).(string)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var ownerID int
	err = database.DB.QueryRow("SELECT user_id FROM questions WHERE id = ? AND deleted_at IS NULL", id).Scan(&ownerID)
	if err != nil {
		http.Error(w, "Question not found or already deleted", http.StatusNotFound)
		return
	}

	if userID != ownerID {
		http.Error(w, "Unauthorized: only owner or premium user can delete", http.StatusUnauthorized)
		return
	}

	_, err = database.DB.Exec("UPDATE questions SET deleted_at = ? WHERE id = ?", time.Now(), id)
	if err != nil {
		http.Error(w, "Failed to delete question", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Question deleted",
	})
}
