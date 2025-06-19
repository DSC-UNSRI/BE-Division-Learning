package repository

import (
	"database/sql"
	"time"

	"github.com/artichys/uts-raki/models" 
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password_hash, user_type, recovery_code, daily_question_count, daily_answer_count, last_activity_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, user.Username, user.Email, user.PasswordHash, user.UserType, user.RecoveryCode, user.DailyQuestionCount, user.DailyAnswerCount, user.LastActivityDate)
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password_hash, user_type, created_at, updated_at, recovery_code, recovery_code_expires_at, daily_question_count, daily_answer_count, last_activity_date FROM users WHERE username = ?`
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.UserType,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.RecoveryCode,
		&user.RecoveryCodeExpiresAt,
		&user.DailyQuestionCount,
		&user.DailyAnswerCount,  
		&user.LastActivityDate,   
	)
	if err == sql.ErrNoRows {
		return nil, nil 
	}
	return user, err
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, user_type, created_at, updated_at, daily_question_count, daily_answer_count, last_activity_date FROM users WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.UserType,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DailyQuestionCount,
		&user.DailyAnswerCount,
		&user.LastActivityDate,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (r *UserRepository) UpdateUserPassword(userID int, newPasswordHash string) error {
	query := `UPDATE users SET password_hash = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	_, err := r.db.Exec(query, newPasswordHash, userID)
	return err
}

func (r *UserRepository) UpdateUserRecoveryCode(userID int, code *string, expiresAt *time.Time) error {
	query := `UPDATE users SET recovery_code = ?, recovery_code_expires_at = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	_, err := r.db.Exec(query, code, expiresAt, userID)
	return err
}

func (r *UserRepository) UpdateUserDailyCounts(userID int, questionCount, answerCount int, lastActivityDate time.Time) error {
	query := `UPDATE users SET daily_question_count = ?, daily_answer_count = ?, last_activity_date = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	_, err := r.db.Exec(query, questionCount, answerCount, lastActivityDate, userID)
	return err
}