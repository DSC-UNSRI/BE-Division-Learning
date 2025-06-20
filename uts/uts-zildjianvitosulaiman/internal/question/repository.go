package question

import (
	"database/sql"
	"uts-zildjianvitosulaiman/domain"
)

type Repository interface {
	Create(question *domain.Question) error
	FindAll() ([]*domain.Question, error)
	FindByID(id int) (*domain.Question, error)
	Update(question *domain.Question) error
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(q *domain.Question) error {
	query := `INSERT INTO questions (title, body, user_id) VALUES (?, ?, ?)`
	res, err := r.db.Exec(query, q.Title, q.Body, q.UserID)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	q.ID = int(id)
	return nil
}

func (r *repository) FindAll() ([]*domain.Question, error) {
	query := `SELECT id, title, body, user_id, created_at, updated_at FROM questions`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []*domain.Question
	for rows.Next() {
		var q domain.Question
		if err := rows.Scan(&q.ID, &q.Title, &q.Body, &q.UserID, &q.CreatedAt, &q.UpdatedAt); err != nil {
			return nil, err
		}
		questions = append(questions, &q)
	}
	return questions, nil
}

func (r *repository) FindByID(id int) (*domain.Question, error) {
	query := `SELECT id, title, body, user_id, created_at, updated_at FROM questions WHERE id = ?`
	var q domain.Question
	err := r.db.QueryRow(query, id).Scan(&q.ID, &q.Title, &q.Body, &q.UserID, &q.CreatedAt, &q.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *repository) Update(q *domain.Question) error {
	query := `UPDATE questions SET title = ?, body = ? WHERE id = ?`
	_, err := r.db.Exec(query, q.Title, q.Body, q.ID)
	return err
}

func (r *repository) Delete(id int) error {
	query := `DELETE FROM questions WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
