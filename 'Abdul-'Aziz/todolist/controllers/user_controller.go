package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"tugas/todolist/dto"
	"tugas/todolist/helper"
	"tugas/todolist/models"
	"tugas/todolist/usecase"

	"github.com/gorilla/mux"
)

func CreateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := usecase.CreateUser(db, user)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}

func GetAllUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	users, err := usecase.GetAllUsers(db)
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(users)
}


func GetUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	fmt.Println("user id", id)

	user, err := usecase.GetUserByID(db, id)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println("user", user)

	err := usecase.UpdateUser(db, id, user)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated"})
}

func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := usecase.DeleteUser(db, id)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
}

func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var loginData dto.LoginData

	json.NewDecoder(r.Body).Decode(&loginData)

	user, err := usecase.LoginUser(db, loginData.Email, loginData.Password)

	if err != nil {
		helper.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}


	json.NewEncoder(w).Encode(user)
}
