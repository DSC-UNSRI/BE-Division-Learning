package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/artichys/uts-raki/models"   
	"github.com/artichys/uts-raki/repository" 
	"github.com/artichys/uts-raki/utils"    

	"github.com/google/uuid" 
)

type AuthHandler struct {
	userRepo *repository.UserRepository
}

func NewAuthHandler(userRepo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req models.UserRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Username == "" || req.Password == "" || req.RecoveryCode == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Username, Password, and Recovery Code are required")
		return
	}
	if len(req.Password) < 6 {
		utils.ErrorResponse(w, http.StatusBadRequest, "Password must be at least 6 characters long")
		return
	}
	if len(req.RecoveryCode) < 6 {
		utils.ErrorResponse(w, http.StatusBadRequest, "Recovery Code must be at least 6 characters long")
		return
	}


	existingUser, err := h.userRepo.GetUserByUsername(req.Username)
	if err != nil && err != sql.ErrNoRows {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	if existingUser != nil {
		utils.ErrorResponse(w, http.StatusConflict, "Username already exists")
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	hashedRecoveryCode, err := utils.HashPassword(req.RecoveryCode)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to hash recovery code")
		return
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		UserType:     "free", 
	}
	user.RecoveryCode = &hashedRecoveryCode 
	if err := h.userRepo.CreateUser(user); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to register user: "+err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Username == "" || req.Password == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Username and Password are required")
		return
	}

	user, err := h.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	if user == nil || !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.UserType)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{"token": token, "user_type": user.UserType})
}

func (h *AuthHandler) InitiatePasswordReset(w http.ResponseWriter, r *http.Request) {
	var req models.ForgotPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Username == "" || req.RecoveryCode == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Username and Recovery Code are required")
		return
	}

	user, err := h.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	if user == nil {
		utils.ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	if user.RecoveryCode == nil || !utils.CheckPasswordHash(req.RecoveryCode, *user.RecoveryCode) {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid recovery code")
		return
	}

	tempRecoveryCode := uuid.New().String()
	expiresAt := time.Now().Add(10 * time.Minute)

	hashedTempRecoveryCode, err := utils.HashPassword(tempRecoveryCode)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to hash temporary recovery code")
		return
	}

	if err := h.userRepo.UpdateUserRecoveryCode(user.ID, &hashedTempRecoveryCode, &expiresAt); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update recovery code")
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{
		"message":          "Temporary recovery code generated. Use it to reset your password.",
		"temp_recovery_code": tempRecoveryCode, 
		"expires_at":       expiresAt.Format(time.RFC3339),
	})
}

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req models.ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Username == "" || req.TempRecoveryCode == "" || req.NewPassword == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Username, Temporary Recovery Code, and New Password are required")
		return
	}
	if len(req.NewPassword) < 6 {
		utils.ErrorResponse(w, http.StatusBadRequest, "New password must be at least 6 characters long")
		return
	}

	user, err := h.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	if user == nil {
		utils.ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	if user.RecoveryCode == nil || user.RecoveryCodeExpiresAt == nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Password reset process not initiated or expired.")
		return
	}
	if time.Now().After(*user.RecoveryCodeExpiresAt) {
		_ = h.userRepo.UpdateUserRecoveryCode(user.ID, nil, nil)
		utils.ErrorResponse(w, http.StatusUnauthorized, "Temporary recovery code expired.")
		return
	}

	if !utils.CheckPasswordHash(req.TempRecoveryCode, *user.RecoveryCode) {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid temporary recovery code.")
		return
	}

	hashedNewPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to hash new password")
		return
	}

	if err := h.userRepo.UpdateUserPassword(user.ID, hashedNewPassword); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update password: "+err.Error())
		return
	}

	if err := h.userRepo.UpdateUserRecoveryCode(user.ID, nil, nil); err != nil {
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "Password reset successfully."})
}