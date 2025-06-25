package middleware

import (
	"UTS-Ahmad-Fadhil-Rizqi/database"
	"UTS-Ahmad-Fadhil-Rizqi/utils"
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: Token tidak ditemukan", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Unauthorized: Format token tidak valid", http.StatusUnauthorized)
			return
		}

		var userID int64
		var userTier string
		
		err := database.DB.QueryRow(
			"SELECT id, tier FROM users WHERE token = ? AND deleted_at IS NULL AND token_expires_at > NOW()",
			tokenString,
		).Scan(&userID, &userTier)

		if err != nil {
			http.Error(w, "Unauthorized: Token tidak valid atau sudah kadaluwarsa", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), utils.UserIDKey, userID)
		ctx = context.WithValue(ctx, utils.UserTierKey, userTier)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}