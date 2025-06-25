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

func withOwnAuth(handler func(http.ResponseWriter, *http.Request, string), id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wrappedControllerHandler := http.HandlerFunc(func(innerW http.ResponseWriter, innerR *http.Request) {
			handler(innerW, innerR, id)
		})

		middlewareChain := middleware.CourseOwnershipMiddleware(
			wrappedControllerHandler,
			id,
		)

		utils.ApplyMiddlewares(middlewareChain, middleware.AuthMiddleware).ServeHTTP(w, r)
	}
}

func withOldAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ApplyMiddlewares(handler, middleware.AuthMiddleware, middleware.OldMiddleware).ServeHTTP(w, r)
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
		withOldAuth(controllers.GetLecturers)(w, r)
	case http.MethodPost:
		withOldAuth(controllers.CreateLecturer)(w, r)
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
		withOldAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetLecturerByID(w, r, id)
		})(w, r)
	case http.MethodPatch:
		withOldAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.UpdateLecturer(w, r, id)
		})(w, r)
	case http.MethodDelete:
		withOldAuth(func(w http.ResponseWriter, r *http.Request) {
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
		withOwnAuth(controllers.GetCourseByID, id)(w, r)
	case http.MethodPatch:
		withOwnAuth(controllers.UpdateCourse, id)(w, r)
	case http.MethodDelete:
		withOwnAuth(controllers.DeleteCourse, id)(w, r)
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