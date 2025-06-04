package routes

import (
	"pertemuan05/controllers"
	"pertemuan05/utils"
	
	"net/http"
)



func RoutesHandlers() {
	http.HandleFunc("/login", controllers.LoginHandler)
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
		controllers.GetLecturers(w, r)
	case http.MethodPost:
		controllers.CreateLecturer(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func lecturersHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id := parts[2]
	switch r.Method {
	case http.MethodGet:
		controllers.GetLecturerByID(w, r, id)
	case http.MethodPatch:
		controllers.UpdateLecturer(w, r, id)
	case http.MethodDelete:
		controllers.DeleteLecturer(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetCourses(w, r)
	case http.MethodPost:
		controllers.CreateCourse(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func coursesHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id := parts[2]
	switch r.Method {
	case http.MethodGet:
		controllers.GetCourseByID(w, r, id)
	case http.MethodPatch:
		controllers.UpdateCourse(w, r, id)
	case http.MethodDelete:
		controllers.DeleteCourse(w, r, id)
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
		controllers.GetCoursesByLecturer(w, r, lecturerID)
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
		controllers.GetLecturersByCity(w, r, city)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}