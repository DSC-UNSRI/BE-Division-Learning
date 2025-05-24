package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"tugas/tugas4/infrastructure/env"
	"tugas/tugas4/models"
)

const (
	dataFilePath = "user_data.json"
)

type FileUserRepo struct {
	user     *models.User
	filePath string
}

// NewFileUserRepo membuat repository user yang menyimpan data ke file
func NewFileUserRepo(env *env.EnvLoader) UserRepository {
	// Buat direktori data jika belum ada
	os.MkdirAll("data", 0755)

	// Path absolut ke file data
	absPath, _ := filepath.Abs(dataFilePath)

	repo := &FileUserRepo{
		filePath: absPath,
		user:     &models.User{},
	}

	// Coba baca data dari file
	if err := repo.loadFromFile(); err != nil {
		// Jika file tidak ada, buat user baru
		repo.user = &models.User{
			Name:            env.Get("NAMA"),
			Email:           env.Get("EMAIL"),
			Password:        env.Get("PASSWORD"),
			Vehicle:         "",
			Equipment:       []string{},
			Recommendations: []models.Recommendation{},
			Friends:         []models.Friend{},
		}
		// Simpan user baru ke file
		repo.saveToFile()
	}

	return repo
}

// loadFromFile memuat data user dari file
func (r *FileUserRepo) loadFromFile() error {
	// Cek apakah file ada
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		return err
	}

	// Baca file
	data, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	// Parse JSON
	var user models.User
	if err := json.Unmarshal(data, &user); err != nil {
		return err
	}

	r.user = &user
	return nil
}

// saveTo0e menyimpan data user ke file
func (r *FileUserRepo) saveToFile() error {
	// Marshal user ke JSON
	data, err := json.MarshalIndent(r.user, "", "  ")
	if err != nil {
		return err
	}

	// Tulis ke file
	return ioutil.WriteFile(r.filePath, data, 0644)
}

// GetUser mengambil data user
func (r *FileUserRepo) GetUser() (*models.User, error) {
	return r.user, nil
}

// SaveUser menyimpan data user baru
func (r *FileUserRepo) SaveUser(user *models.User) error {
	r.user = user
	return r.saveToFile()
}

// UpdateUser memperbarui data user
func (r *FileUserRepo) UpdateUser(user *models.User) error {
	r.user = user
	return r.saveToFile()
}

// DeleteUser menghapus data user
func (r *FileUserRepo) DeleteUser() error {
	r.user = nil
	return r.saveToFile()
}
