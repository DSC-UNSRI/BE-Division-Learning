package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas/todolist/models"
	"tugas/todolist/usecase"

	"github.com/gorilla/mux"
)

func CreateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := usecase.CreateUser(db, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}

func GetUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := usecase.GetUserByID(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := usecase.UpdateUser(db, id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated"})
}

func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := usecase.DeleteUser(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
}

func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&loginData)

	user, err := usecase.LoginUser(db, loginData.Name, loginData.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(user)
}
