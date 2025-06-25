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

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	userIDCtx := r.Context().Value(middleware.UserIDKey)
	if userIDCtx == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := userIDCtx.(int)

	r.ParseForm()
	questionIDStr := r.FormValue("question_id")
	content := r.FormValue("content")

	if questionIDStr == "" || content == "" {
		http.Error(w, "question_id and content are required", http.StatusBadRequest)
		return
	}

	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		http.Error(w, "Invalid question_id", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		"INSERT INTO answers (question_id, user_id, content) VALUES (?, ?, ?)",
		questionID, userID, content,
	)
	if err != nil {
		http.Error(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Answer submitted successfully",
	})
}

func GetAnswersByQuestion(w http.ResponseWriter, r *http.Request, questionIDStr string) {
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	rows, err := database.DB.Query(`
		SELECT a.id, a.question_id, a.user_id, a.content, a.created_at, u.role
		FROM answers a
		JOIN users u ON a.user_id = u.id
		WHERE a.question_id = ? AND a.deleted_at IS NULL
	`, questionID)
	if err != nil {
		http.Error(w, "Failed to retrieve answers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var answers []models.Answer
	for rows.Next() {
		var a models.Answer
		var role string

		err := rows.Scan(&a.ID, &a.QuestionID, &a.UserID, &a.Content, &a.CreatedAt, &role)
		if err != nil {
			http.Error(w, "Error while reading data", http.StatusInternalServerError)
			return
		}

		if role == "premium" {
			a.Highlight = true
		} else {
			a.Highlight = false
		}

		answers = append(answers, a)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"answers": answers,
	})
}

func UpdateAnswer(w http.ResponseWriter, r *http.Request, idStr string) {
	userIDCtx := r.Context().Value(middleware.UserIDKey)
	if userIDCtx == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := userIDCtx.(int)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	// Cek apakah user adalah pemilik jawaban
	var ownerID int
	err = database.DB.QueryRow("SELECT user_id FROM answers WHERE id = ? AND deleted_at IS NULL", id).Scan(&ownerID)
	if err != nil {
		http.Error(w, "Answer not found", http.StatusNotFound)
		return
	}

	if ownerID != userID {
		http.Error(w, "Unauthorized: only the owner can update", http.StatusUnauthorized)
		return
	}

	r.ParseForm()
	content := r.FormValue("content")
	if content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE answers SET content = ? WHERE id = ?", content, id)
	if err != nil {
		http.Error(w, "Failed to update answer", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Answer updated successfully",
	})
}

func DeleteAnswer(w http.ResponseWriter, r *http.Request, answerIDStr string) {
	userIDCtx := r.Context().Value(middleware.UserIDKey)
	if userIDCtx == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := userIDCtx.(int)

	answerID, err := strconv.Atoi(answerIDStr)
	if err != nil {
		http.Error(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	var ownerID int
	err = database.DB.QueryRow("SELECT user_id FROM answers WHERE id = ?", answerID).Scan(&ownerID)
	if err != nil || ownerID != userID {
		http.Error(w, "Unauthorized: You can only delete your own answer", http.StatusUnauthorized)
		return
	}

	_, err = database.DB.Exec("UPDATE answers SET deleted_at = ? WHERE id = ?", time.Now(), answerID)
	if err != nil {
		http.Error(w, "Failed to delete answer", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Answer deleted successfully",
	})
}
