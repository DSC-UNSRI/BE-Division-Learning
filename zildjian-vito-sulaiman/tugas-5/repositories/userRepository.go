package repositories

import (
	"database/sql"
	"errors"
	"tugas-5/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	result, err := r.db.Exec("INSERT INTO users (name, password, email) VALUES (?, ?, ?)",
		user.Name, user.Password, user.Email)
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

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ? AND deleted_at IS NULL", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}

func (r *UserRepository) FindAll() ([]*models.User, error) {
	rows, err := r.db.Query("SELECT id, name, email, created_at FROM users WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Update(user *models.User) error {
	result, err := r.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ? AND deleted_at IS NULL",
		user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found or already deleted")
	}
	return nil
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.db.Exec("UPDATE users SET deleted_at = NOW() WHERE id = ?", id)
	return err
}

func (r *UserRepository) FindByAuth(auth string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE password = ?", auth).
		Scan(&user.ID, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		return nil, errors.New("auth token is invalid")
	}
	return user, err
}
