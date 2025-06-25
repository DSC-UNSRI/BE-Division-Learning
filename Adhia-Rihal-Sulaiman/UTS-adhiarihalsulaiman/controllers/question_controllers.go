package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
	"uts_adhia/database"
	"uts_adhia/models"
	"uts_adhia/utils"
)

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	ctxUserIDStr, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok || ctxUserIDStr == "" {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	ctxUserID, err := strconv.Atoi(ctxUserIDStr)
	if err != nil {
		http.Error(w, "Invalid User ID in context", http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	if title == "" || content == "" {
		http.Error(w, "Missing required fields: title, content", http.StatusBadRequest)
		return
	}

	question := models.Question{
		UserID:       ctxUserID,
		Title:        title,
		Content:      content,
		BestQuestion: false,
	}

	res, err := database.DB.Exec(`
		INSERT INTO questions (user_id, title, content, best_question)
		VALUES (?, ?, ?, ?)`,
		question.UserID, question.Title, question.Content, question.BestQuestion)

	if err != nil {
		http.Error(w, "Failed to create question", http.StatusInternalServerError)
		return
	}

	newQuestionID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Question created, but failed to retrieve Question ID", http.StatusInternalServerError)
		return
	}
	question.ID = int(newQuestionID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Question created successfully",
		"question": question,
	})
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT question_id, user_id, title, content, best_question, created_at, updated_at, deleted_at FROM questions WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	questions := []models.Question{}
	for rows.Next() {
		question := models.Question{}
		err := rows.Scan(&question.ID, &question.UserID, &question.Title, &question.Content, &question.BestQuestion, &question.CreatedAt, &question.UpdatedAt, &question.DeletedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		questions = append(questions, question)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"questions": questions,
	})
}

func GetQuestionByID(w http.ResponseWriter, r *http.Request, questionID string) {
	if questionID == "" {
		http.Error(w, "Question ID missing in URL", http.StatusBadRequest)
		return
	}

	qIDInt, err := strconv.Atoi(questionID)
	if err != nil {
		http.Error(w, "Invalid Question ID", http.StatusBadRequest)
		return
	}

	question := models.Question{}
	err = database.DB.QueryRow("SELECT question_id, user_id, title, content, best_question, created_at, updated_at, deleted_at FROM questions WHERE question_id = ? AND deleted_at IS NULL", qIDInt).
		Scan(&question.ID, &question.UserID, &question.Title, &question.Content, &question.BestQuestion, &question.CreatedAt, &question.UpdatedAt, &question.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Question not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request, questionID string) {
	if questionID == "" {
		http.Error(w, "Question ID missing in URL", http.StatusBadRequest)
		return
	}

	qIDInt, err := strconv.Atoi(questionID)
	if err != nil {
		http.Error(w, "Invalid Question ID", http.StatusBadRequest)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	updateFields := []string{}
	updateValues := []interface{}{}

	if title != "" {
		updateFields = append(updateFields, "title = ?")
		updateValues = append(updateValues, title)
	}

	if content != "" {
		updateFields = append(updateFields, "content = ?")
		updateValues = append(updateValues, content)
	}

	if len(updateFields) == 0 {
		http.Error(w, "No fields provided for update", http.StatusBadRequest)
		return
	}

	query := "UPDATE questions SET " + strings.Join(updateFields, ", ") + ", updated_at = ? WHERE question_id = ? AND deleted_at IS NULL"
	updateValues = append(updateValues, time.Now())
	updateValues = append(updateValues, qIDInt)

	_, err = database.DB.Exec(query, updateValues...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	question := models.Question{}
	err = database.DB.QueryRow("SELECT question_id, user_id, title, content, best_question, created_at, updated_at, deleted_at FROM questions WHERE question_id = ? AND deleted_at IS NULL", qIDInt).
		Scan(&question.ID, &question.UserID, &question.Title, &question.Content, &question.BestQuestion, &question.CreatedAt, &question.UpdatedAt, &question.DeletedAt)

	if err != nil {
		http.Error(w, "Failed to fetch updated question", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Question updated successfully",
		"question": question,
	})
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request, questionID string) {
	if questionID == "" {
		http.Error(w, "Question ID missing in URL", http.StatusBadRequest)
		return
	}

	qIDInt, err := strconv.Atoi(questionID)
	if err != nil {
		http.Error(w, "Invalid Question ID", http.StatusBadRequest)
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE questions SET deleted_at = CURRENT_TIMESTAMP WHERE question_id = ?", qIDInt)
	if err != nil {
		http.Error(w, "Failed to delete question", http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("UPDATE answers SET deleted_at = CURRENT_TIMESTAMP WHERE question_id = ?", qIDInt)
	if err != nil {
		http.Error(w, "Failed to delete answers for this question", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Question deleted successfully",
	})
}

func MarkQuestionBest(w http.ResponseWriter, r *http.Request, questionID string) {
	if questionID == "" {
		http.Error(w, "Question ID missing in URL", http.StatusBadRequest)
		return
	}

	qIDInt, err := strconv.Atoi(questionID)
	if err != nil {
		http.Error(w, "Invalid Question ID", http.StatusBadRequest)
		return
	}

	ctxUserIDStr, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok || ctxUserIDStr == "" {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	ctxUserID, err := strconv.Atoi(ctxUserIDStr)
	if err != nil {
		http.Error(w, "Invalid User ID in context", http.StatusInternalServerError)
		return
	}

	var questionOwnerID int
	err = database.DB.QueryRow("SELECT question_id FROM questions WHERE question_id = ? AND deleted_at IS NULL", qIDInt).Scan(&questionOwnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Question not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var userRole string
	userRole, ok = r.Context().Value(utils.UserRoleKey).(string)
	if !ok {
		http.Error(w, "User role not found in context", http.StatusInternalServerError)
		return
	}

	if ctxUserID != questionOwnerID && userRole != "admin" {
		http.Error(w, "Unauthorized: Only question owner or admin can mark best question", http.StatusForbidden)
		return
	}

	isBestStr := r.FormValue("is_best")
	var setBest bool
	switch isBestStr {
	case "true":
		setBest = true
	case "false":
		setBest = false
	default:
		http.Error(w, "Invalid or missing 'is_best' field. Must be 'true' or 'false'.", http.StatusBadRequest)
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE questions SET best_question = ? WHERE question_id = ?", setBest, qIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Question best status updated successfully",
	})
}
