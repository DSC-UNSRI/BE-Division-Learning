package controllers

import (
	"UTS-Ahmad-Fadhil-Rizqi/database"
	"UTS-Ahmad-Fadhil-Rizqi/models"
	"UTS-Ahmad-Fadhil-Rizqi/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(utils.UserIDKey).(int64)
	userTier, _ := r.Context().Value(utils.UserTierKey).(string)
	
	if userTier == "free" {

		var lastQuestionTime sql.NullTime
		
		err := database.DB.QueryRow("SELECT MAX(created_at) FROM questions WHERE user_id = ?", userID).Scan(&lastQuestionTime)
		if err != nil && err != sql.ErrNoRows {
			http.Error(w, "Gagal memeriksa kuota pertanyaan", http.StatusInternalServerError)
			return
		}

		if lastQuestionTime.Valid {
			secondsSinceLastPost := time.Since(lastQuestionTime.Time).Seconds()
			if secondsSinceLastPost < 120 {
				remainingTime := 120 - int(secondsSinceLastPost)
				errorMessage := fmt.Sprintf("Anda harus menunggu %d detik lagi. Ingin bisa bertanya sepuasnya tanpa harus menunggu? Cobalah fitur premium!", remainingTime)
				http.Error(w, errorMessage, http.StatusTooManyRequests)
				return 
			}
		}
	}
	
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Gagal mem-parsing form", http.StatusBadRequest)
		return
	}
	
	questionText := r.FormValue("question")
	
	if questionText == "" {
		http.Error(w, "Isi pertanyaan tidak boleh kosong", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("INSERT INTO questions (user_id, question) VALUES (?, ?)",
		userID, questionText)
	if err != nil {
		http.Error(w, "Gagal membuat pertanyaan", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Pertanyaan berhasil dibuat"})
}


func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, user_id, question, created_at FROM questions WHERE deleted_at IS NULL ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, "Gagal mengambil data pertanyaan", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var questions []models.Question
	
	for rows.Next() {
		var q models.Question
		
		if err := rows.Scan(&q.ID, &q.UserID, &q.Question, &q.CreatedAt); err != nil {
			http.Error(w, "Gagal memindai data pertanyaan", http.StatusInternalServerError)
			return
		}
		
		questions = append(questions, q)
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]models.Question{"questions": questions})
}



func GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/questions/")
	
	var question models.Question
	
	err := database.DB.QueryRow("SELECT id, user_id, question, created_at, updated_at FROM questions WHERE id = ? AND deleted_at IS NULL", id).
		Scan(&question.ID, &question.UserID, &question.Question, &question.CreatedAt, &question.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Pertanyaan tidak ditemukan", http.StatusNotFound)
		} else {
			http.Error(w, "Gagal mengambil pertanyaan", http.StatusInternalServerError)
		}
		return
	}

	rows, err := database.DB.Query("SELECT id, user_id, answer, created_at, updated_at FROM answers WHERE question_id = ? AND deleted_at IS NULL", id)
	if err != nil {
		http.Error(w, "Gagal mengambil jawaban", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var answers []models.Answer
	for rows.Next() {
		var a models.Answer
		if err := rows.Scan(&a.ID, &a.UserID, &a.Answer, &a.CreatedAt, &a.UpdatedAt); err != nil {
			http.Error(w, "Gagal memindai data jawaban", http.StatusInternalServerError)
			return
		}
		answers = append(answers, a)
	}
	
	response := map[string]interface{}{
		"question": question,
		"answers":  answers,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}



func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(utils.UserIDKey).(int64)
	userTier, _ := r.Context().Value(utils.UserTierKey).(string)
	
	if userTier == "free" {
		
		errorMessage := "Ingin mengubah pertanyaan kamu? Upgrade ke premium sekarang untuk bisa mengubah pertanyaan kamu!"
		http.Error(w, errorMessage, http.StatusForbidden) 
		return 
	}
	
	id := strings.TrimPrefix(r.URL.Path, "/questions/")

	var ownerID int64
	err := database.DB.QueryRow("SELECT user_id FROM questions WHERE id = ?", id).Scan(&ownerID)
	if err != nil {
		http.Error(w, "Pertanyaan tidak ditemukan", http.StatusNotFound)
		return
	}
	if ownerID != userID {
		http.Error(w, "Akses ditolak: Anda bukan pemilik pertanyaan ini", http.StatusForbidden)
		return
	}

	r.ParseForm()
	questionText := r.FormValue("question")
	if questionText == "" {
		http.Error(w, "Isi pertanyaan tidak boleh kosong", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE questions SET question = ? WHERE id = ?", questionText, id)
	if err != nil {
		http.Error(w, "Gagal memperbarui pertanyaan", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Pertanyaan berhasil diperbarui"})
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(utils.UserIDKey).(int64)
	id := strings.TrimPrefix(r.URL.Path, "/questions/")

	var ownerID int64
	err := database.DB.QueryRow("SELECT user_id FROM questions WHERE id = ?", id).Scan(&ownerID)
	if err != nil {
		http.Error(w, "Pertanyaan tidak ditemukan", http.StatusNotFound)
		return
	}
	if ownerID != userID {
		http.Error(w, "Akses ditolak: Anda bukan pemilik pertanyaan ini", http.StatusForbidden)
		return
	}

	_, err = database.DB.Exec("UPDATE questions SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Gagal menghapus pertanyaan", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Pertanyaan berhasil dihapus"})
}