package routes

import (
	"net/http"
	"uts-gdg/controllers"
	"uts-gdg/middleware"
	"uts-gdg/utils"
)

func QuestionRoutes(){
	http.HandleFunc("/questions", questionHandler)
	http.HandleFunc("/questions/", questionHandlerWithID)
	http.HandleFunc("/questions/user", withPremiumAuth(controllers.GetQuestionsByUser))
}

func withAuth(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        utils.ApplyMiddlewares(handler, middleware.AuthMiddleware).ServeHTTP(w, r)
    }
}

func withPremiumAuth(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        utils.ApplyMiddlewares(handler, middleware.AuthMiddleware, middleware.PremiumMiddleware).ServeHTTP(w, r)
    }
}

func questionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		withAuth(controllers.GetQuestions)(w,r)
	case http.MethodPost:
		withPremiumAuth(controllers.CreateQuestions)(w,r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func questionHandlerWithID(w http.ResponseWriter, r *http.Request) {
	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id := parts[2]

	switch r.Method {
	case http.MethodGet:
		withAuth(func(w http.ResponseWriter, r *http.Request) {
            controllers.GetQuestion(w, r, id)
        })(w, r)
	case http.MethodPatch:
		withPremiumAuth(func(w http.ResponseWriter, r *http.Request) {
            controllers.UpdateQuestion(w, r, id)
        })(w, r)
	case http.MethodDelete:
		withPremiumAuth(func(w http.ResponseWriter, r *http.Request) {
            controllers.DeleteQuestion(w, r, id)
        })(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}