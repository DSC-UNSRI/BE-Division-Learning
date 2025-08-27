package controllers

import (
	"database/sql"
	"encoding/json"
	"forum-app/internal/database"
	"forum-app/internal/middlewares"
	"forum-app/internal/models"
	"net/http"
	"strconv"
	"time"
)

const maxFreeQuestionsPerDay = 5

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	claims, _ := middlewares.GetClaimsFromContext(r)

	if claims.UserType == "free" {
		var questionCount int
		err := database.DB.QueryRow(
			"SELECT COUNT(*) FROM questions WHERE user_id = ? AND created_at >= ?",
			claims.UserID, time.Now().Add(-24*time.Hour),
		).Scan(&questionCount)

		if err != nil {
			http.Error(w, "Failed to check user quota", http.StatusInternalServerError)
			return
		}

		if questionCount >= maxFreeQuestionsPerDay {
			http.Error(w, "Free user question limit reached. Please upgrade to premium.", http.StatusTooManyRequests)
			return
		}
	}

	var payload models.CreateQuestionPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if payload.Title == "" || payload.Body == "" {
		http.Error(w, "Title and body are required", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("INSERT INTO questions (user_id, title, body) VALUES (?, ?, ?)",
		claims.UserID, payload.Title, payload.Body)
	if err != nil {
		http.Error(w, "Failed to create question", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Question created successfully"})
}

func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT q.id, q.title, q.body, q.user_id, u.username, q.created_at
		FROM questions q
		JOIN users u ON q.user_id = u.id
		ORDER BY q.created_at DESC
	`)
	if err != nil {
		http.Error(w, "Failed to retrieve questions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	questions := []models.Question{}
	for rows.Next() {
		var q models.Question
		if err := rows.Scan(&q.ID, &q.Title, &q.Body, &q.UserID, &q.Username, &q.CreatedAt); err != nil {
			http.Error(w, "Failed to scan question data", http.StatusInternalServerError)
			return
		}
		questions = append(questions, q)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	var q models.Question
	err = database.DB.QueryRow(`
		SELECT q.id, q.title, q.body, q.user_id, u.username, q.created_at
		FROM questions q JOIN users u ON q.user_id = u.id
		WHERE q.id = ?`, id).Scan(&q.ID, &q.Title, &q.Body, &q.UserID, &q.Username, &q.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Question not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve question", http.StatusInternalServerError)
		return
	}

	rows, err := database.DB.Query(`
		SELECT a.id, a.body, a.question_id, a.user_id, u.username, a.created_at
		FROM answers a JOIN users u ON a.user_id = u.id
		WHERE a.question_id = ? ORDER BY a.created_at ASC`, id)
	if err != nil {
		http.Error(w, "Failed to retrieve answers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var answers []models.Answer
	for rows.Next() {
		var a models.Answer
		if err := rows.Scan(&a.ID, &a.Body, &a.QuestionID, &a.UserID, &a.Username, &a.CreatedAt); err != nil {
			continue
		}
		answers = append(answers, a)
	}

	response := models.QuestionWithAnswers{
		Question: q,
		Answers:  answers,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	claims, _ := middlewares.GetClaimsFromContext(r)
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	var payload models.CreateQuestionPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var ownerID int64
	err = database.DB.QueryRow("SELECT user_id FROM questions WHERE id = ?", id).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Question not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to verify ownership", http.StatusInternalServerError)
		return
	}

	if ownerID != claims.UserID {
		http.Error(w, "Forbidden: You are not the owner of this question", http.StatusForbidden)
		return
	}

	_, err = database.DB.Exec("UPDATE questions SET title = ?, body = ? WHERE id = ?", payload.Title, payload.Body, id)
	if err != nil {
		http.Error(w, "Failed to update question", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Question updated successfully"})
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	claims, _ := middlewares.GetClaimsFromContext(r)
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	var ownerID int64
	err = database.DB.QueryRow("SELECT user_id FROM questions WHERE id = ?", id).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Question not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to verify ownership", http.StatusInternalServerError)
		return
	}

	if ownerID != claims.UserID {
		http.Error(w, "Forbidden: You are not the owner of this question", http.StatusForbidden)
		return
	}

	_, err = database.DB.Exec("DELETE FROM questions WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete question", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}