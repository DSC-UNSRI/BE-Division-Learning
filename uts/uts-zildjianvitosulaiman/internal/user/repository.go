package user

import (
	"database/sql"
	"uts-zildjianvitosulaiman/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
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
