package routes

import (
	"net/http"
	"uts-gdg/controllers"
	"uts-gdg/utils"
)

func AnswerRoutes(){
	http.HandleFunc("/answers", answerHandler)
	http.HandleFunc("/answers/", answerHandlerWithID)
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		withAuth(controllers.CreateAnswer)(w,r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func answerHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id := parts[2]

	switch r.Method {
	case http.MethodGet:
		withAuth(func(w http.ResponseWriter, r *http.Request) {
            controllers.GetAnswersByQuestionID(w, r, id)
        })(w, r)
	case http.MethodPatch:
		withPremiumAuth(func(w http.ResponseWriter, r *http.Request) {
            controllers.UpdateAnswer(w, r, id)
        })(w, r)
	case http.MethodDelete:
		withPremiumAuth(func(w http.ResponseWriter, r *http.Request) {
            controllers.DeleteAnswer(w, r, id)
        })(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
