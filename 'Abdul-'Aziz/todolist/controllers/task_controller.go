package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas/todolist/helper"
	"tugas/todolist/models"
	"tugas/todolist/usecase"

	"github.com/gorilla/mux"
)

func CreateTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	err := usecase.CreateTask(db, task)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task created"})
}

func GetAllTasksWithUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	
	tasks, err := usecase.GetAllTasksWithUser(db) 
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func GetAllTasks(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	
	tasks, err := usecase.GetAllTasks(db) 
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskByOnlyID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, err := usecase.GetTaskByOnlyId(db, id)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(task)
}

func GetTaskByUserID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, err := usecase.GetTaskByUserID(db, id)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(task)
}

func UpdateTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	err := usecase.UpdateTask(db, id, task)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated"})
}

func DeleteTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := usecase.DeleteTask(db, id)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
}
