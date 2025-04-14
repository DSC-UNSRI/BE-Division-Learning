package routes

import (
	"database/sql"
	"net/http"
	"tugas/todolist/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateUser(db, w, r)
	}).Methods("POST")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllUsers(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUserByID(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateUser(db, w, r)
	}).Methods("PUT")

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUser(db, w, r)
	}).Methods("DELETE")

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.Login(db, w, r)
	}).Methods("POST")

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateTask(db, w, r)
	}).Methods("POST")

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllTasks(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetTaskByID(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateTask(db, w, r)
	}).Methods("PUT")

	router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteTask(db, w, r)
	}).Methods("DELETE")

	return router
}
