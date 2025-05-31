package middleware

import (
	"be_pert7/database"
	"be_pert7/utils"

	"database/sql"
	"context"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Unauthorized - no token", http.StatusUnauthorized)
			return
		}

		const prefix = "Bearer "
		if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
			http.Error(w, "Unauthorized - invalid token format", http.StatusUnauthorized)
			return
		}

		token := authHeader[len(prefix):]

		var exists bool
		var chef_id string
		var chef_role string

		err := database.DB.QueryRow(
			"SELECT EXISTS(SELECT 1 FROM chefs WHERE token = ? AND deleted_at IS NULL), chef_id, role FROM chefs WHERE token = ? AND deleted_at IS NULL",
			token, token,
		).Scan(&exists, &chef_id, &chef_role)
		if err != nil || !exists {
			http.Error(w, "Unauthorized - invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), utils.ChefIDKey, chef_id)
		ctx = context.WithValue(ctx, utils.RoleKey, chef_role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func HeadMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value(utils.RoleKey)
		if role == nil || role.(string) != "head" {
			http.Error(w, "Forbidden - Head chef access only hahahah", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CourseOwnershipMiddleware(next http.Handler, MenuID string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxChefID := r.Context().Value(utils.ChefIDKey)
		ChefID, ok := ctxChefID.(string)
		if !ok || ChefID == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if MenuID == "" {
			http.Error(w, "menu_id is null", http.StatusBadRequest)
			return
		}

		var courseOwnerID string
		err := database.DB.QueryRow("SELECT chef_id FROM menus WHERE menu_id = ? AND deleted_at IS NULL", MenuID).Scan(&courseOwnerID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Chef not found", http.StatusNotFound)
			} else {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}

		if courseOwnerID != ChefID {
			http.Error(w, "Forbidden - the menu isn't yours", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
