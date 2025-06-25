package routes

import (
	"net/http"
	"strings"
	"uts_adhia/controllers"
	"uts_adhia/middlewares"
)

func QuestionRoutes() {
	http.HandleFunc("/questions", questionHandler)
	http.HandleFunc("/questions/", questionHandlerWithID)
	http.HandleFunc("/questions/best/", questionMarkBestHandler)
}

func questionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		middlewares.WithAuth(controllers.CreateQuestion)(w, r)
	case http.MethodGet:
		middlewares.WithAuth(controllers.GetQuestions)(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func questionHandlerWithID(w http.ResponseWriter, r *http.Request) {
	pathSegment := strings.TrimPrefix(r.URL.Path, "/questions/")
	parts := strings.Split(pathSegment, "/")
	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Question ID missing in path", http.StatusBadRequest)
		return
	}
	questionID := parts[0]

	switch r.Method {
	case http.MethodGet:
		middlewares.WithAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetQuestionByID(w, r, questionID)
		})(w, r)
	case http.MethodPatch:
		middlewares.WithOwnsQuestionAuth(controllers.UpdateQuestion, questionID)(w, r)
	case http.MethodDelete:
		middlewares.WithOwnsQuestionAuth(controllers.DeleteQuestion, questionID)(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func questionMarkBestHandler(w http.ResponseWriter, r *http.Request) {
	pathSegment := strings.TrimPrefix(r.URL.Path, "/questions/best/")
	parts := strings.Split(pathSegment, "/")
	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Question ID missing in path", http.StatusBadRequest)
		return
	}
	questionID := parts[0]

	switch r.Method {
	case http.MethodPost:
		middlewares.WithAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.MarkQuestionBest(w, r, questionID)
		})(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

