package answer

import (
	"database/sql"
	"uts-zildjianvitosulaiman/domain"
)

type Repository interface {
	Create(answer *domain.Answer) error
	FindByQuestionID(questionID int) ([]*domain.Answer, error)
	FindByID(id int) (*domain.Answer, error)
	Update(answer *domain.Answer) error
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(a *domain.Answer) error {
	query := `INSERT INTO answers (body, user_id, question_id) VALUES (?, ?, ?)`
	res, err := r.db.Exec(query, a.Body, a.UserID, a.QuestionID)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	a.ID = int(id)
	return nil
}

func (r *repository) FindByQuestionID(questionID int) ([]*domain.Answer, error) {
	query := `SELECT id, body, user_id, question_id, upvotes, downvotes, created_at, updated_at 
	          FROM answers WHERE question_id = ?`
	rows, err := r.db.Query(query, questionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var answers []*domain.Answer
	for rows.Next() {
		var a domain.Answer
		err := rows.Scan(&a.ID, &a.Body, &a.UserID, &a.QuestionID, &a.Upvotes, &a.Downvotes, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}
		answers = append(answers, &a)
	}
	return answers, nil
}

func (r *repository) FindByID(id int) (*domain.Answer, error) {
	query := `SELECT id, body, user_id, question_id, upvotes, downvotes, created_at, updated_at 
	          FROM answers WHERE id = ?`
	var a domain.Answer
	err := r.db.QueryRow(query, id).Scan(&a.ID, &a.Body, &a.UserID, &a.QuestionID, &a.Upvotes, &a.Downvotes, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *repository) Update(a *domain.Answer) error {
	query := `UPDATE answers SET body = ? WHERE id = ?`
	_, err := r.db.Exec(query, a.Body, a.ID)
	return err
}

func (r *repository) Delete(id int) error {
	query := `DELETE FROM answers WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
