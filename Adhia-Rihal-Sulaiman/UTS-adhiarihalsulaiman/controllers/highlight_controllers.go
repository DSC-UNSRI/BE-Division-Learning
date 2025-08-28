package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"uts_adhia/database"
	"uts_adhia/models"
	"uts_adhia/utils"
)

func CreateHighlight(w http.ResponseWriter, r *http.Request) {
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

	contentType := r.FormValue("content_type")
	contentIDStr := r.FormValue("content_id")

	if contentType == "" || contentIDStr == "" {
		http.Error(w, "Missing required fields: content_type, content_id", http.StatusBadRequest)
		return
	}

	contentID, err := strconv.Atoi(contentIDStr)
	if err != nil {
		http.Error(w, "Invalid Content ID format", http.StatusBadRequest)
		return
	}

	var contentExists bool
	switch contentType {
	case "question":
		err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM questions WHERE question_id = ? AND deleted_at IS NULL)", contentID).Scan(&contentExists)
	case "answer":
		err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM answers WHERE answer_id = ? AND deleted_at IS NULL)", contentID).Scan(&contentExists)
	default:
		http.Error(w, "Invalid content_type. Must be 'question' or 'answer'", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !contentExists {
		http.Error(w, "Target content not found", http.StatusNotFound)
		return
	}

	highlight := models.Highlight{
		UserID:      ctxUserID,
		ContentType: contentType,
		ContentID:   contentID,
		IsActive:    true,
	}

	res, err := database.DB.Exec(`
		INSERT INTO highlights (user_id, content_type, content_id, is_active)
		VALUES (?, ?, ?, ?)`,
		highlight.UserID, highlight.ContentType, highlight.ContentID, highlight.IsActive)

	if err != nil {
		http.Error(w, "Failed to create highlight", http.StatusInternalServerError)
		return
	}

	newHighlightID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Highlight created, but failed to retrieve Highlight ID", http.StatusInternalServerError)
		return
	}
	highlight.ID = int(newHighlightID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Highlight created successfully",
		"highlight": highlight,
	})
}

func GetAllHighlights(w http.ResponseWriter, r *http.Request) {
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

	userRole, okRole := r.Context().Value(utils.UserRoleKey).(string)
	if !okRole {
		http.Error(w, "Unauthorized: User role not found in context", http.StatusInternalServerError)
		return
	}

	query := "SELECT highlight_id, user_id, content_type, content_id, is_active, created_at, expires_at, deleted_at FROM highlights WHERE deleted_at IS NULL AND is_active = TRUE"
	queryArgs := []interface{}{}

	if userRole != "admin" {
		query += " AND user_id = ?"
		queryArgs = append(queryArgs, ctxUserID)
	}

	rows, err := database.DB.Query(query, queryArgs...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	highlights := []models.Highlight{}
	for rows.Next() {
		highlight := models.Highlight{}
		err := rows.Scan(&highlight.ID, &highlight.UserID, &highlight.ContentType, &highlight.ContentID, &highlight.IsActive, &highlight.CreatedAt, &highlight.ExpiresAt, &highlight.DeletedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		highlights = append(highlights, highlight)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"highlights": highlights,
	})
}

func GetHighlightByID(w http.ResponseWriter, r *http.Request, highlightID string) {
	if highlightID == "" {
		http.Error(w, "Highlight ID missing in URL", http.StatusBadRequest)
		return
	}

	hIDInt, err := strconv.Atoi(highlightID)
	if err != nil {
		http.Error(w, "Invalid Highlight ID format", http.StatusBadRequest)
		return
	}

	ctxUserIDStr, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok || ctxUserIDStr == "" {
		http.Error(w, "Unauthorized: User ID not found in context", http.StatusInternalServerError)
		return
	}
	ctxUserID, err := strconv.Atoi(ctxUserIDStr)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid User ID in context", http.StatusInternalServerError)
		return
	}

	userRole, okRole := r.Context().Value(utils.UserRoleKey).(string)
	if !okRole {
		http.Error(w, "Unauthorized: User role not found in context", http.StatusInternalServerError)
		return
	}

	highlight := models.Highlight{}
	query := "SELECT highlight_id, user_id, content_type, content_id, is_active, created_at, expires_at, deleted_at FROM highlights WHERE highlight_id = ? AND deleted_at IS NULL AND is_active = TRUE"
	queryArgs := []interface{}{hIDInt}

	if userRole != "admin" {
		query += " AND user_id = ?"
		queryArgs = append(queryArgs, ctxUserID)
	}

	err = database.DB.QueryRow(query, queryArgs...).
		Scan(&highlight.ID, &highlight.UserID, &highlight.ContentType, &highlight.ContentID,
			&highlight.IsActive, &highlight.CreatedAt, &highlight.ExpiresAt, &highlight.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Highlight not found or not accessible", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(highlight)
}

func DeleteHighlight(w http.ResponseWriter, r *http.Request, highlightID string) {
	if highlightID == "" {
		http.Error(w, "Highlight ID missing in URL", http.StatusBadRequest)
		return
	}

	hIDInt, err := strconv.Atoi(highlightID)
	if err != nil {
		http.Error(w, "Invalid Highlight ID format", http.StatusBadRequest)
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

	userRole, okRole := r.Context().Value(utils.UserRoleKey).(string)
	if !okRole {
		http.Error(w, "Unauthorized: User role not found in context", http.StatusInternalServerError)
		return
	}

	query := "UPDATE highlights SET deleted_at = CURRENT_TIMESTAMP, is_active = FALSE WHERE highlight_id = ?"
	queryArgs := []interface{}{hIDInt}

	if userRole != "admin" {
		query += " AND user_id = ?"
		queryArgs = append(queryArgs, ctxUserID)
	}

	res, err := database.DB.Exec(query, queryArgs...)
	if err != nil {
		http.Error(w, "Failed to delete highlight", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to check rows affected after delete", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Highlight not found or you do not have permission to delete it", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Highlight deleted successfully",
	})
}