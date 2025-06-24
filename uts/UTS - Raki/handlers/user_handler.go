
package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/artichys/uts-raki/repository"
	"github.com/artichys/uts-raki/utils"   
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

func (h *UserHandler) GetMyProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "User ID not found in context")
		return
	}

	user, err := h.userRepo.GetUserByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorResponse(w, http.StatusNotFound, "User not found")
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, user)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID format")
		return
	}

	user, err := h.userRepo.GetUserByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorResponse(w, http.StatusNotFound, "User not found")
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	if user == nil {
		utils.ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	utils.JSONResponse(w, http.StatusOK, user)
}