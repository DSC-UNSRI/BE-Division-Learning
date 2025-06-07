package repository

import (
	"database/sql"
	"errors"
	"tugas/todolist/dto"
	"tugas/todolist/models"
)

func InsertTask(db *sql.DB, task models.Task) error {
	_, err := db.Exec("INSERT INTO tasks(title, description, is_done, user_id) VALUES(?, ?, ?, ?)", task.Title, task.Description, task.IsDone, task.UserID)
	return err
}

func SelectAllUsersWithTasks(db *sql.DB) ([]dto.UserWithTasks, error) {
	var result []dto.UserWithTasks

	query := `
		SELECT 
			u.id, u.name, u.email,
			t.id, t.title, t.description, t.is_done, t.user_id
		FROM users u
		LEFT JOIN tasks t ON t.user_id = u.id
		ORDER BY u.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userMap := make(map[int]*dto.UserWithTasks)

	for rows.Next() {
		var (
			userID    int
			userName  string
			userEmail string
			taskID    sql.NullInt64
			title     sql.NullString
			desc      sql.NullString
			isDone    sql.NullBool
			taskUserID sql.NullInt64
		)

		err := rows.Scan(
			&userID, &userName, &userEmail,
			&taskID, &title, &desc, &isDone, &taskUserID,
		)
		if err != nil {
			return nil, err
		}

		
		user, exists := userMap[userID]
		if !exists {
			user = &dto.UserWithTasks{
				ID:    userID,
				Name:  userName,
				Email: userEmail,
				Tasks: []models.Task{},
			}
			userMap[userID] = user
		}

		if taskID.Valid {
			task := models.Task{
				ID:          int(taskID.Int64),
				Title:       title.String,
				Description: desc.String,
				IsDone:      isDone.Bool,
				UserID:      int(taskUserID.Int64),
			}
			user.Tasks = append(user.Tasks, task)
		}
	}

	
	for _, user := range userMap {
		result = append(result, *user)
	}

	return result, nil
}

func SelectAllTasks(db *sql.DB) ([]models.Task, error) {
	query := `
		SELECT id, title, description, is_done,
		FROM tasks
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.IsDone,
			&task.UserID,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}


func SelectTaskOnlyByID(db *sql.DB, id string) (models.Task, error) {
	var task models.Task

	query := `
		SELECT id, title, description, is_done, user_id
		FROM tasks
		WHERE id = ?
	`
	err := db.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.IsDone,
		&task.UserID,
	)

	return task, err
}

func SelectTaskByUserID(db *sql.DB, id string) (dto.UserWithTasks, error) {
	var result dto.UserWithTasks
	isFirstRow := true

	query := `
		SELECT 
			u.id, u.name, u.email,
			t.id, t.title, t.description, t.is_done, t.user_id
		FROM users u
		LEFT JOIN tasks t ON t.user_id = u.id
		WHERE u.id = ?
	`

	rows, err := db.Query(query, id)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			userID    int
			userName  string
			userEmail string
			taskID    sql.NullInt64
			title     sql.NullString
			desc      sql.NullString
			isDone    sql.NullBool
			taskUserID sql.NullInt64
		)

		err := rows.Scan(
			&userID, &userName, &userEmail,
			&taskID, &title, &desc, &isDone, &taskUserID,
		)
		if err != nil {
			return result, err
		}

		if isFirstRow {
			result.ID = userID
			result.Name = userName
			result.Email = userEmail
			isFirstRow = false
		}

		if taskID.Valid {
			task := models.Task{
				ID:          int(taskID.Int64),
				Title:       title.String,
				Description: desc.String,
				IsDone:      isDone.Bool,
				UserID:      int(taskUserID.Int64),
			}
			result.Tasks = append(result.Tasks, task)
		}
	}

	return result, nil
}

func UpdateTaskByID(db *sql.DB, id string, task models.Task) error {
	var existingTask models.Task

	query := `
	SELECT id, title, description, is_done, user_id
	FROM tasks
	WHERE id = ?
	`
	err := db.QueryRow(query, id).Scan(
		&existingTask.ID,
		&existingTask.Title,
		&existingTask.Description,
		&existingTask.IsDone,
		&existingTask.UserID,
	)
	if err != nil {
		return errors.New("task not found")
	}

	
	if task.Title == "" {
		task.Title = existingTask.Title
	}
	if task.Description == "" {
		task.Description = existingTask.Description
	}
	
	task.UserID = existingTask.UserID 

	
	_, err = db.Exec(
		"UPDATE tasks SET title = ?, description = ?, is_done = ?, user_id = ? WHERE id = ?",
		task.Title, task.Description, task.IsDone, task.UserID, id,
	)
	if err != nil {
		return errors.New("failed to update task")
	}

	return nil
}


func DeleteTaskByID(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
