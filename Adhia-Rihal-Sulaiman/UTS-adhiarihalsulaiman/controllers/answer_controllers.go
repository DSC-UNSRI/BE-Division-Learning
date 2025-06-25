package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"uts_adhia/database"
	"uts_adhia/models"
	"uts_adhia/utils"
)

func CreateAnswer(w http.ResponseWriter, r *http.Request, questionID string) {
	if questionID == "" {
		http.Error(w, "Please input question_id in the url", http.StatusBadRequest)
		return
	}

	qIDInt, err := strconv.Atoi(questionID)
	if err != nil {
		http.Error(w, "Invalid Question ID", http.StatusBadRequest)
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM questions WHERE question_id = ? AND deleted_at IS NULL)", qIDInt).Scan(&exists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}

	ctxUserIDStr, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok || ctxUserIDStr == "" {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	userIDInt, err := strconv.Atoi(ctxUserIDStr)
	if err != nil {
		http.Error(w, "Invalid User ID in context", http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	if content == "" {
		http.Error(w, "Missing required fields: content", http.StatusBadRequest)
		return
	}

	answer := models.Answer{
		UserID:     userIDInt,
		QuestionID: qIDInt,
		Content:    content,
		BestAnswer: false,
	}

	res, err := database.DB.Exec(`
		INSERT INTO answers (question_id, user_id, content, best_answer)
		VALUES (?, ?, ?, ?)`,
		answer.QuestionID, answer.UserID, answer.Content, answer.BestAnswer)

	if err != nil {
		http.Error(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}

	newAnswerID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Answer created, but failed to retrieve Answer ID", http.StatusInternalServerError)
		return
	}
	answer.ID = int(newAnswerID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Answer created successfully",
		"answer":  answer,
	})
}

func GetAnswersByQuestionID(w http.ResponseWriter, r *http.Request, questionID string) {
	if questionID == "" {
		http.Error(w, "Please input question_id in the url", http.StatusBadRequest)
		return
	}

	qIDInt, err := strconv.Atoi(questionID)
	if err != nil {
		http.Error(w, "Invalid Question ID", http.StatusBadRequest)
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM questions WHERE question_id = ? AND deleted_at IS NULL)", qIDInt).Scan(&exists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}

	rows, err := database.DB.Query("SELECT answer_id, question_id, user_id, content, best_answer, created_at, updated_at, deleted_at FROM answers WHERE question_id = ? AND deleted_at IS NULL", qIDInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	answers := []models.Answer{}
	for rows.Next() {
		answer := models.Answer{}
		err := rows.Scan(&answer.ID, &answer.QuestionID, &answer.UserID, &answer.Content, &answer.BestAnswer, &answer.CreatedAt, &answer.UpdatedAt, &answer.DeletedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		answers = append(answers, answer)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"answers": answers,
	})
}

func GetAnswerByAnswerID(w http.ResponseWriter, r *http.Request, answerID string) {
	if answerID == "" {
		http.Error(w, "Please input answer_id in the url", http.StatusBadRequest)
		return
	}

	aIDInt, err := strconv.Atoi(answerID)
	if err != nil {
		http.Error(w, "Invalid Answer ID", http.StatusBadRequest)
		return
	}

	answer := models.Answer{}
	err = database.DB.QueryRow("SELECT answer_id, question_id, user_id, content, best_answer, created_at, updated_at, deleted_at FROM answers WHERE answer_id = ? AND deleted_at IS NULL", aIDInt).
		Scan(&answer.ID, &answer.QuestionID, &answer.UserID, &answer.Content, &answer.BestAnswer, &answer.CreatedAt, &answer.UpdatedAt, &answer.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Answer not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)
}

func UpdateAnswer(w http.ResponseWriter, r *http.Request, answerID string) {
	if answerID == "" {
		http.Error(w, "Please input answer_id in the url", http.StatusBadRequest)
		return
	}

	aIDInt, err := strconv.Atoi(answerID)
	if err != nil {
		http.Error(w, "Invalid Answer ID", http.StatusBadRequest)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")

	updateFields := []string{}
	updateValues := []interface{}{}

	if content != "" {
		updateFields = append(updateFields, "content = ?")
		updateValues = append(updateValues, content)
	}

	if len(updateFields) == 0 {
		http.Error(w, "No fields provided for update", http.StatusBadRequest)
		return
	}

	query := "UPDATE answers SET " + strings.Join(updateFields, ", ") + ", updated_at = CURRENT_TIMESTAMP WHERE answer_id = ? AND deleted_at IS NULL"
	updateValues = append(updateValues, aIDInt)

	_, err = database.DB.Exec(query, updateValues...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	answer := models.Answer{}
	err = database.DB.QueryRow("SELECT answer_id, question_id, user_id, content, best_answer, created_at, updated_at, deleted_at FROM answers WHERE answer_id = ? AND deleted_at IS NULL", aIDInt).
		Scan(&answer.ID, &answer.QuestionID, &answer.UserID, &answer.Content, &answer.BestAnswer, &answer.CreatedAt, &answer.UpdatedAt, &answer.DeletedAt)

	if err != nil {
		http.Error(w, "Failed to fetch updated Answer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Answer updated successfully",
		"answer":  answer,
	})
}

func DeleteAnswer(w http.ResponseWriter, r *http.Request, answerID string) {
	if answerID == "" {
		http.Error(w, "Please input answer_id in the url", http.StatusBadRequest)
		return
	}

	aIDInt, err := strconv.Atoi(answerID)
	if err != nil {
		http.Error(w, "Invalid Answer ID", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE answers SET deleted_at = CURRENT_TIMESTAMP WHERE answer_id = ?", aIDInt)
	if err != nil {
		http.Error(w, "Failed to delete answer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Answer deleted successfully",
	})
}

func MarkAnswerBest(w http.ResponseWriter, r *http.Request, answerID string) {
	if answerID == "" {
		http.Error(w, "Answer ID missing in URL", http.StatusBadRequest)
		return
	}

	aIDInt, err := strconv.Atoi(answerID)
	if err != nil {
		http.Error(w, "Invalid Answer ID", http.StatusBadRequest)
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

	var qID int
	var answerUserID int
	err = database.DB.QueryRow("SELECT question_id, user_id FROM answers WHERE answer_id = ? AND deleted_at IS NULL", aIDInt).Scan(&qID, &answerUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Answer not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var questionOwnerID int
	err = database.DB.QueryRow("SELECT user_id FROM questions WHERE question_id = ? AND deleted_at IS NULL", qID).Scan(&questionOwnerID)
	if err != nil {
		http.Error(w, "Question not found for answer", http.StatusInternalServerError)
		return
	}

	var userRole string
	userRole, ok = r.Context().Value(utils.UserRoleKey).(string)
	if !ok {
		http.Error(w, "User role not found in context", http.StatusInternalServerError)
		return
	}

	if ctxUserID != questionOwnerID && userRole != "admin" {
		http.Error(w, "Unauthorized: Only question owner or admin can mark best answer", http.StatusForbidden)
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

	_, err = tx.Exec("UPDATE answers SET best_answer = FALSE WHERE question_id = ?", qID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = tx.Exec("UPDATE answers SET best_answer = ? WHERE answer_id = ?", setBest, aIDInt)
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
		"message": "Answer best status updated successfully",
	})
}
