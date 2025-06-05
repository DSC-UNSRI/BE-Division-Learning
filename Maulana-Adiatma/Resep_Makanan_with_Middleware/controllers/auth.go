package controllers

import (
	"encoding/json"
	"net/http"
	"resepku/database"
	"resepku/utils"
)

type RegisterRequest struct {
	NamaNegara  string `json:"nama_negara"`
	KodeNegara  int    `json:"kode_negara"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateToken(32)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO negara (nama_negara, kode_negara, auth_key) VALUES (?, ?, ?)", req.NamaNegara, req.KodeNegara, token)
	if err != nil {
		http.Error(w, "Failed to register", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(AuthResponse{Token: token})
}
