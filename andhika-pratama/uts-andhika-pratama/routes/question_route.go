package routes

import (
	"uts/controllers"
	"uts/middlewares"
	"uts/utils"

	"net/http"
)

func QuestionRoutes() {
	http.HandleFunc("/questions", questionHandler)
	http.HandleFunc("/questions/", questionHandlerWithID)
	http.HandleFunc("/questions/upvote/", questionUpvoteHandler)
	http.HandleFunc("/questions/downvote/", questionDownvoteHandler)
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
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[2]
	switch r.Method {
	case http.MethodGet:
		middlewares.WithAuth(func(w http.ResponseWriter, r *http.Request) {
        controllers.GetQuestionByID(w, r, id)
    })(w, r)
	case http.MethodPatch:
		middlewares.WithOwnsQuestionAuth(controllers.UpdateQuestion, id)(w, r)
	case http.MethodDelete:
		middlewares.WithOwnsQuestionAuth(controllers.DeleteQuestion, id)(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func questionUpvoteHandler(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[3]
	switch r.Method {
	case http.MethodPost:
		middlewares.WithPremiumAuth(func(w http.ResponseWriter, r *http.Request) {
        controllers.UpvoteQuestion(w, r, id)
    })(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func questionDownvoteHandler(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[3]
	switch r.Method {
	case http.MethodPost:
		middlewares.WithPremiumAuth(func(w http.ResponseWriter, r *http.Request) {
        controllers.DownvoteQuestion(w, r, id)
    })(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
