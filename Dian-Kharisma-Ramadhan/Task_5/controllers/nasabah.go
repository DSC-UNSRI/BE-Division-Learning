package controllers

import (
	"Task_5/database"
	"Task_5/models"
	"Task_5/utils"
	"encoding/json"
	"net/http"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func RegisterNasabah(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama     string `json:"nama"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	token := utils.GenerateToken(32)

	result, err := database.DB.Exec("INSERT INTO nasabah (nama, password, token) VALUES (?, ?, ?)", input.Nama, hash, token)
	if err != nil {
		fmt.Println("Error insert nasabah:", err) // Print error ke console/server log
		http.Error(w, fmt.Sprintf("Failed to register nasabah: %v", err), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last insert ID:", err)
	}

	response := map[string]interface{}{
		"message": "Register nasabah success",
		"id":      id,
		"token":   token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func LoginNasabah(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nama := r.FormValue("nama")
	password := r.FormValue("password")

	var n models.Nasabah
	err := database.DB.QueryRow("SELECT id, nama, password, token FROM nasabah WHERE nama = ?", nama).
		Scan(&n.ID, &n.Nama, &n.Password, &n.Token)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(n.Password), []byte(password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Nasabah login",
		"token":   n.Token,
	})
}
