package user

import (
	"database/sql"
	"uts-zildjianvitosulaiman/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByID(id int) (*domain.User, error)
	FindSecurityQuestionByEmail(email string) (string, error)
	ResetPassword(email, newPasswordHash string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (name, email, password, tier, security_question, security_answer)
	          VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query,
		user.Name,
		user.Email,
		user.Password,
		user.Tier,
		user.SecurityQuestion,
		user.SecurityAnswer,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)

	return nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	query := `SELECT id, name, email, password, tier, security_question, security_answer, created_at 
	          FROM users WHERE email = ? AND deleted_at IS NULL`

	user := &domain.User{}

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Tier,
		&user.SecurityQuestion,
		&user.SecurityAnswer,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByID(id int) (*domain.User, error) {
	query := `SELECT id, name, email, tier, created_at FROM users WHERE id = ? AND deleted_at IS NULL`
	user := &domain.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Tier, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindSecurityQuestionByEmail(email string) (string, error) {
	var question string
	query := `SELECT security_question FROM users WHERE email = ? AND deleted_at IS NULL`
	err := r.db.QueryRow(query, email).Scan(&question)
	if err != nil {
		return "", err
	}
	return question, nil
}

func (r *userRepository) ResetPassword(email, newPasswordHash string) error {
	query := `UPDATE users SET password = ? WHERE email = ?`
	_, err := r.db.Exec(query, newPasswordHash, email)
	return err
}
