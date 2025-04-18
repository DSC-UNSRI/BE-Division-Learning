package repositories

import (
	"database/sql"
	"errors"
	"tugas-5/models"
)

type ProgrammerRepository struct {
	db *sql.DB
}

func NewProgrammerRepository(db *sql.DB) *ProgrammerRepository {
	return &ProgrammerRepository{db: db}
}

func (r *ProgrammerRepository) Create(programmer *models.Programmer) error {
	if programmer == nil {
		return errors.New("programmer cannot be nil")
	}

	result, err := r.db.Exec("INSERT INTO programmers (name, email, language, years_of_experience, user_id) VALUES (?, ?, ?, ?, ?)",
		programmer.Name, programmer.Email, programmer.Language, programmer.YearsOfExperience, programmer.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	programmer.ID = int(id)
	return nil
}

func (r *ProgrammerRepository) FindByID(id int) (*models.Programmer, error) {
	programmer := &models.Programmer{}
	err := r.db.QueryRow("SELECT id, name, email, language, years_of_experience, created_at, user_id FROM programmers WHERE id = ? AND deleted_at IS NULL", id).
		Scan(&programmer.ID, &programmer.Name, &programmer.Email, &programmer.Language, &programmer.YearsOfExperience, &programmer.CreatedAt, &programmer.UserID)

	if err == sql.ErrNoRows {
		return nil, errors.New("programmer not found")
	}
	return programmer, err
}

func (r *ProgrammerRepository) FindAll() ([]*models.Programmer, error) {
	rows, err := r.db.Query("SELECT id, name, email, language, years_of_experience, created_at, user_id FROM programmers WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programmers []*models.Programmer
	for rows.Next() {
		p := &models.Programmer{}
		err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.Language, &p.YearsOfExperience, &p.CreatedAt, &p.UserID)
		if err != nil {
			return nil, err
		}
		programmers = append(programmers, p)
	}
	return programmers, nil
}

func (r *ProgrammerRepository) Update(programmer *models.Programmer) error {
	result, err := r.db.Exec("UPDATE programmers SET name = ?, language = ?, email = ?, years_of_experience = ?, user_id = ? WHERE id = ? AND deleted_at IS NULL",
		programmer.Name, programmer.Language, programmer.Email, programmer.YearsOfExperience, programmer.UserID, programmer.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("programmer not found or already deleted")
	}
	return nil
}

func (r *ProgrammerRepository) Delete(id int) error {
	_, err := r.db.Exec("UPDATE programmers SET deleted_at = NOW() WHERE id = ?", id)
	return err
}

func (r *ProgrammerRepository) FindByUserID(userID int) ([]*models.Programmer, error) {
	query := `SELECT id, name, email, language, years_of_experience, user_id, created_at, deleted_at
			  FROM programmers
			  WHERE user_id = ? AND deleted_at IS NULL`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programmers []*models.Programmer
	for rows.Next() {
		var p models.Programmer
		err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.Language, &p.YearsOfExperience, &p.UserID, &p.CreatedAt, &p.DeletedAt)
		if err != nil {
			return nil, err
		}
		programmers = append(programmers, &p)
	}

	return programmers, nil
}

func (r *ProgrammerRepository) CountByUserID(userID int) (int, error) {
	query := `SELECT COUNT(*) FROM programmers WHERE user_id = ? AND deleted_at IS NULL`

	var count int
	err := r.db.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
