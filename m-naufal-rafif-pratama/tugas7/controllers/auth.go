package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"tugas7/database"
	"tugas7/models"
	"tugas7/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string        `json:"token"`
	User  models.Student `json:"user"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		log.Printf("Login: Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Login: Attempting login for email: %s", loginReq.Email)

	var student models.Student
	err := database.DB.QueryRow(
		"SELECT id, name, email, password, major, year, org_id FROM students WHERE email = ?",
		loginReq.Email,
	).Scan(&student.ID, &student.Name, &student.Email, &student.Pass, &student.Major, &student.Year, &student.OrgID)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Login: No student found with email: %s", loginReq.Email)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			log.Printf("Login: Database error: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	if student.Pass != loginReq.Password {
		log.Printf("Login: Invalid password for email: %s", loginReq.Email)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	log.Printf("Login: Student authenticated successfully - ID: %d, Email: %s, OrgID: %d",
		student.ID, student.Email, student.OrgID)

	token, err := utils.GenerateToken(student.ID, student.Email, student.OrgID)
	if err != nil {
		log.Printf("Login: Error generating token: %v", err)
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	student.Pass = ""

	response := LoginResponse{
		Token: token,
		User:  student,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("Login: Successfully generated token for student ID: %d", student.ID)
}