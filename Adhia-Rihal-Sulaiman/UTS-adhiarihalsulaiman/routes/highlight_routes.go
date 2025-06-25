package routes

import (
	"net/http"
	"strings"
	"uts_adhia/controllers"
	"uts_adhia/middlewares"
)

func HighlightRoutes() {
	http.HandleFunc("/highlights", highlightGenericHandler)
	http.HandleFunc("/highlights/", highlightGenericHandlerWithID)
}

func highlightGenericHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		middlewares.WithAdminOrPremiumAuth(controllers.CreateHighlight)(w, r)
	case http.MethodGet:
		middlewares.WithAdminOrPremiumAuth(controllers.GetAllHighlights)(w, r)
	default:
		http.Error(w, "Method not allowed for /highlights", http.StatusMethodNotAllowed)
	}
}

func highlightGenericHandlerWithID(w http.ResponseWriter, r *http.Request) {
	pathSegment := strings.TrimPrefix(r.URL.Path, "/highlights/")
	parts := strings.Split(pathSegment, "/")
	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Highlight ID missing in path", http.StatusBadRequest)
		return
	}
	highlightID := parts[0]

	switch r.Method {
	case http.MethodGet:
		middlewares.WithAdminOrPremiumAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetHighlightByID(w, r, highlightID)
		})(w, r)
	case http.MethodDelete:
		middlewares.WithAdminOrPremiumAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.DeleteHighlight(w, r, highlightID)
		})(w, r)

	default:
		http.Error(w, "Method not allowed for /highlights/{id}", http.StatusMethodNotAllowed)
	}
}
