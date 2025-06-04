package controllers

import (
	"uts/database"
	"uts/models"
	"uts/utils"

	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

func CreateAnswer(w http.ResponseWriter, r *http.Request, questionID string) {
	if questionID == "" {
		http.Error(w, "Please input question_id in the url", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM questions WHERE question_id = ? AND deleted_at IS NULL)", questionID).Scan(&exists)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if !exists {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}

	ctxUserID := utils.Atoi(r.Context().Value(utils.UserIDKey).(string))
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
		UserID:  ctxUserID,
		QuestionID: utils.Atoi(questionID),
		Content: content,
	}

	res, err := database.DB.Exec(`
		INSERT INTO answers (question_id, user_id, content)
		VALUES (?, ?, ?)`,
		answer.QuestionID, answer.UserID, answer.Content)

	if err != nil {
		http.Error(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}

	newAnswerID, err := res.LastInsertId()
		if err != nil {
		http.Error(w, "Answer created, but failed to retrieve AnswernID", http.StatusInternalServerError)
		return
	}
	answer.AnswerID = int(newAnswerID)

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

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM questions WHERE question_id = ? AND deleted_at IS NULL)", questionID).Scan(&exists)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if !exists {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}

	rows, err := database.DB.Query("SELECT answer_id, question_id, user_id, content, upvotes, downvotes FROM answers WHERE question_id = ? AND deleted_at IS NULL", questionID)
	
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	answers := []models.Answer{}
	for rows.Next() {
		answer := models.Answer{}
		rows.Scan(&answer.AnswerID, &answer.QuestionID, &answer.UserID, &answer.Content, &answer.Upvotes, &answer.Downvotes)
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

	answer := models.Answer{}
	err := database.DB.QueryRow("SELECT answer_id, question_id, user_id, content, upvotes, downvotes FROM answers WHERE answer_id = ? AND deleted_at IS NULL", answerID).
		Scan(&answer.AnswerID, &answer.QuestionID, &answer.UserID, &answer.Content, &answer.Upvotes, &answer.Downvotes)

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

	err := r.ParseForm()
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

	query := "UPDATE answers SET " + strings.Join(updateFields, ", ") + " WHERE answer_id = ? AND deleted_at IS NULL"
	updateValues = append(updateValues, utils.Atoi(answerID))

	_, err = database.DB.Exec(query, updateValues...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	answer := models.Answer{}
	err = database.DB.QueryRow("SELECT answer_id, question_id, user_id, content, upvotes, downvotes FROM answers WHERE answer_id = ? AND deleted_at IS NULL", answerID).
		Scan(&answer.AnswerID, &answer.QuestionID, &answer.UserID, &answer.Content, &answer.Upvotes, &answer.Downvotes)

	if err != nil {
		http.Error(w, "Failed to fetch updated Answer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Answer updated successfully",
		"answer": answer,
	})
}

func DeleteAnswer(w http.ResponseWriter, r *http.Request, answerID string) {
	if answerID == "" {
		http.Error(w, "Please input answer_id in the url", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("UPDATE answers SET deleted_at = NOW() WHERE answer_id = ?", answerID)
	if err != nil {
		http.Error(w, "Failed to delete answer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Answer deleted successfully",
	})
}

func UpvoteAnswer(w http.ResponseWriter, r *http.Request, answerID string) {
	if answerID == "" {
		http.Error(w, "Please input answer_id in the url", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM answers WHERE answer_id = ? AND deleted_at IS NULL)", answerID).Scan(&exists)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if !exists {
		http.Error(w, "Answer not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE answers SET upvotes = upvotes + 1 WHERE answer_id = ? AND deleted_at IS NULL", answerID)

	if err != nil {
		http.Error(w, "Failed to upvote this answer", http.StatusInternalServerError)
		return
	}
	
	answer := models.Answer{}
	err = database.DB.QueryRow("SELECT answer_id, question_id, user_id, content, upvotes, downvotes FROM answers WHERE answer_id = ? AND deleted_at IS NULL", answerID).
		Scan(&answer.AnswerID, &answer.QuestionID, &answer.UserID,  &answer.Content, &answer.Upvotes, &answer.Downvotes)

	if err != nil {
		http.Error(w, "Failed to fetch upvoted answer", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Answer upvoted successfully",
		"answer": answer,
	})
}

func DownvoteAnswer(w http.ResponseWriter, r *http.Request, answerID string) {
	if answerID == "" {
		http.Error(w, "Please input answer_id in the url", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM answers WHERE answer_id = ? AND deleted_at IS NULL)", answerID).Scan(&exists)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if !exists {
		http.Error(w, "Answer not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE answers SET downvotes = downvotes + 1 WHERE answer_id = ? AND deleted_at IS NULL", answerID)

	if err != nil {
		http.Error(w, "Failed to downvote this answer", http.StatusInternalServerError)
		return
	}
	
	answer := models.Answer{}
	err = database.DB.QueryRow("SELECT answer_id, question_id, user_id, content, upvotes, downvotes FROM answers WHERE answer_id = ? AND deleted_at IS NULL", answerID).
		Scan(&answer.AnswerID, &answer.QuestionID, &answer.UserID,  &answer.Content, &answer.Upvotes, &answer.Downvotes)

	if err != nil {
		http.Error(w, "Failed to fetch downvoted answer", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Answer downvoted successfully",
		"answer": answer,
	})
}