package routes

import (
	"net/http"
	"strings"

	"pmm/controllers"
	"pmm/utils"
)

func InitRoutes() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /mahasiswa", controllers.CreateMahasiswa)
	mux.HandleFunc("GET /mahasiswa", controllers.GetAllMahasiswa)

	mux.HandleFunc("/mahasiswa/", mahasiswaSubHandler)

	mux.HandleFunc("/minat", minatBaseHandler)
	mux.HandleFunc("/minat/", minatIDHandler)

	http.Handle("/", mux)
}

func mahasiswaSubHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	if len(parts) < 2 || parts[0] != "mahasiswa" {
		http.NotFound(w, r)
		return
	}

	mahasiswaIDStr := parts[1]

	switch len(parts) {
	case 2:
		handleMahasiswaIDEndpoint(w, r, mahasiswaIDStr)
	case 3:
		if parts[2] == "minat" {
			handleMahasiswaMinatBaseEndpoint(w, r, mahasiswaIDStr)
		} else {
			http.NotFound(w, r)
		}
	case 4:
		if parts[2] == "minat" {
			minatIDStr := parts[3]
			handleMahasiswaMinatSpecificEndpoint(w, r, mahasiswaIDStr, minatIDStr)
		} else {
			http.NotFound(w, r)
		}
	default:
		http.NotFound(w, r)
	}
}

func handleMahasiswaIDEndpoint(w http.ResponseWriter, r *http.Request, mahasiswaIDStr string) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetMahasiswaByID(w, r, mahasiswaIDStr)
	case http.MethodPatch:
		controllers.UpdateMahasiswa(w, r, mahasiswaIDStr)
	case http.MethodDelete:
		controllers.DeleteMahasiswa(w, r, mahasiswaIDStr)
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed for /mahasiswa/{id}")
	}
}

func handleMahasiswaMinatBaseEndpoint(w http.ResponseWriter, r *http.Request, mahasiswaIDStr string) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetMinatByMahasiswaID(w, r, mahasiswaIDStr)
	case http.MethodPost:
		controllers.AddMinatToMahasiswa(w, r, mahasiswaIDStr)
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed for /mahasiswa/{id}/minat")
	}
}

func handleMahasiswaMinatSpecificEndpoint(w http.ResponseWriter, r *http.Request, mahasiswaIDStr string, minatIDStr string) {
	switch r.Method {
	case http.MethodDelete:
		controllers.RemoveMinatFromMahasiswa(w, r, mahasiswaIDStr, minatIDStr)
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed for /mahasiswa/{id}/minat/{minat_id}")
	}
}

func minatBaseHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/minat" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		controllers.GetAllMinat(w, r)
	case http.MethodPost:
		controllers.CreateMinat(w, r)
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed for /minat")
	}
}

func minatIDHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	if len(parts) != 2 || parts[0] != "minat" {
		http.NotFound(w, r)
		return
	}
	idStr := parts[1]

	switch r.Method {
	case http.MethodGet:
		controllers.GetMinatByID(w, r, idStr)
	case http.MethodPatch:
		controllers.UpdateMinat(w, r, idStr)
	case http.MethodDelete:
		controllers.DeleteMinat(w, r, idStr)
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed for /minat/{id}")
	}
}