package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas/todolist/models"
	"tugas/todolist/repository"

	"github.com/gorilla/mux"
)

func CreateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Name == "" || user.Password == "" {
		http.Error(w, "Name and password are required", http.StatusBadRequest)
		return
	}

	err := repository.InsertUser(db, user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}

func GetAllUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	users, err := repository.SelectAllUsers(db)
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := repository.SelectUserByID(db, id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := repository.UpdateUserByID(db, id, user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated"})
}

func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := repository.DeleteUserByID(db, id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
}

func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	user, err := repository.ValidateLogin(db, input.Name, input.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "user_id": string(rune(user.ID))})
}
