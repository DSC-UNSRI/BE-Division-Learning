package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "resepku/database"
    "resepku/models"
)

func GetAllNegara(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT * FROM data_negara")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var negaras []models.Negara
    for rows.Next() {
        var r models.Negara
        if err := rows.Scan(&r.ID, &r.NamaNegara, &r.KodeNegara); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        negaras = append(negaras, r)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(negaras)
}

func GetNegaraByID(w http.ResponseWriter, r *http.Request) {
    idParam := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var rcp models.Negara
    err = database.DB.QueryRow("SELECT * FROM data_negara WHERE id = ?", id).
        Scan(&rcp.ID, &rcp.NamaNegara, &rcp.KodeNegara)

    if err != nil {
        http.Error(w, "Negara not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(rcp)
}

func CreateNegara(w http.ResponseWriter, r *http.Request) {
    var rcp models.Negara
    err := json.NewDecoder(r.Body).Decode(&rcp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err = database.DB.Exec("INSERT INTO data_negara (negara_asal, kode_negara) VALUES (?, ?)",
        rcp.NamaNegara, rcp.KodeNegara)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Negara created"))
}

func UpdateNegara(w http.ResponseWriter, r *http.Request) {
    idParam := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var rcp models.Negara
    err = json.NewDecoder(r.Body).Decode(&rcp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err = database.DB.Exec("UPDATE data_negara SET negara_asal=?, kode_negara=? WHERE id=?",
        rcp.NamaNegara, rcp.KodeNegara, id)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Negara updated"))
}

func DeleteNegara(w http.ResponseWriter, r *http.Request) {
    idParam := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    _, err = database.DB.Exec("DELETE FROM data_negara WHERE id=?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Negara deleted"))
}
