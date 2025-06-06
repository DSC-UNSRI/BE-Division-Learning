package models

type MahasiswaMinat struct {
	MahasiswaID int `json:"mahasiswa_id"`
	MinatID     int `json:"minat_id"`
}

var MahasiswaMinatQuery = `
CREATE TABLE IF NOT EXISTS mahasiswa_minat (
	mahasiswa_id INT NOT NULL,
	minat_id INT NOT NULL,
	PRIMARY KEY (mahasiswa_id, minat_id),
	FOREIGN KEY (mahasiswa_id) REFERENCES mahasiswa(id) ON DELETE CASCADE,
	FOREIGN KEY (minat_id) REFERENCES minat(id) ON DELETE CASCADE
);`