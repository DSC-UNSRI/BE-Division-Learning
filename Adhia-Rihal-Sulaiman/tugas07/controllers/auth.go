package controllers

import (
	"be_pert7/database"
	"be_pert7/models"
	"be_pert7/utils"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Name := r.FormValue("chef_name")
	Password := r.FormValue("password")

	if Name == "" || Password == "" {
		http.Error(w, "Missing required fields: Name, Password", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM chefs WHERE chef_name = ? AND deleted_at IS NULL)", Name).Scan(&exists)
	if err != nil {
		fmt.Println("Error saat memeriksa keberadaan Nama Chef:", err)
		http.Error(w, "Database error checking Chef Name existence", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Chef with Name already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec("INSERT INTO chefs (chef_name, password) VALUES (?, ?)",
		Name, hashedPassword,
	)
	if err != nil {
		http.Error(w, "Failed to register Chef", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Chef registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Name := r.FormValue("chef_name")
	Password := r.FormValue("password")

	if Name == "" || Password == "" {
		http.Error(w, "Missing required fields: Name, Password", http.StatusBadRequest)
		return
	}

	var chef models.Chef
	err := database.DB.QueryRow("SELECT chef_id, chef_name, password, token, role FROM chefs WHERE chef_name = ? AND deleted_at IS NULL", Name).
		Scan(&chef.ChefID, &chef.Name, &chef.Password, &chef.Token, &chef.Role)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(chef.Password), []byte(Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	newToken := utils.GenerateToken(32)

	_, err = database.DB.Exec("UPDATE chefs SET token = ? WHERE chef_name = ? AND deleted_at IS NULL", newToken, Name)
	if err != nil {
		http.Error(w, "Failed to update token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": newToken,
	})
}
