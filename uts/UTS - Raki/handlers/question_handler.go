package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/artichys/uts-raki/models"
	"github.com/artichys/uts-raki/repository"
	"github.com/artichys/uts-raki/utils"
)

const (
	MaxFreeQuestionsPerDay = 10
)

type QuestionHandler struct {
	questionRepo *repository.QuestionRepository
	userRepo     *repository.UserRepository
}

func NewQuestionHandler(questionRepo *repository.QuestionRepository, userRepo *repository.UserRepository) *QuestionHandler {
	return &QuestionHandler{questionRepo: questionRepo, userRepo: userRepo}
}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "User ID not found in context")
		return
	}
	userType := r.Context().Value("userType").(string)

	var req models.CreateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Title == "" || req.Content == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Title and Content are required")
		return
	}
	if len(req.Title) < 5 || len(req.Content) < 10 {
		utils.ErrorResponse(w, http.StatusBadRequest, "Title must be at least 5 chars, Content at least 10 chars")
		return
	}

	if userType == "free" {
		user, err := h.userRepo.GetUserByID(userID)
		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to get user data for limit check: " + err.Error())
			return
		}

		today := time.Now().Truncate(24 * time.Hour)

		if user.LastActivityDate.Before(today) {
			user.DailyQuestionCount = 0
			user.DailyAnswerCount = 0
			user.LastActivityDate = today
		}

		if user.DailyQuestionCount >= MaxFreeQuestionsPerDay {
			utils.ErrorResponse(w, http.StatusForbidden, "Free users are limited to " + strconv.Itoa(MaxFreeQuestionsPerDay) + " questions per day.")
			return
		}
		user.DailyQuestionCount++
		if err := h.userRepo.UpdateUserDailyCounts(userID, user.DailyQuestionCount, user.DailyAnswerCount, user.LastActivityDate); err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update daily question count: " + err.Error())
			return
		}
	}

	question := &models.Question{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
		IsPromoted: false,
	}

	if err := h.questionRepo.CreateQuestion(question); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create question: " + err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]string{"message": "Question created successfully"})
}

func (h *QuestionHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.questionRepo.GetAllQuestions()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve questions: " + err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, questions)
}

func (h *QuestionHandler) GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid question ID format")
		return
	}

	question, err := h.questionRepo.GetQuestionByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorResponse(w, http.StatusNotFound, "Question not found")
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Database error: " + err.Error())
		return
	}
	if question == nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Question not found")
		return
	}

	utils.JSONResponse(w, http.StatusOK, question)
}

func (h *QuestionHandler) GetMyQuestions(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "User ID not found in context")
		return
	}

	questions, err := h.questionRepo.GetUserQuestions(userID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve your questions: " + err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, questions)
}

func (h *QuestionHandler) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "User ID not found in context")
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	questionID, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid question ID format")
		return
	}

	var req models.CreateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Title == "" || req.Content == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Title and Content are required")
		return
	}

	question := &models.Question{
		ID:      questionID,
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.questionRepo.UpdateQuestion(question); err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorResponse(w, http.StatusNotFound, "Question not found or you don't have permission to update it")
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update question: " + err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "Question updated successfully"})
}

func (h *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "User ID not found in context")
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	questionID, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid question ID format")
		return
	}

	if err := h.questionRepo.DeleteQuestion(questionID, userID); err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorResponse(w, http.StatusNotFound, "Question not found or you don't have permission to delete it")
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to delete question: " + err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "Question deleted successfully"})
}

func (h *QuestionHandler) PromoteQuestion(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "User ID not found in context")
		return
	}

	vars := mux.Vars(r)
	questionIDStr := vars["id"]
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid question ID format")
		return
	}

	if err := h.questionRepo.PromoteQuestion(questionID, userID); err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorResponse(w, http.StatusNotFound, "Question not found or you don't own this question.")
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to promote question: " + err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "Question successfully promoted!"})
}