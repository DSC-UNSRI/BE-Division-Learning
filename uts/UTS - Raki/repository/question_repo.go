package repository

import (
	"database/sql"
	"github.com/artichys/uts-raki/models" 
)

type QuestionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) CreateQuestion(question *models.Question) error {
	query := `INSERT INTO questions (user_id, title, content, is_promoted) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, question.UserID, question.Title, question.Content, question.IsPromoted)
	return err
}

func (r *QuestionRepository) GetAllQuestions() ([]models.Question, error) {
	query := `
		SELECT q.id, q.user_id, q.title, q.content, q.created_at, q.updated_at, u.username, q.is_promoted
		FROM questions q
		JOIN users u ON q.user_id = u.id
		ORDER BY q.is_promoted DESC, q.created_at DESC -- Prioritize promoted questions
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var q models.Question
		err := rows.Scan(
			&q.ID, &q.UserID, &q.Title, &q.Content,
			&q.CreatedAt, &q.UpdatedAt, &q.Username, &q.IsPromoted, 
		)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}


func (r *QuestionRepository) GetQuestionByID(id int) (*models.Question, error) {
	var q models.Question
	query := `
		SELECT q.id, q.user_id, q.title, q.content, q.created_at, q.updated_at, u.username, q.is_promoted
		FROM questions q
		JOIN users u ON q.user_id = u.id
		WHERE q.id = ?
	`
	err := r.db.QueryRow(query, id).Scan(
		&q.ID, &q.UserID, &q.Title, &q.Content,
		&q.CreatedAt, &q.UpdatedAt, &q.Username, &q.IsPromoted, 
	)
	if err == sql.ErrNoRows {
		return nil, nil 
	}
	return &q, err
}

func (r *QuestionRepository) GetUserQuestions(userID int) ([]models.Question, error) {
	query := `
		SELECT q.id, q.user_id, q.title, q.content, q.created_at, q.updated_at, u.username, q.is_promoted
		FROM questions q
		JOIN users u ON q.user_id = u.id
		WHERE q.user_id = ?
		ORDER BY q.created_at DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var q models.Question
		err := rows.Scan(
			&q.ID, &q.UserID, &q.Title, &q.Content,
			&q.CreatedAt, &q.UpdatedAt, &q.Username, &q.IsPromoted, 
		)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func (r *QuestionRepository) UpdateQuestion(question *models.Question) error {
	query := `UPDATE questions SET title = ?, content = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ? AND user_id = ?`
	result, err := r.db.Exec(query, question.Title, question.Content, question.ID, question.UserID)
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

func (r *QuestionRepository) DeleteQuestion(questionID, userID int) error {
	query := `DELETE FROM questions WHERE id = ? AND user_id = ?`
	result, err := r.db.Exec(query, questionID, userID)
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

func (r *QuestionRepository) PromoteQuestion(questionID int, userID int) error {
	query := `UPDATE questions SET is_promoted = TRUE, updated_at = CURRENT_TIMESTAMP WHERE id = ? AND user_id = ?`
	result, err := r.db.Exec(query, questionID, userID)
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