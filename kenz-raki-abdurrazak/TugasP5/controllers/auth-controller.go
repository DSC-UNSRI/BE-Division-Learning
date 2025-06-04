package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/config"
	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/models"
)

type LoginRequest struct {
	Name    string `json:"name"`
	AuthKey string `json:"auth_key"`
}

type LoginResponse struct {
	Speaker  models.Speaker `json:"speaker"`
	AuthKey  string         `json:"auth_key"`
	Message  string         `json:"message"`
}

func LoginSpeaker(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.Name == "" || req.AuthKey == "" {
		http.Error(w, "Name and auth_key are required", http.StatusBadRequest)
		return
	}

	// Check if speaker exists with the provided credentials
	var speaker models.Speaker
	err := config.DB.QueryRow(
		"SELECT id, name, expertise, auth_key FROM speakers WHERE name = ? AND auth_key = ?",
		req.Name,
		req.AuthKey,
	).Scan(&speaker.ID, &speaker.Name, &speaker.Expertise, &speaker.AuthKey)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Return response with speaker data and authentication token
	resp := LoginResponse{
		Speaker:  speaker,
		AuthKey:  speaker.AuthKey,
		Message:  "Login successful",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
