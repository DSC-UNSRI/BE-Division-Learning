package routes

import(
	"uts-gdg/controllers"
	"uts-gdg/middleware"
	"uts-gdg/utils"
	"net/http"
)

func QuestionRoutes(){
	http.HandleFunc("/questions", questionHandler)
	//http.HandleFunc("/questions/", questionHandlerWithID)
}

func withPremiumAuth(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        utils.ApplyMiddlewares(handler, middleware.AuthMiddleware, middleware.PremiumMiddleware).ServeHTTP(w, r)
    }
}

func questionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetQuestions(w,r)
	case http.MethodPost:
		withPremiumAuth(controllers.CreateQuestions)(w,r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}