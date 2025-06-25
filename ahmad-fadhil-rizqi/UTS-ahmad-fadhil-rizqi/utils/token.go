package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateBearerToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}