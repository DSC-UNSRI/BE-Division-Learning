package middlewares

import (
	"uts/database"
	"uts/utils"

	"net/http"
	"database/sql"
)

func OwnsQuestionMiddleware(next http.Handler, questionID string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxUserID := r.Context().Value(utils.UserIDKey)
		userID, ok := ctxUserID.(string)
		if !ok || userID == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if questionID == "" {
			http.Error(w, "question_id is null", http.StatusBadRequest)
			return
		}

		var questionOwnerID string
		err := database.DB.QueryRow("SELECT user_id FROM questions WHERE question_id = ? AND deleted_at IS NULL", questionID).Scan(&questionOwnerID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Question not found", http.StatusNotFound)
			} else {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}


		if questionOwnerID != userID {
			http.Error(w, "Forbidden - the question isn't yours", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}