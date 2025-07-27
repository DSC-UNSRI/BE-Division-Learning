package model

type User struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender" gorm:"type:enum('man','woman')"`
}