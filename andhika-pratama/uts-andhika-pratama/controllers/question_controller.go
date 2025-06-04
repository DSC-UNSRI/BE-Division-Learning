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

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	ctxUserID := utils.Atoi(r.Context().Value(utils.UserIDKey).(string))

	err := r.ParseForm()
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
		UserID:  ctxUserID,
		Title:   title,
		Content: content,
	}

	res, err := database.DB.Exec(`
		INSERT INTO questions (user_id, title, content)
		VALUES (?, ?, ?)`,
		question.UserID, question.Title, question.Content)

	if err != nil {
		http.Error(w, "Failed to create question", http.StatusInternalServerError)
		return
	}

	newQuestionID, err := res.LastInsertId()
		if err != nil {
		http.Error(w, "Question created, but failed to retrieve QuestionID", http.StatusInternalServerError)
		return
	}
	question.QuestionID = int(newQuestionID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Question created successfully",
		"question":  question,
	})
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT question_id, user_id, title, content, upvotes, downvotes FROM questions WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	questions := []models.Question{}
	for rows.Next() {
		question := models.Question{}
		rows.Scan(&question.QuestionID, &question.UserID, &question.Title, &question.Content, &question.Upvotes, &question.Downvotes)
		questions = append(questions, question)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"questions": questions,
	})
}

func GetQuestionByID(w http.ResponseWriter, r *http.Request, questionID string) {
	if questionID == "" {
		http.Error(w, "Please input question_id in the url", http.StatusBadRequest)
		return
	}

	question := models.Question{}
	err := database.DB.QueryRow("SELECT question_id, user_id, title, content, upvotes, downvotes FROM questions WHERE question_id = ? AND deleted_at IS NULL", questionID).
		Scan(&question.QuestionID, &question.UserID, &question.Title, &question.Content, &question.Upvotes, &question.Downvotes)

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
		http.Error(w, "Please input question_id in the url", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
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

	query := "UPDATE questions SET " + strings.Join(updateFields, ", ") + " WHERE question_id = ? AND deleted_at IS NULL"
	updateValues = append(updateValues, utils.Atoi(questionID))

	_, err = database.DB.Exec(query, updateValues...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	question := models.Question{}
	err = database.DB.QueryRow("SELECT question_id, user_id, title, content, upvotes, downvotes FROM questions WHERE question_id = ? AND deleted_at IS NULL", questionID).
		Scan(&question.QuestionID, &question.UserID, &question.Title, &question.Content, &question.Upvotes, &question.Downvotes)

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
		http.Error(w, "Please input question_id in the url", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("UPDATE questions SET deleted_at = NOW() WHERE question_id = ?", utils.Atoi(questionID))
	if err != nil {
		http.Error(w, "Failed to delete question", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("UPDATE answers SET deleted_at = NOW() WHERE question_id = ?", utils.Atoi(questionID))
	if err != nil {
		http.Error(w, "Failed to delete answers for this question", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Question deleted successfully",
	})
}

func UpvoteQuestion(w http.ResponseWriter, r *http.Request, questionID string) {
	if questionID == "" {
		http.Error(w, "Please input question_id in the url", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("UPDATE questions SET upvotes = upvotes + 1 WHERE question_id = ?", utils.Atoi(questionID))

	if err != nil {
		http.Error(w, "Failed to upvote this question", http.StatusInternalServerError)
		return
	}
	
	question := models.Question{}
	err = database.DB.QueryRow("SELECT question_id, user_id, title, content, upvotes, downvotes FROM questions WHERE question_id = ? AND deleted_at IS NULL", questionID).
		Scan(&question.QuestionID, &question.UserID, &question.Title, &question.Content, &question.Upvotes, &question.Downvotes)

	if err != nil {
		http.Error(w, "Failed to fetch upvoted question", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Question upvoted successfully",
		"question": question,
	})
}

func DownvoteQuestion(w http.ResponseWriter, r *http.Request, questionID string) {
	if questionID == "" {
		http.Error(w, "Please input question_id in the url", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("UPDATE questions SET downvotes = downvotes + 1 WHERE question_id = ?", utils.Atoi(questionID))
	
	if err != nil {
		http.Error(w, "Failed to downvote this question", http.StatusInternalServerError)
		return
	}
	
	question := models.Question{}
	err = database.DB.QueryRow("SELECT question_id, user_id, title, content, upvotes, downvotes FROM questions WHERE question_id = ? AND deleted_at IS NULL", questionID).
		Scan(&question.QuestionID, &question.UserID, &question.Title, &question.Content, &question.Upvotes, &question.Downvotes)

	if err != nil {
		http.Error(w, "Failed to fetch downvoted question", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Question downvoted successfully",
		"question": question,
	})
}