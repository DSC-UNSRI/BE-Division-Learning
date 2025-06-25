package controllers

import (
	"UTS-Ahmad-Fadhil-Rizqi/database"
	"UTS-Ahmad-Fadhil-Rizqi/models"
	"UTS-Ahmad-Fadhil-Rizqi/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Gagal mem-parsing form", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	
	if username == "" || password == "" {
		http.Error(w, "Username dan password wajib diisi", http.StatusBadRequest)
		return
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Gagal melakukan hash pada password", http.StatusInternalServerError)
		return
	}
	hashedPassword := string(hashedPasswordBytes)

	_, err = database.DB.Exec("INSERT INTO users (username, password_hash, tier) VALUES (?, ?, ?)",
		username, hashedPassword, "free")
	if err != nil {
		http.Error(w, "Gagal membuat pengguna. Username mungkin sudah ada.", http.StatusInternalServerError)
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Pengguna berhasil terdaftar"})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Gagal mem-parsing form", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, "Username dan password wajib diisi", http.StatusBadRequest)
		return
	}

	var user models.User
	err := database.DB.QueryRow("SELECT id, username, password_hash, tier FROM users WHERE username = ? AND deleted_at IS NULL", username).
		Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Tier)
	if err != nil {
		if err == sql.ErrNoRows { http.Error(w, "Username atau password salah", http.StatusUnauthorized)
		} else { http.Error(w, "Terjadi kesalahan internal", http.StatusInternalServerError) }
		return
	}

	
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		http.Error(w, "Username atau password salah", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateBearerToken(32)
	if err != nil {
		http.Error(w, "Gagal membuat token", http.StatusInternalServerError)
		return
	}

	expiresAt := time.Now().Add(5 * time.Minute) 

	_, err = database.DB.Exec("UPDATE users SET token = ?, token_expires_at = ? WHERE id = ?", token, expiresAt, user.ID)
	if err != nil {
		http.Error(w, "Gagal menyimpan sesi login", http.StatusInternalServerError)
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login berhasil",
		"token":   token,
	})
}