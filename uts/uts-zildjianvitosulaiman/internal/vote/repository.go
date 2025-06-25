package vote

import (
	"database/sql"
	"uts-zildjianvitosulaiman/domain"
)

type Repository interface {
	FindByUserAndAnswer(userID, answerID int) (*domain.Vote, error)
	Upsert(vote *domain.Vote) error
	Delete(userID, answerID int) error
	GetVoteCounts(answerID int) (upvotes, downvotes int, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByUserAndAnswer(userID, answerID int) (*domain.Vote, error) {
	var v domain.Vote
	query := `SELECT id, user_id, answer_id, vote_type FROM votes WHERE user_id = ? AND answer_id = ?`
	err := r.db.QueryRow(query, userID, answerID).Scan(&v.ID, &v.UserID, &v.AnswerID, &v.Type)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *repository) Upsert(v *domain.Vote) error {
	query := `INSERT INTO votes (user_id, answer_id, vote_type) VALUES (?, ?, ?)
	          ON DUPLICATE KEY UPDATE vote_type = VALUES(vote_type)`
	_, err := r.db.Exec(query, v.UserID, v.AnswerID, v.Type)
	return err
}

func (r *repository) Delete(userID, answerID int) error {
	query := `DELETE FROM votes WHERE user_id = ? AND answer_id = ?`
	_, err := r.db.Exec(query, userID, answerID)
	return err
}

func (r *repository) GetVoteCounts(answerID int) (upvotes, downvotes int, err error) {
	query := `SELECT
	    COALESCE(SUM(CASE WHEN vote_type = 1 THEN 1 ELSE 0 END), 0) as upvotes,
	    COALESCE(SUM(CASE WHEN vote_type = -1 THEN 1 ELSE 0 END), 0) as downvotes
	    FROM votes WHERE answer_id = ?`
	err = r.db.QueryRow(query, answerID).Scan(&upvotes, &downvotes)
	return
}
