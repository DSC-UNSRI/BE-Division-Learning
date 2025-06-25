package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"uts-gdg/database"
	"uts-gdg/models"
	"uts-gdg/utils"
)

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	answer := models.Answer{}
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

	answer.UserID = userID

	questionIDSTR := r.FormValue("questionID")
	questionIDINT, err := strconv.Atoi(questionIDSTR)
	if err != nil {
		http.Error(w, "Invalid format Question ID", http.StatusBadRequest)
		return
	}
	answer.QuestionID = questionIDINT

	answer.Content = r.FormValue("content")

	if (answer.QuestionID == 0 || answer.Content == "") {
		http.Error(w, "Title and Content cannot be empty", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("INSERT INTO answers (user_id, question_id, content) VALUES (?, ?, ?)", answer.UserID, answer.QuestionID, answer.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	answer.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Question Sended",
		"answer": answer,
	})
}

func GetAnswersByQuestionID(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "question id is null", http.StatusBadRequest)
		return
	}

	rows, err := database.DB.Query(`
	SELECT a.id, a.question_id, a.user_id, u.name AS user_name, a.content, a.created_at 
	FROM answers AS a
	JOIN users AS u
	ON a.user_id = u.id
	WHERE a.question_id = ? 
	AND a.deleted_at IS NULL`, id)

	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()


	answers := []models.AnswerWithUser{}
	for rows.Next() {
		answer := models.AnswerWithUser{}
		rows.Scan(&answer.ID, &answer.QuestionID, &answer.UserID, &answer.UserName, &answer.Content, &answer.CreatedAt)
		answers = append(answers, answer)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Success Found",
		"answers": answers,
	})
}

func UpdateAnswer(w http.ResponseWriter, r *http.Request, id string) {
	answer := models.Answer{}

	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	userIDCtx := r.Context().Value(utils.UserIDKey)
	if userIDCtx == nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return	
    }
	userID := userIDCtx.(int)

	var answerUserID int
	err := database.DB.QueryRow("SELECT content, user_id FROM answers WHERE id = ? AND deleted_at IS NULL", id).Scan(&answer.Content, &answerUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Answer not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if answerUserID != userID {
		http.Error(w, "Forbidden - You can only update your own answer", http.StatusForbidden)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	if content != "" {
		answer.Content = content
	}

	_, err = database.DB.Exec("UPDATE answers SET content = ? WHERE id = ? AND deleted_at IS NULL", answer.Content, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Answer updated successfully",
		"answer":    answer,
	})
}

func DeleteAnswer(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {		
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	userIDCtx := r.Context().Value(utils.UserIDKey)
	if userIDCtx == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := userIDCtx.(int)

	var answerUserID int
	err := database.DB.QueryRow("SELECT user_id FROM answers WHERE id = ? AND deleted_at IS NULL", id).Scan(&answerUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Answer not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	if answerUserID != userID {
		http.Error(w, "Forbidden - You can only delete your own answer", http.StatusForbidden)
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM answers WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "question not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE answers SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, "failed to delete question", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Answer deleted",
		"id":      id,
	})
}