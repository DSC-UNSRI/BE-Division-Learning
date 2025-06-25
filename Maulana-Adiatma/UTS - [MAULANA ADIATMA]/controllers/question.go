package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"utsquora/database"
	"utsquora/middlewares"
	"utsquora/models"
)

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUser).(models.User)

	var input struct {
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(`
		INSERT INTO questions (user_id, content) VALUES (?, ?)`,
		user.ID, input.Content)

	if err != nil {
		http.Error(w, "Failed to create question", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Question created successfully",
	})
}

func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
	SELECT questions.id, questions.user_id, users.username, questions.content, questions.created_at
	FROM questions 
	JOIN users ON questions.user_id = users.id
	ORDER BY questions.created_at DESC
	`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var q models.Question
		if err := rows.Scan(&q.ID, &q.UserID, &q.Username, &q.Content, &q.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		questions = append(questions, q)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUser).(models.User)

	idStr := strings.TrimPrefix(r.URL.Path, "/question/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	var ownerID int
	err = database.DB.QueryRow("SELECT user_id FROM questions WHERE id = ?", id).Scan(&ownerID)
	if err != nil || ownerID != user.ID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input struct {
		Content string `json:"content"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	_, err = database.DB.Exec("UPDATE questions SET content = ? WHERE id = ?", input.Content, id)
	if err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Question updated"})
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUser).(models.User)

	idStr := strings.TrimPrefix(r.URL.Path, "/question/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	var ownerID int
	err = database.DB.QueryRow("SELECT user_id FROM questions WHERE id = ?", id).Scan(&ownerID)
	if err != nil || ownerID != user.ID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, err = database.DB.Exec("DELETE FROM questions WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Question deleted"})
}