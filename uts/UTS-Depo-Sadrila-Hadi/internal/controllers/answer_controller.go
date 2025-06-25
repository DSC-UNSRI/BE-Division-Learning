package controllers

import (
	"database/sql"
	"encoding/json"
	"forum-app/internal/database"
	"forum-app/internal/middlewares"
	"forum-app/internal/models"
	"net/http"
	"strconv"
)

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	claims, _ := middlewares.GetClaimsFromContext(r)

	questionID, err := strconv.ParseInt(r.PathValue("questionID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	var payload models.AnswerPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if payload.Body == "" {
		http.Error(w, "Answer body cannot be empty", http.StatusBadRequest)
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM questions WHERE id = ?)", questionID).Scan(&exists)
	if err != nil || !exists {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("INSERT INTO answers (question_id, user_id, body) VALUES (?, ?, ?)",
		questionID, claims.UserID, payload.Body)
	if err != nil {
		http.Error(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Answer posted successfully"})
}

func GetAnswersForQuestion(w http.ResponseWriter, r *http.Request) {
	questionID, err := strconv.ParseInt(r.PathValue("questionID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	rows, err := database.DB.Query(`
		SELECT a.id, a.body, a.question_id, a.user_id, u.username, a.created_at
		FROM answers a
		JOIN users u ON a.user_id = u.id
		WHERE a.question_id = ?
		ORDER BY a.created_at ASC
	`, questionID)
	if err != nil {
		http.Error(w, "Failed to retrieve answers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	answers := []models.Answer{}
	for rows.Next() {
		var a models.Answer
		if err := rows.Scan(&a.ID, &a.Body, &a.QuestionID, &a.UserID, &a.Username, &a.CreatedAt); err != nil {
			http.Error(w, "Failed to scan answer data", http.StatusInternalServerError)
			return
		}
		answers = append(answers, a)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answers)
}

func UpdateAnswer(w http.ResponseWriter, r *http.Request) {
	claims, _ := middlewares.GetClaimsFromContext(r)
	answerID, err := strconv.ParseInt(r.PathValue("answerID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	var payload models.AnswerPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var ownerID int64
	err = database.DB.QueryRow("SELECT user_id FROM answers WHERE id = ?", answerID).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Answer not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to verify ownership", http.StatusInternalServerError)
		return
	}

	if ownerID != claims.UserID {
		http.Error(w, "Forbidden: You are not the owner of this answer", http.StatusForbidden)
		return
	}

	_, err = database.DB.Exec("UPDATE answers SET body = ? WHERE id = ?", payload.Body, answerID)
	if err != nil {
		http.Error(w, "Failed to update answer", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Answer updated successfully"})
}

func DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	claims, _ := middlewares.GetClaimsFromContext(r)
	answerID, err := strconv.ParseInt(r.PathValue("answerID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	var ownerID int64
	err = database.DB.QueryRow("SELECT user_id FROM answers WHERE id = ?", answerID).Scan(&ownerID)
	if err != nil {
		http.Error(w, "Answer not found", http.StatusNotFound)
		return
	}

	if ownerID != claims.UserID {
		http.Error(w, "Forbidden: You are not the owner of this answer", http.StatusForbidden)
		return
	}

	_, err = database.DB.Exec("DELETE FROM answers WHERE id = ?", answerID)
	if err != nil {
		http.Error(w, "Failed to delete answer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}