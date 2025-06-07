package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"tugas/todolist/dto"
	"tugas/todolist/models"
	"tugas/todolist/repository"
)

func CreateTask(db *sql.DB, task models.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	return repository.InsertTask(db, task)
}

func GetAllTasksWithUser(db *sql.DB) ([]dto.UserWithTasks, error) {
	tasks, err := repository.SelectAllUsersWithTasks(db)
	if err != nil {
		return nil, errors.New("no tasks found")
	}
	return tasks, nil
}

func GetAllTasks(db *sql.DB) ([]models.Task, error) {
	tasks, err := repository.SelectAllTasks(db)
	if err != nil {
		return nil, errors.New("no tasks found")
	}
	return tasks, nil
}

func GetTaskByOnlyId(db *sql.DB, id string) (models.Task, error) {
	if id == "" {
		return models.Task{}, errors.New("task ID is required")
	}
	task, err := repository.SelectTaskOnlyByID(db, id)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func GetTaskByUserID(db *sql.DB, id string) (dto.UserWithTasks, error) {
	if id == "" {
		return dto.UserWithTasks{}, errors.New("task ID is required")
	}
	task, err := repository.SelectTaskByUserID(db, id)
	if err != nil {
		fmt.Println("err:" ,err)
		return dto.UserWithTasks{}, errors.New("task not found")
	}
	return task, nil
}

func UpdateTask(db *sql.DB, id string, task models.Task) error {
	if id == "" {
		return errors.New("task ID is required")
	}
	
	return repository.UpdateTaskByID(db, id, task)
}

func DeleteTask(db *sql.DB, id string) error {
	if id == "" {
		return errors.New("task ID is required")
	}
	return repository.DeleteTaskByID(db, id)
}
