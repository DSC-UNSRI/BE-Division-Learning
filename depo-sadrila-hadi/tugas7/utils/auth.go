package utils

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"tugas7/database"
	"tugas7/models"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUserFromRequest(r *http.Request) (models.Mahasiswa, error) {
	var authenticatedUser models.Mahasiswa
	var dbHashedPassword string

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return authenticatedUser, errors.New("authorization header required")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Basic") {
		return authenticatedUser, errors.New("invalid authorization format, expected Basic auth")
	}

	payload, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return authenticatedUser, errors.New("invalid base64 encoding")
	}

	credentials := strings.SplitN(string(payload), ":", 2)
	if len(credentials) != 2 {
		return authenticatedUser, errors.New("invalid credentials format, expected nama:password")
	}

	nama := credentials[0]
	providedPassword := credentials[1]

	query := "SELECT id, nama, password FROM mahasiswa WHERE nama = ? AND deleted_at IS NULL"
	err = database.DB.QueryRow(query, nama).Scan(&authenticatedUser.ID, &authenticatedUser.Nama, &dbHashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return authenticatedUser, errors.New("invalid credentials")
		}
		fmt.Printf("Database error during authentication for user %s: %v\n", nama, err)
		return authenticatedUser, errors.New("internal server error during authentication")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbHashedPassword), []byte(providedPassword))
	if err != nil {
		return authenticatedUser, errors.New("invalid credentials")
	}

	authenticatedUser.Password = ""
	return authenticatedUser, nil
}