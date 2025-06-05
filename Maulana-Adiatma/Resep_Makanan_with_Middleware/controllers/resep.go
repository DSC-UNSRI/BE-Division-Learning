package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "resepku/database"
    "resepku/models"
)

func GetAllResep(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT * FROM data_resep")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var reseps []models.Resep
    for rows.Next() {
        var r models.Resep
        if err := rows.Scan(&r.ID, &r.NamaResep, &r.DeskripsiResep, &r.BahanUtama, &r.WaktuMasak, &r.NegaraID); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        reseps = append(reseps, r)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(reseps)
}

func GetResepByID(w http.ResponseWriter, r *http.Request) {
    idParam := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var rcp models.Resep
    err = database.DB.QueryRow("SELECT * FROM data_resep WHERE id = ?", id).
        Scan(&rcp.ID, &rcp.NamaResep, &rcp.DeskripsiResep, &rcp.BahanUtama, &rcp.WaktuMasak, &rcp.NegaraID)

    if err != nil {
        http.Error(w, "Resep not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(rcp)
}

func CreateResep(w http.ResponseWriter, r *http.Request) {
    var rcp models.Resep
    err := json.NewDecoder(r.Body).Decode(&rcp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err = database.DB.Exec("INSERT INTO data_resep(nama_resep, deskripsi_resep, bahan_utama, waktu_masak, negara_id) VALUES(?, ?, ?, ?, ?)",
        rcp.NamaResep, rcp.DeskripsiResep, rcp.BahanUtama, rcp.WaktuMasak, rcp.NegaraID)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Resep created"))
}

func UpdateResep(w http.ResponseWriter, r *http.Request) {
    idParam := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var rcp models.Resep
    err = json.NewDecoder(r.Body).Decode(&rcp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err = database.DB.Exec("UPDATE data_resep SET nama_resep=?, deskripsi_resep=?, bahan_utama=?, waktu_masak=?, negara_id=? WHERE id=?",
        rcp.NamaResep, rcp.DeskripsiResep, rcp.BahanUtama, rcp.WaktuMasak, rcp.NegaraID, id)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Resep updated"))
}

func DeleteResep(w http.ResponseWriter, r *http.Request) {
    idParam := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    _, err = database.DB.Exec("DELETE FROM data_resep WHERE id=?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Resep deleted"))
}