package usecase

import (
	"database/sql"
	"errors"
	"tugas/todolist/models"
	"tugas/todolist/repository"
)

func CreateUser(db *sql.DB, user models.User) error {
	if user.Name == "" || user.Password == "" {
		return errors.New("name and password are required")
	}
	return repository.InsertUser(db, user)
}

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	users, err := repository.SelectAllUsers(db)
	if err != nil {
		return nil, errors.New("no users found")
	}

	return users, nil
}

func GetUserByID(db *sql.DB, id string) (models.User, error) {
	if id == "" {
		return models.User{}, errors.New("user ID is required")
	}
	user, err := repository.SelectUserByID(db, id)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func UpdateUser(db *sql.DB, id string, user models.User) error {
	if id == "" {
		return errors.New("user ID is required")
	}
	if user.Name == "" || user.Password == "" {
		return errors.New("name and password are required")
	}
	return repository.UpdateUserByID(db, id, user)
}

func DeleteUser(db *sql.DB, id string) error {
	if id == "" {
		return errors.New("user ID is required")
	}
	return repository.DeleteUserByID(db, id)
}

func LoginUser(db *sql.DB, name, password string) (models.User, error) {
	if name == "" || password == "" {
		return models.User{}, errors.New("name and password are required")
	}
	return repository.ValidateLogin(db, name, password)
}
