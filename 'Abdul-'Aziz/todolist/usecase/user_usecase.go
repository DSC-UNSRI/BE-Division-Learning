package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"tugas/todolist/lib"
	"tugas/todolist/models"
	"tugas/todolist/repository"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *sql.DB, user models.User) error {
	existingUser := repository.CheckUserExist(db, user.Email)

	if existingUser.Email != "" {
		return errors.New("user already exist. please use different email")
	}

	if user.Name == "" || user.Password == "" || user.Email == "" {
		return errors.New("please fill all the fields")
	}

	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	
	user.Password = string(hashedPassword)

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
		
	return repository.UpdateUserByID(db, id, user)
}

func DeleteUser(db *sql.DB, id string) error {
	fmt.Println("id", id)
	if id == "" {
		return errors.New("user ID is required")
	}
	return repository.DeleteUserByID(db, id)
}

func LoginUser(db *sql.DB, email, password string) (models.User, error) {
	if email == "" || password == "" {
		return models.User{}, errors.New("please fill all the fields")
	}

	user, err := repository.ValidateLogin(db, email, password)
	if err != nil {
		return models.User{}, err // misalnya: email/password salah
	}

	token, err := lib.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return models.User{}, errors.New("failed to generate token")
	}

	user.Token = &token

	return user, nil
}

