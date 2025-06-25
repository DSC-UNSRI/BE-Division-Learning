package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"tugas/todolist/models"

	"golang.org/x/crypto/bcrypt"
)

func InsertUser(db *sql.DB, user models.User) error {
	_, err := db.Exec("INSERT INTO users(name, password, email) VALUES(?, ?, ?)", user.Name, user.Password, user.Email)
	return err
}

func SelectAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.Email)
		users = append(users, user)
	}
	return users, nil
}

func SelectUserByID(db *sql.DB, id string) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	fmt.Println("error ", err)
	return user, err
}


func UpdateUserByID(db *sql.DB, id string, user models.User) error {
	var existingUser models.User
	err := db.QueryRow("SELECT name, password, email FROM users WHERE id = ?", id).Scan(
		&existingUser.Name, &existingUser.Password, &existingUser.Email,
	)
	if err != nil {
		return err
	}

	if user.Name == "" {
		user.Name = existingUser.Name
	}
	if user.Email == "" {
		user.Email = existingUser.Email
	}

	if user.Password == "" {
		user.Password = existingUser.Password
	} else {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	_, err = db.Exec(
		"UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?",
		user.Name, user.Email, user.Password, id,
	)
	return err
}


func DeleteUserByID(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	fmt.Println("err", err)
	return err
}

func ValidateLogin(db *sql.DB, email, password string) (models.User, error) {
	var user models.User
	var hashedPassword string

	err := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Name, &user.Email, &hashedPassword)
	if err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return models.User{}, errors.New("invalid password")
	}


	return user, nil
}

func CheckUserExist (db *sql.DB, email string) (models.User) {
	var user models.User
	_ = db.QueryRow("SELECT email FROM users WHERE email = ?", email).Scan(&user.Email)

	return user
}
