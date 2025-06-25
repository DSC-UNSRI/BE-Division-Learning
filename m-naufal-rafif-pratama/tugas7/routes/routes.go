package routes

import (
	"net/http"
	"strings"
	"tugas7/controllers"
	"tugas7/middleware"
	"tugas7/utils"
)

func chainMiddleware(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

func SetupRoutes() {
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.CreateStudent)

	http.HandleFunc("/students", chainMiddleware(StudentHandler, middleware.OrgAuthMiddleware, middleware.AuthMiddleware))
	http.HandleFunc("/students/", chainMiddleware(StudentHandlerWithID, middleware.OrgAuthMiddleware, middleware.AuthMiddleware))
	http.HandleFunc("/organizations", chainMiddleware(OrganizationHandler, middleware.OrgAuthMiddleware, middleware.AuthMiddleware))
	http.HandleFunc("/organizations/", chainMiddleware(OrganizationHandlerWithID, middleware.OrgAuthMiddleware, middleware.AuthMiddleware))
	http.HandleFunc("/organizations/members/", chainMiddleware(OrganizationMembersHandler, middleware.OrgAuthMiddleware, middleware.AuthMiddleware))
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetAllStudents(w, r)
	case http.MethodPost:
		controllers.CreateStudent(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func StudentHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[2]
	
	switch r.Method {
	case http.MethodGet:
		controllers.GetStudentByID(w, r, id)
	case http.MethodPut:
		controllers.UpdateStudent(w, r, id)
	case http.MethodDelete:
		controllers.DeleteStudent(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func OrganizationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetAllOrganizations(w, r)
	case http.MethodPost:
		controllers.CreateOrganization(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func OrganizationHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[2]
	
	switch r.Method {
	case http.MethodGet:
		controllers.GetOrganizationByID(w, r, id)
	case http.MethodPut:
		controllers.UpdateOrganization(w, r, id)
	case http.MethodDelete:
		controllers.DeleteOrganization(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func OrganizationMembersHandler(w http.ResponseWriter, r *http.Request) {
	orgID := strings.TrimPrefix(r.URL.Path, "/organizations/members/")
	if orgID == "" {
		http.Error(w, "Organization ID is required", http.StatusBadRequest)
		return
	}
	
	if r.Method == http.MethodGet {
		controllers.GetOrganizationMembers(w, r, orgID)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
