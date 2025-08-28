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
	MaxFreeAnswersPerDay = 5
)

type AnswerHandler struct {
	answerRepo   *repository.AnswerRepository
	questionRepo *repository.QuestionRepository
	userRepo     *repository.UserRepository
}

func NewAnswerHandler(answerRepo *repository.AnswerRepository, questionRepo *repository.QuestionRepository, userRepo *repository.UserRepository) *AnswerHandler {
	return &AnswerHandler{answerRepo: answerRepo, questionRepo: questionRepo, userRepo: userRepo}
}

func (h *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "User ID not found in context")
		return
	}
	userType := r.Context().Value("userType").(string)


	vars := mux.Vars(r)
	questionIDStr := vars["question_id"]
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid question ID format")
		return
	}

	question, err := h.questionRepo.GetQuestionByID(questionID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Database error checking question: " + err.Error())
		return
	}
	if question == nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Question not found")
		return
	}

	var req models.CreateAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Content == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Answer content is required")
		return
	}
	if len(req.Content) < 10 {
		utils.ErrorResponse(w, http.StatusBadRequest, "Answer content must be at least 10 characters long")
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

		if user.DailyAnswerCount >= MaxFreeAnswersPerDay {
			utils.ErrorResponse(w, http.StatusForbidden, "Free users are limited to " + strconv.Itoa(MaxFreeAnswersPerDay) + " answers per day.")
			return
		}
		user.DailyAnswerCount++
		if err := h.userRepo.UpdateUserDailyCounts(userID, user.DailyQuestionCount, user.DailyAnswerCount, user.LastActivityDate); err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update daily answer count: " + err.Error())
			return
		}
	}


	answer := &models.Answer{
		QuestionID: questionID,
		UserID:     userID,
		Content:    req.Content,
	}

	if err := h.answerRepo.CreateAnswer(answer); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create answer: " + err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]string{"message": "Answer created successfully"})
}

func (h *AnswerHandler) GetAnswersByQuestionID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questionIDStr := vars["question_id"]
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid question ID format")
		return
	}

	question, err := h.questionRepo.GetQuestionByID(questionID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Database error checking question: " + err.Error())
		return
	}
	if question == nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Question not found")
		return
	}


	answers, err := h.answerRepo.GetAnswersByQuestionID(questionID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve answers: " + err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, answers)
}

func (h *AnswerHandler) UpdateAnswer(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "User ID not found in context")
		return
	}

	vars := mux.Vars(r)
	questionIDStr := vars["question_id"]
	answerIDStr := vars["answer_id"]

	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid question ID format")
		return
	}
	answerID, err := strconv.Atoi(answerIDStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid answer ID format")
		return
	}

	var req models.CreateAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Content == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Answer content is required")
		return
	}

	answer := &models.Answer{
		ID:         answerID,
		QuestionID: questionID,
		UserID:     userID,
		Content:    req.Content,
	}

	if err := h.answerRepo.UpdateAnswer(answer); err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorResponse(w, http.StatusNotFound, "Answer not found or you don't have permission to update it")
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update answer: " + err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "Answer updated successfully"})
}

func (h *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "User ID not found in context")
		return
	}

	vars := mux.Vars(r)
	questionIDStr := vars["question_id"]
	answerIDStr := vars["answer_id"]

	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid question ID format")
		return
	}
	answerID, err := strconv.Atoi(answerIDStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid answer ID format")
		return
	}

	if err := h.answerRepo.DeleteAnswer(answerID, questionID, userID); err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorResponse(w, http.StatusNotFound, "Answer not found or you don't have permission to delete it")
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to delete answer: " + err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "Answer deleted successfully"})
}