package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas/todolist/models"
	"tugas/todolist/repository"

	"github.com/gorilla/mux"
)

func CreateTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	err := repository.InsertTask(db, task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task created"})
}

func GetAllTasks(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.SelectAllTasks(db)
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, err := repository.SelectTaskByID(db, id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func UpdateTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	err := repository.UpdateTaskByID(db, id, task)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated"})
}

func DeleteTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := repository.DeleteTaskByID(db, id)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
}
