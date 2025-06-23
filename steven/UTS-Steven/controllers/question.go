package controllers

import (
	"database/sql"
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

func GetQuestion(w http.ResponseWriter, r *http.Request, id string){
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	var question models.Question

	err := database.DB.QueryRow(`SELECT id, user_id, title, content, created_at FROM questions WHERE id = ? AND deleted_at IS NULL`, id).Scan(&question.ID, &question.UserID, &question.Title, &question.Content, &question.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Question not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success Found",
		"question": question,
	})
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request, id string){
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	question := models.Question{}
	err := database.DB.QueryRow("SELECT id, title, content FROM questions WHERE id = ? AND deleted_at IS NULL", id).Scan(&question.ID, &question.Title, &question.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Question not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	if title != "" {
		question.Title = title
	}
	if content != "" {
		question.Content = content
	}

	_, err = database.DB.Exec("UPDATE questions SET title = ?, content = ? WHERE id = ? AND deleted_at IS NULL", question.Title, question.Content, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Question updated successfully",
		"question":    question,
	})
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request, id string){
	if id == "" {		
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM questions WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "question not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE questions SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, "failed to delete question", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Question deleted successfully",
		"id":      id,
	})
}