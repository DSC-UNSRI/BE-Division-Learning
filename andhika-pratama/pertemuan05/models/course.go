package models

import "time"

type Course struct {
	CourseID   int        `json:"course_id"`
	Name       string     `json:"name"`
	LecturerID int        `json:"lecturer_id"`
	Semester   int        `json:"semester"`
	Credit     int        `json:"credit"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

var CourseQuery = `
	CREATE TABLE IF NOT EXISTS courses (
		course_id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(40) NOT NULL,
		lecturer_id INT NOT NULL,
		semester INT NOT NULL,
		credit INT NOT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		FOREIGN KEY (lecturer_id) REFERENCES lecturers(lecturer_id)
	);`
