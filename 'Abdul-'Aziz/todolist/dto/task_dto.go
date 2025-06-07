package dto

import (
	"tugas/todolist/models"
)



type UserWithTasks struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Tasks []models.Task `json:"tasks"`
}