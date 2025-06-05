package controllers

import (
	"encoding/json"
	"net/http"
	"resepku/database"
	"resepku/models"
	"resepku/utils"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email      string `json:"email_users"`
		Password   string `json:"pass_users"`
		NegaraAsal string `json:"negara_asal"`
		KodeNegara int    `json:"kode_negara"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var exists bool
	err = database.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM data_negara WHERE email_users = ?)",
		input.Email,
	).Scan(&exists)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	_, err = database.DB.Exec(`
    INSERT INTO data_negara (email_users, pass_users, negara_asal, kode_negara, role_users)
    VALUES (?, ?, ?, ?, 'user')`,
		input.Email, string(hashedPassword), input.NegaraAsal, input.KodeNegara)

	if err != nil {
		log.Println("Insert error:", err) // 🟢 Tambah log ini!
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email_users"`
		Password string `json:"pass_users"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var user models.Negara
	err = database.DB.QueryRow(
		"SELECT id, email_users, pass_users, role_users FROM data_negara WHERE email_users = ?",
		input.Email,
	).Scan(&user.ID, &user.Email, &user.Password, &user.Role)

	if err != nil {
		http.Error(w, "Email not found", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}

	token := utils.GenerateToken(32)
	_, err = database.DB.Exec("UPDATE data_negara SET token_users = ? WHERE id = ?", token, user.ID)

	if err != nil {
		http.Error(w, "Failed to update token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
	})
}
