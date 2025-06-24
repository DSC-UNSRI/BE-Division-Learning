package middleware

import (
	"log"
	"net/http"
	"strconv"
)

func OrgAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgIDStr := r.Header.Get("X-Student-OrgID")
		if orgIDStr == "" {
			log.Println("OrgAuth Middleware: Organization ID not found in request")
			http.Error(w, "Organization ID not found", http.StatusUnauthorized)
			return
		}

		orgID, err := strconv.Atoi(orgIDStr)
		if err != nil {
			log.Printf("OrgAuth Middleware: Invalid organization ID format: %v", err)
			http.Error(w, "Invalid organization ID format", http.StatusBadRequest)
			return
		}

		log.Printf("OrgAuth Middleware: Found organization ID: %d", orgID)

		if queryOrgID := r.URL.Query().Get("org_id"); queryOrgID != "" {
			requestOrgID, err := strconv.Atoi(queryOrgID)
			if err != nil {
				log.Printf("OrgAuth Middleware: Invalid requested organization ID format: %v", err)
				http.Error(w, "Invalid requested organization ID format", http.StatusBadRequest)
				return
			}

			if orgID != requestOrgID {
				log.Printf("OrgAuth Middleware: Unauthorized access attempt. Token OrgID: %d, Requested OrgID: %d", 
					orgID, requestOrgID)
				http.Error(w, "Unauthorized: You can only access data from your own organization", http.StatusForbidden)
				return
			}

			log.Printf("OrgAuth Middleware: Access granted to organization %d", requestOrgID)
		}

		next.ServeHTTP(w, r)
	}
} 