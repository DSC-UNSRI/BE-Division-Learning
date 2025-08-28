package model

type Major struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Batch  int    `json:"batch"`
	Number int    `json:"number"`
}
