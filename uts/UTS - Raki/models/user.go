package models

import "time"

type User struct {
	ID                    int        `json:"id"`
	Username              string     `json:"username"`
	Email                 string     `json:"email,omitempty"`
	PasswordHash          string     `json:"-"`
	UserType              string     `json:"user_type"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	RecoveryCode          *string    `json:"-"`
	RecoveryCodeExpiresAt *time.Time `json:"-"`
	DailyQuestionCount    int        `json:"daily_question_count"`
	DailyAnswerCount      int        `json:"daily_answer_count"`
	LastActivityDate      time.Time  `json:"last_activity_date"`
}

type UserRegistrationRequest struct {
	Username     string `json:"username" validate:"required,min=3,max=50"`
	Email        string `json:"email" validate:"omitempty,email"`
	Password     string `json:"password" validate:"required,min=6"`
	RecoveryCode string `json:"recovery_code" validate:"required,min=6"` 
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ForgotPasswordRequest struct {
	Username     string `json:"username" validate:"required"`
	RecoveryCode string `json:"recovery_code" validate:"required"`
}


type ResetPasswordRequest struct {
	Username         string `json:"username" validate:"required"`
	TempRecoveryCode string `json:"temp_recovery_code" validate:"required"` 
	NewPassword      string `json:"new_password" validate:"required,min=6"`
}