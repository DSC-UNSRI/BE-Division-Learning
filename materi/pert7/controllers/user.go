package controllers

import (
	"be_pert5/database"
	"be_pert5/models"
	"be_pert5/utils"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, hash)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	var user models.User
	err := database.DB.QueryRow("SELECT id, email, password, token FROM users WHERE email = ? AND deleted_at IS NULL", email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Token)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := utils.GenerateToken(32)

	_, err = database.DB.Exec("UPDATE users SET token = ? WHERE email = ? AND deleted_at IS NULL", token, email)
	if err != nil {
		http.Error(w, "Failed to set token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
