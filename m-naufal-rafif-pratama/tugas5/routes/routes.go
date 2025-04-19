package routes

import (
	"net/http"
	"strings"
	"tugas5/controllers"
	"tugas5/utils"
)

func SetupRoutes() {
	http.HandleFunc("/students", StudentHandler)
	http.HandleFunc("/students/", StudentHandlerWithID)
	http.HandleFunc("/organizations", OrganizationHandler)
	http.HandleFunc("/organizations/", OrganizationHandlerWithID)
	http.HandleFunc("/organizations/members/", OrganizationMembersHandler)
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
