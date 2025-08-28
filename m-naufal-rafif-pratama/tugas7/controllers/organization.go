package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas7/database"
	"tugas7/models"
)

func GetAllOrganizations(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM organizations")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orgs []models.Organization
	for rows.Next() {
		var org models.Organization
		if err := rows.Scan(&org.ID, &org.Name, &org.Type); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orgs = append(orgs, org)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orgs)
}

func GetOrganizationByID(w http.ResponseWriter, r *http.Request, id string) {
	var org models.Organization
	err := database.DB.QueryRow("SELECT * FROM organizations WHERE id = ?", id).Scan(
		&org.ID, &org.Name, &org.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Organization not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(org)
}

func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var org models.Organization
	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !org.IsValid() {
		http.Error(w, "Invalid organization data", http.StatusBadRequest)
		return
	}

	result, err := database.DB.Exec("INSERT INTO organizations (name, type) VALUES (?, ?)",
		org.Name, org.Type)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	org.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(org)
}

func UpdateOrganization(w http.ResponseWriter, r *http.Request, id string) {
	var org models.Organization
	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec("UPDATE organizations SET name = ?, type = ? WHERE id = ?",
		org.Name, org.Type, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(org)
}

func DeleteOrganization(w http.ResponseWriter, r *http.Request, id string) {
	result, err := database.DB.Exec("DELETE FROM organizations WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Organization not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetOrganizationMembers(w http.ResponseWriter, r *http.Request, orgID string) {
	rows, err := database.DB.Query(`
		SELECT s.id, s.name, s.email, s.major, s.year 
		FROM students s
		WHERE s.org_id = ?
	`, orgID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var members []models.Student
	for rows.Next() {
		var member models.Student
		if err := rows.Scan(&member.ID, &member.Name, &member.Email, &member.Major, &member.Year); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		members = append(members, member)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}
