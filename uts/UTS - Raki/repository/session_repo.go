
package repository

import (
	"database/sql"

	"github.com/artichys/uts-raki/models"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) CreateSession(session *models.Session) error {
	query := `INSERT INTO sessions (token, user_id, user_type, expires_at) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, session.Token, session.UserID, session.UserType, session.ExpiresAt)
	return err
}

func (r *SessionRepository) GetSessionByToken(token string) (*models.Session, error) {
	session := &models.Session{}
	query := `SELECT token, user_id, user_type, expires_at, created_at FROM sessions WHERE token = ?`
	err := r.db.QueryRow(query, token).Scan(
		&session.Token,
		&session.UserID,
		&session.UserType,
		&session.ExpiresAt,
		&session.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return session, err
}

func (r *SessionRepository) DeleteSessionByToken(token string) error {
	query := `DELETE FROM sessions WHERE token = ?`
	result, err := r.db.Exec(query, token)
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