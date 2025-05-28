package routes

import (
	"pertemuan05/controllers"
	"pertemuan05/middlewares"
	"pertemuan05/utils"

	"net/http"
)

func withAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ApplyMiddlewares(handler, middleware.AuthMiddleware).ServeHTTP(w, r)
	}
}

func withAdminAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ApplyMiddlewares(handler, middleware.AuthMiddleware, middleware.AdminMiddleware).ServeHTTP(w, r)
	}
}

func RoutesHandlers() {
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/lecturers", lecturersHandler)
	http.HandleFunc("/lecturers/", lecturersHandlerWithID)
	http.HandleFunc("/courses", coursesHandler)
	http.HandleFunc("/courses/", coursesHandlerWithID)
	http.HandleFunc("/coursesbylecturer/", coursesByLecturerHandler)
	http.HandleFunc("/lecturersbycity/", lecturersByCityHandler)
}

func lecturersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		withAdminAuth(controllers.GetLecturers)(w, r)
	case http.MethodPost:
		withAdminAuth(controllers.CreateLecturer)(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func lecturersHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[2]
	switch r.Method {
	case http.MethodGet:
		withAdminAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetLecturerByID(w, r, id)
		})(w, r)
	case http.MethodPatch:
		withAdminAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.UpdateLecturer(w, r, id)
		})(w, r)
	case http.MethodDelete:
		withAdminAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.DeleteLecturer(w, r, id)
		})(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		withAuth(controllers.GetCourses)(w, r)
	case http.MethodPost:
		withAuth(controllers.CreateCourse)(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func coursesHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[2]
	switch r.Method {
	case http.MethodGet:
		withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetCourseByID(w, r, id)
		})(w, r)
	case http.MethodPatch:
		withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.UpdateCourse(w, r, id)
		})(w, r)
	case http.MethodDelete:
		withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.DeleteCourse(w, r, id)
		})(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func coursesByLecturerHandler(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lecturerID := parts[2]
	switch r.Method {
	case http.MethodGet:
		withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetCoursesByLecturer(w, r, lecturerID)
		})(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func lecturersByCityHandler(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	city := parts[2]
	switch r.Method {
	case http.MethodGet:
		withAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetLecturersByCity(w, r, city)
		})(w, r)
		
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}