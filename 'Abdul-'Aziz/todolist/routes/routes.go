package routes

import (
	"database/sql"
	"net/http"
	"tugas/todolist/controllers"
	middleware "tugas/todolist/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()


	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateUser(db, w, r)
	}).Methods("POST")
	userRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllUsers(db, w, r)
	}).Methods("GET")
	userRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetUserByID(db, w, r)
	}).Methods("GET")
	userRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateUser(db, w, r)
	}).Methods("PUT")
	userRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUser(db, w, r)
	}).Methods("DELETE")



	taskRouter := router.PathPrefix("/tasks").Subrouter()
	
	taskRouter.Handle("", middleware.AuthMiddleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	controllers.CreateTask(db, w, r)
	}))).Methods("POST")
	taskRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllTasks(db, w, r)
	}).Methods("GET")
	taskRouter.Handle("/all-task-user", middleware.AuthMiddleware()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	controllers.GetAllTasksWithUser(db, w, r)
	}))).Methods("GET")
	taskRouter.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetTaskByUserID(db, w, r)
	}).Methods("GET")
	taskRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetTaskByOnlyID(db, w, r)
	}).Methods("GET")
	taskRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateTask(db, w, r)
	}).Methods("PUT")
	taskRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteTask(db, w, r)
	}).Methods("DELETE")

	


	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.Login(db, w, r)
	}).Methods("POST")

	return router
}
