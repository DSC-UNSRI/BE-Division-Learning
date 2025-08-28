package models

import "time"

type Minat struct {
	ID          int        `json:"id"`
	NamaMinat   string     `json:"nama_minat"`
	Deskripsi   string     `json:"deskripsi,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

var MinatQuery = `
CREATE TABLE IF NOT EXISTS minat (
	id INT AUTO_INCREMENT PRIMARY KEY,
	nama_minat VARCHAR(100) NOT NULL UNIQUE,
	deskripsi TEXT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL DEFAULT NULL
);`