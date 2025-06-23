package controllers

import (
	"encoding/json"
	"net/http"
	"uts-gdg/database"
	"uts-gdg/models"
	"uts-gdg/utils"
)

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, user_id, title, content, created_at FROM questions WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	questions := []models.Question{}
	for rows.Next() {
		question := models.Question{}
		rows.Scan(&question.ID, &question.UserID, &question.Title, &question.Content, &question.CreatedAt)
		questions = append(questions, question)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message" : "Success",
		"questions": questions,
	})
}

func CreateQuestions(w http.ResponseWriter, r *http.Request) {
	question := models.Question{}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	userIDCtx := r.Context().Value(utils.UserIDKey)
	if userIDCtx == nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return	
    }

	userID, ok := userIDCtx.(int)
    if !ok {
        http.Error(w, "Invalid user ID", http.StatusInternalServerError)
        return
    }

	question.UserID = userID
	question.Title = r.FormValue("title")
	question.Content = r.FormValue("content")

	if (question.Title == "" || question.Content == "") {
		http.Error(w, "Title and Content cannot be empty", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("INSERT INTO questions (user_id, title, content) VALUES (?, ?, ?)", question.UserID, question.Title, question.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	question.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Question Created",
		"question": question,
	})
}