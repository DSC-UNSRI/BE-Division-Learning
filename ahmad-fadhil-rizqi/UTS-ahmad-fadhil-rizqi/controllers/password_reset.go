package controllers

import (
	"UTS-Ahmad-Fadhil-Rizqi/database"
	"UTS-Ahmad-Fadhil-Rizqi/utils"
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SetSecurityQuestion(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(utils.UserIDKey).(int64)

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Gagal mem-parsing form", http.StatusBadRequest)
		return
	}
	question := r.FormValue("question")
	answer := r.FormValue("answer")
	if question == "" || answer == "" {
		http.Error(w, "Pertanyaan dan jawaban wajib diisi", http.StatusBadRequest)
		return
	}

	hashedAnswerBytes, err := bcrypt.GenerateFromPassword([]byte(answer), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Gagal melakukan hash pada jawaban", http.StatusInternalServerError)
		return
	}
	hashedAnswer := string(hashedAnswerBytes)

	_, err = database.DB.Exec(`
		INSERT INTO security_questions (user_id, question, hashed_answer)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE question = VALUES(question), hashed_answer = VALUES(hashed_answer)
	`, userID, question, hashedAnswer)

	if err != nil {
		http.Error(w, "Gagal mengatur pertanyaan keamanan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Pertanyaan keamanan berhasil diatur"})
}

func GetSecurityQuestionForUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Gagal mem-parsing form", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Username wajib diisi", http.StatusBadRequest)
		return
	}

	var question string
	err := database.DB.QueryRow(`
		SELECT sq.question FROM security_questions sq
		JOIN users u ON sq.user_id = u.id
		WHERE u.username = ?
	`, username).Scan(&question)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Pengguna tidak ditemukan atau belum mengatur pertanyaan keamanan", http.StatusNotFound)
			return
		}
		http.Error(w, "Gagal mengambil pertanyaan keamanan", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"question": question})
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Gagal mem-parsing form", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	answer := r.FormValue("answer")
	newPassword := r.FormValue("new_password")

	
	if username == "" || answer == "" || newPassword == "" {
		http.Error(w, "Username, jawaban, dan password baru wajib diisi", http.StatusBadRequest)
		return
	}

	var userID int64
	var hashedAnswer string
	err := database.DB.QueryRow(`
		SELECT u.id, sq.hashed_answer FROM users u
		JOIN security_questions sq ON u.id = sq.user_id
		WHERE u.username = ?
	`, username).Scan(&userID, &hashedAnswer)

	if err != nil {
		http.Error(w, "Pengguna tidak ditemukan", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedAnswer), []byte(answer))
	if err != nil {
		http.Error(w, "Jawaban untuk pertanyaan keamanan salah", http.StatusForbidden)
		return
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Gagal melakukan hash pada password baru", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("UPDATE users SET password_hash = ? WHERE id = ?", newHashedPassword, userID)
	if err != nil {
		http.Error(w, "Gagal memperbarui password", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Password telah berhasil direset"})
}