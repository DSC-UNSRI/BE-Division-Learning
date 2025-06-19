package repository

import (
	"database/sql"
	"github.com/artichys/uts-raki/models" 
)

type AnswerRepository struct {
	db *sql.DB
}

func NewAnswerRepository(db *sql.DB) *AnswerRepository {
	return &AnswerRepository{db: db}
}


func (r *AnswerRepository) CreateAnswer(answer *models.Answer) error {
	query := `INSERT INTO answers (question_id, user_id, content) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, answer.QuestionID, answer.UserID, answer.Content)
	return err
}

func (r *AnswerRepository) GetAnswersByQuestionID(questionID int) ([]models.Answer, error) {
	query := `
		SELECT a.id, a.question_id, a.user_id, a.content, a.created_at, a.updated_at, u.username
		FROM answers a
		JOIN users u ON a.user_id = u.id
		WHERE a.question_id = ?
		ORDER BY a.created_at ASC
	`
	rows, err := r.db.Query(query, questionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var answers []models.Answer
	for rows.Next() {
		var a models.Answer
		err := rows.Scan(
			&a.ID, &a.QuestionID, &a.UserID, &a.Content,
			&a.CreatedAt, &a.UpdatedAt, &a.Username,
		)
		if err != nil {
			return nil, err
		}
		answers = append(answers, a)
	}
	return answers, nil
}

func (r *AnswerRepository) UpdateAnswer(answer *models.Answer) error {
	query := `UPDATE answers SET content = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ? AND user_id = ? AND question_id = ?`
	result, err := r.db.Exec(query, answer.Content, answer.ID, answer.UserID, answer.QuestionID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows 
	}
	return nil
}

func (r *AnswerRepository) DeleteAnswer(answerID, questionID, userID int) error {
	query := `DELETE FROM answers WHERE id = ? AND question_id = ? AND user_id = ?`
	result, err := r.db.Exec(query, answerID, questionID, userID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows 
	}
	return nil
}

func (r *AnswerRepository) GetAnswerByID(answerID int) (*models.Answer, error) {
	var a models.Answer
	query := `SELECT id, question_id, user_id, content, created_at, updated_at FROM answers WHERE id = ?`
	err := r.db.QueryRow(query, answerID).Scan(
		&a.ID, &a.QuestionID, &a.UserID, &a.Content, &a.CreatedAt, &a.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil 
	}
	return &a, err
}