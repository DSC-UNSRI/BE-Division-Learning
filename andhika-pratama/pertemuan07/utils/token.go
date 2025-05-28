
package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateToken(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
