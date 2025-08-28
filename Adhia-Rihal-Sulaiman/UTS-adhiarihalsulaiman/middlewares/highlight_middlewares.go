package middlewares

import (
	"database/sql"
	"net/http"
	"strconv"
	"uts_adhia/database"
	"uts_adhia/utils"
)

func OwnsHighlightMiddleware(next http.HandlerFunc, highlightIDParam string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxUserIDStr, ok := r.Context().Value(utils.UserIDKey).(string)
		if !ok || ctxUserIDStr == "" {
			http.Error(w, "Unauthorized: User ID not found in context", http.StatusUnauthorized)
			return
		}
		ctxUserID, err := strconv.Atoi(ctxUserIDStr)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid User ID in context", http.StatusInternalServerError)
			return
		}

		if highlightIDParam == "" {
			http.Error(w, "Highlight ID missing in path", http.StatusBadRequest)
			return
		}

		highlightID, err := strconv.Atoi(highlightIDParam)
		if err != nil {
			http.Error(w, "Invalid Highlight ID format", http.StatusBadRequest)
			return
		}

		var ownerID int
		err = database.DB.QueryRow("SELECT user_id FROM highlights WHERE highlight_id = ? AND deleted_at IS NULL", highlightID).Scan(&ownerID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Forbidden: Highlight not found or not accessible", http.StatusForbidden)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if ownerID != ctxUserID {
			return
		}

		http.Error(w, "Forbidden: You do not own this highlight, and lack admin or premium privileges", http.StatusForbidden)
	})
}

