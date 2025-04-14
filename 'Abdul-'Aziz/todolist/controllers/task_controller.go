package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas/todolist/models"
	"tugas/todolist/usecase"

	"github.com/gorilla/mux"
)

func CreateTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	err := usecase.CreateTask(db, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task created"})
}

func GetAllTasks(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// ini bisa tetap pakai repository langsung kalau gak ada logic khusus
	tasks, err := usecase.GetAllTasks(db) // bisa dibungkus juga kalau mau
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, err := usecase.GetTaskByID(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated"})
}

func DeleteTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := usecase.DeleteTask(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
}
