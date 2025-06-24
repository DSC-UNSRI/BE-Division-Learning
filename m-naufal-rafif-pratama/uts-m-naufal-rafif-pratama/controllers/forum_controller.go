package controllers

import (
	"database/sql"
	"encoding/json"
	"uts/database"
	"uts/middleware"
	"uts/models"
	"net/http"
	"strconv"
	"time"
)

type CreateQuestionRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	user, ok := middleware.GetUserFromContext(r.Context())
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	if user.Role == "free" {
		var questionCount int
		query := "SELECT COUNT(*) FROM questions WHERE user_id = ? AND created_at >= ?"
		err := database.DB.QueryRow(query, user.ID, time.Now().Add(-24*time.Hour)).Scan(&questionCount)
		if err != nil {
			http.Error(w, "Database error on question count", http.StatusInternalServerError)
			return
		}

		if questionCount >= 1 {
			http.Error(w, "Free users are limited to 1 question per day. Go Premium!", http.StatusForbidden)
			return
		}
	}

	var req CreateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Body == "" {
		http.Error(w, "Title and body are required", http.StatusBadRequest)
		return
	}

	newQuestion := models.Question{
		UserID:    user.ID,
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: time.Now(),
	}

	result, err := database.DB.Exec("INSERT INTO questions (user_id, title, body, created_at) VALUES (?, ?, ?, ?)",
		newQuestion.UserID, newQuestion.Title, newQuestion.Body, newQuestion.CreatedAt)
	if err != nil {
		http.Error(w, "Failed to create question", http.StatusInternalServerError)
		return
	}
	questionID, _ := result.LastInsertId()
	newQuestion.ID = int(questionID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newQuestion)
}

func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, user_id, title, body, created_at FROM questions ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	questions := make([]models.Question, 0)
	for rows.Next() {
		var q models.Question
		if err := rows.Scan(&q.ID, &q.UserID, &q.Title, &q.Body, &q.CreatedAt); err != nil {
			http.Error(w, "Failed to scan question data", http.StatusInternalServerError)
			return
		}
		questions = append(questions, q)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	questionIDStr := r.URL.Query().Get("id")
	if questionIDStr == "" {
		http.Error(w, "Question ID is required", http.StatusBadRequest)
		return
	}
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		http.Error(w, "Question ID must be an integer", http.StatusBadRequest)
		return
	}

	var question models.Question
	err = database.DB.QueryRow("SELECT id, user_id, title, body, created_at FROM questions WHERE id = ?", questionID).Scan(
		&question.ID, &question.UserID, &question.Title, &question.Body, &question.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Question not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	rows, err := database.DB.Query("SELECT id, user_id, question_id, body, created_at FROM answers WHERE question_id = ? ORDER BY created_at ASC", questionID)
	if err != nil {
		http.Error(w, "Database error fetching answers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	answers := make([]models.Answer, 0)
	for rows.Next() {
		var a models.Answer
		if err := rows.Scan(&a.ID, &a.UserID, &a.QuestionID, &a.Body, &a.CreatedAt); err != nil {
			http.Error(w, "Failed to scan answer data", http.StatusInternalServerError)
			return
		}
		answers = append(answers, a)
	}

	response := map[string]interface{}{
		"question": question,
		"answers":  answers,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type CreateAnswerRequest struct {
	QuestionID int    `json:"question_id"`
	Body       string `json:"body"`
}

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	user, ok := middleware.GetUserFromContext(r.Context())
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	if user.Role == "free" {
		var answerCount int
		query := "SELECT COUNT(*) FROM answers WHERE user_id = ? AND created_at >= ?"
		err := database.DB.QueryRow(query, user.ID, time.Now().Add(-24*time.Hour)).Scan(&answerCount)
		if err != nil {
			http.Error(w, "Database error on answer count", http.StatusInternalServerError)
			return
		}

		if answerCount >= 5 {
			http.Error(w, "Free users are limited to 5 answers per day. Go Premium!", http.StatusForbidden)
			return
		}
	}

	var req CreateAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.QuestionID == 0 || req.Body == "" {
		http.Error(w, "Question ID and body are required", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM questions WHERE id = ?)", req.QuestionID).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error checking question", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}

	newAnswer := models.Answer{
		UserID:     user.ID,
		QuestionID: req.QuestionID,
		Body:       req.Body,
		CreatedAt:  time.Now(),
	}

	result, err := database.DB.Exec("INSERT INTO answers (user_id, question_id, body, created_at) VALUES (?, ?, ?, ?)",
		newAnswer.UserID, newAnswer.QuestionID, newAnswer.Body, newAnswer.CreatedAt)
	if err != nil {
		http.Error(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}
	answerID, _ := result.LastInsertId()
	newAnswer.ID = int(answerID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAnswer)
} 