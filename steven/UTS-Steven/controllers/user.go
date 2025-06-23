package controllers

import (
	"encoding/json"
	"net/http"
	"uts-gdg/database"
	"uts-gdg/models"
	"uts-gdg/utils"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	role := r.FormValue("role")

	var user models.User
	err := database.DB.QueryRow("SELECT email FROM users WHERE email = ? AND deleted_at IS NULL", email).
		Scan(&user.Email)
	if err == nil {
		http.Error(w, "Email used", http.StatusUnauthorized)
		return
	}

	if(email == "" || name == "" || password == "" || role == ""){
		http.Error(w, "email, name, password can not be empty", http.StatusInternalServerError)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO users (email, name, password, role) VALUES (?, ?, ?, ?)", email, name, hash, role)
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