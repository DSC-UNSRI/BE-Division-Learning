package controllers

import (
	"UTS-Ahmad-Fadhil-Rizqi/database"
	"UTS-Ahmad-Fadhil-Rizqi/utils"
	"encoding/json"
	"net/http"
	"strings"
)

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(utils.UserIDKey).(int64)
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Format URL tidak valid", http.StatusBadRequest)
		return
	}
	questionID := parts[2] 
	
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Gagal mem-parsing form", http.StatusBadRequest)
		return
	}
	answerText := r.FormValue("answer")
	if answerText == "" {
		http.Error(w, "Isi jawaban tidak boleh kosong", http.StatusBadRequest)
		return
	}
	
	_, err := database.DB.Exec("INSERT INTO answers (user_id, question_id, answer) VALUES (?, ?, ?)",
		userID, questionID, answerText)
	if err != nil {
		
		http.Error(w, "Gagal membuat jawaban. Pastikan pertanyaan tersebut ada.", http.StatusInternalServerError)
		return
	}

	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Jawaban berhasil diposting"})
}


func UpdateAnswer(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(utils.UserIDKey).(int64)
	answerID := strings.TrimPrefix(r.URL.Path, "/answers/")
	
	var ownerID int64
	err := database.DB.QueryRow("SELECT user_id FROM answers WHERE id = ?", answerID).Scan(&ownerID)
	if err != nil {
		http.Error(w, "Jawaban tidak ditemukan", http.StatusNotFound)
		return
	}
	if ownerID != userID {
		http.Error(w, "Akses ditolak: Anda bukan pemilik jawaban ini", http.StatusForbidden)
		return
	}
	
	r.ParseForm()
	answerText := r.FormValue("answer")
	if answerText == "" {
		http.Error(w, "Isi jawaban tidak boleh kosong", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE answers SET answer = ? WHERE id = ?", answerText, answerID)
	if err != nil {
		http.Error(w, "Gagal memperbarui jawaban", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Jawaban berhasil diperbarui"})
}

func DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(utils.UserIDKey).(int64)
	answerID := strings.TrimPrefix(r.URL.Path, "/answers/")

	var ownerID int64
	err := database.DB.QueryRow("SELECT user_id FROM answers WHERE id = ?", answerID).Scan(&ownerID)
	if err != nil {
		http.Error(w, "Jawaban tidak ditemukan", http.StatusNotFound)
		return
	}
	if ownerID != userID {
		http.Error(w, "Akses ditolak: Anda bukan pemilik jawaban ini", http.StatusForbidden)
		return
	}

	_, err = database.DB.Exec("UPDATE answers SET deleted_at = NOW() WHERE id = ?", answerID)
	if err != nil {
		http.Error(w, "Gagal menghapus jawaban", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Jawaban berhasil dihapus"})
}