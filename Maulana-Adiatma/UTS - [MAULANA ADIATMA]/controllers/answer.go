package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"utsquora/database"
	"utsquora/middlewares"
	"utsquora/models"
)

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUser).(models.User)

	var input struct {
		QuestionID int    `json:"question_id"`
		Content    string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if input.Content == "" || input.QuestionID == 0 {
		http.Error(w, "Content and question_id are required", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(`
		INSERT INTO answers (question_id, user_id, content)
		VALUES (?, ?, ?)`,
		input.QuestionID, user.ID, input.Content,
	)

	if err != nil {
		http.Error(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Answer created successfully",
	})
}

func GetAnswersByQuestionID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	questionID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	rows, err := database.DB.Query(`
		SELECT a.id, a.question_id, a.user_id, u.username, a.content, a.created_at
		FROM answers a
		JOIN users u ON a.user_id = u.id
		WHERE a.question_id = ?
		ORDER BY a.created_at ASC
	`, questionID)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var answers []models.Answer
	for rows.Next() {
		var ans models.Answer
		err := rows.Scan(&ans.ID, &ans.QuestionID, &ans.UserID, &ans.Username, &ans.Content, &ans.CreatedAt)
		if err != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		answers = append(answers, ans)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answers)
}