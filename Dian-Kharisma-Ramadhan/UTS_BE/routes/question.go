package routes

import (
	"UTS_BE/controllers"
	"UTS_BE/middleware"
	"UTS_BE/utils"
	"net/http"
)
func QuestionRoutes() {
	http.HandleFunc("/questions/create", middleware.WithAuth(controllers.CreateQuestion))
	http.HandleFunc("/questions/all", controllers.GetAllQuestions)
	http.HandleFunc("/questions/mine", middleware.WithAuth(controllers.GetMyQuestions))
	http.HandleFunc("/questions/", middleware.WithAuth(questionByIDHandler))
	http.HandleFunc("/questions/update/", middleware.WithAuth(updateQuestionHandler))
	http.HandleFunc("/questions/delete/", middleware.WithAuth(deleteQuestionHandler))
}

func questionByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil || len(parts) < 2 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	id := parts[1]
	controllers.GetQuestionByID(w, r, id)
}

func updateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil || len(parts) < 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	id := parts[2]
	controllers.UpdateQuestion(w, r, id)
}

func deleteQuestionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil || len(parts) < 3 {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}
	id := parts[2]
	controllers.DeleteQuestion(w, r, id)
}
