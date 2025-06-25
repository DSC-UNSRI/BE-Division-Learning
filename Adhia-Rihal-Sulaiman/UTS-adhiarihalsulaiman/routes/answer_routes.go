package routes

import (
	"net/http"
	"strings"
	"uts_adhia/controllers"
	"uts_adhia/middlewares"
)

func AnswerRoutes() {
	http.HandleFunc("/answers/question/", answerHandlerWithQuestionID)
	http.HandleFunc("/answers/", answerHandlerWithAnswerID)
	http.HandleFunc("/answers/best/", answerMarkBestHandler)
}

func answerHandlerWithQuestionID(w http.ResponseWriter, r *http.Request) {
	pathSegment := strings.TrimPrefix(r.URL.Path, "/answers/question/")
	parts := strings.Split(pathSegment, "/")
	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Question ID missing in path", http.StatusBadRequest)
		return
	}
	questionID := parts[0]

	switch r.Method {
	case http.MethodPost:
		middlewares.WithAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.CreateAnswer(w, r, questionID)
		})(w, r)
	case http.MethodGet:
		middlewares.WithAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetAnswersByQuestionID(w, r, questionID)
		})(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func answerHandlerWithAnswerID(w http.ResponseWriter, r *http.Request) {
	pathSegment := strings.TrimPrefix(r.URL.Path, "/answers/")
	parts := strings.Split(pathSegment, "/")
	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Answer ID missing in path", http.StatusBadRequest)
		return
	}
	answerID := parts[0]

	switch r.Method {
	case http.MethodGet:
		middlewares.WithAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.GetAnswerByAnswerID(w, r, answerID)
		})(w, r)
	case http.MethodPatch:
		middlewares.WithOwnsAnswerAuth(controllers.UpdateAnswer, answerID)(w, r)
	case http.MethodDelete:
		middlewares.WithOwnsAnswerAuth(controllers.DeleteAnswer, answerID)(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func answerMarkBestHandler(w http.ResponseWriter, r *http.Request) {
	pathSegment := strings.TrimPrefix(r.URL.Path, "/answers/best/")
	parts := strings.Split(pathSegment, "/")
	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Answer ID missing in path", http.StatusBadRequest)
		return
	}
	answerID := parts[0]

	switch r.Method {
	case http.MethodPost:
		middlewares.WithAuth(func(w http.ResponseWriter, r *http.Request) {
			controllers.MarkAnswerBest(w, r, answerID)
		})(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


