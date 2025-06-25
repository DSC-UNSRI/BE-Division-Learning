package routes

import (
	"uts/controllers"
	"uts/middlewares"
	"uts/utils"

	"net/http"
)

func AnswerRoutes() {
	http.HandleFunc("/answers/question/", answerHandlerWithQuestionID)
	http.HandleFunc("/answers/", answerHandlerWithAnswerID)
	http.HandleFunc("/answers/upvote/", answerUpvoteHandler)
	http.HandleFunc("/answers/downvote/", answerDownvoteHandler)
}

func answerHandlerWithQuestionID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[3]
	switch r.Method {
	case http.MethodPost:
		middlewares.WithAuth(func(w http.ResponseWriter, r *http.Request) {
        controllers.CreateAnswer(w, r, id)
    })(w, r)
	case http.MethodGet:
		middlewares.WithAuth(func(w http.ResponseWriter, r *http.Request) {
        controllers.GetAnswersByQuestionID(w, r, id)
    })(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func answerHandlerWithAnswerID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[2]
	switch r.Method {
	case http.MethodGet:
		middlewares.WithAuth(func(w http.ResponseWriter, r *http.Request) {
        controllers.GetAnswerByAnswerID(w, r, id)
    })(w, r)
	case http.MethodPatch:
		middlewares.WithOwnsAnswerAuth(controllers.UpdateAnswer, id)(w, r)
	case http.MethodDelete:
		middlewares.WithOwnsAnswerAuth(controllers.DeleteAnswer, id)(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func answerUpvoteHandler(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[3]
	switch r.Method {
	case http.MethodPost:
		middlewares.WithPremiumAuth(func(w http.ResponseWriter, r *http.Request) {
        controllers.UpvoteAnswer(w, r, id)
    })(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func answerDownvoteHandler(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := parts[3]
	switch r.Method {
	case http.MethodPost:
		middlewares.WithPremiumAuth(func(w http.ResponseWriter, r *http.Request) {
        controllers.DownvoteAnswer(w, r, id)
    })(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
