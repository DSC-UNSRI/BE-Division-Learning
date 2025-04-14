package repository

import (
	"database/sql"
	"tugas/todolist/models"
)

func InsertTask(db *sql.DB, task models.Task) error {
	_, err := db.Exec("INSERT INTO tasks(title, description, is_done, user_id) VALUES(?, ?, ?, ?)", task.Title, task.Description, task.IsDone, task.UserID)
	return err
}

func SelectAllTasks(db *sql.DB) ([]models.Task, error) {
	rows, err := db.Query("SELECT id, title, description, is_done, user_id FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsDone, &task.UserID)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func SelectTaskByID(db *sql.DB, id string) (models.Task, error) {
	var task models.Task
	err := db.QueryRow("SELECT id, title, description, is_done, user_id FROM tasks WHERE id = ?", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.IsDone, &task.UserID)
	return task, err
}

func UpdateTaskByID(db *sql.DB, id string, task models.Task) error {
	_, err := db.Exec("UPDATE tasks SET title = ?, description = ?, is_done = ?, user_id = ? WHERE id = ?", task.Title, task.Description, task.IsDone, task.UserID, id)
	return err
}

func DeleteTaskByID(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
