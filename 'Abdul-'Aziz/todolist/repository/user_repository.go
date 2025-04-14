package repository

import (
	"database/sql"
	"tugas/todolist/models"
)

func InsertUser(db *sql.DB, user models.User) error {
	_, err := db.Exec("INSERT INTO users(name, password) VALUES(?, ?)", user.Name, user.Password)
	return err
}

func SelectAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}
	return users, nil
}

func SelectUserByID(db *sql.DB, id string) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	return user, err
}

func UpdateUserByID(db *sql.DB, id string, user models.User) error {
	_, err := db.Exec("UPDATE users SET name = ?, password = ? WHERE id = ?", user.Name, user.Password, id)
	return err
}

func DeleteUserByID(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

func ValidateLogin(db *sql.DB, name, password string) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, name FROM users WHERE name = ? AND password = ?", name, password).Scan(&user.ID, &user.Name)
	return user, err
}
