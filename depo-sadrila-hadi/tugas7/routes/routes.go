package routes

import (
	"net/http"
	"tugas7/controllers"
	"tugas7/middleware"
	"tugas7/utils"
	"strings"
)

func applyMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func InitRoutes() {
	mux := http.NewServeMux()

	loggingOnlyStack := func(h http.HandlerFunc) http.Handler {
		return applyMiddlewares(h, middleware.LoggingMiddleware)
	}

	authAndLoggingStack := func(h http.HandlerFunc) http.Handler {
		return applyMiddlewares(h, middleware.LoggingMiddleware, middleware.AuthMiddleware)
	}

	mux.Handle("POST /mahasiswa", loggingOnlyStack(http.HandlerFunc(controllers.CreateMahasiswa)))
	mux.Handle("POST /login", loggingOnlyStack(http.HandlerFunc(controllers.LoginMahasiswa)))

	mux.Handle("GET /mahasiswa", authAndLoggingStack(http.HandlerFunc(controllers.GetAllMahasiswa)))

	mux.HandleFunc("/mahasiswa/", func(w http.ResponseWriter, r *http.Request) {
		var h http.HandlerFunc
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		if len(parts) < 2 || parts[0] != "mahasiswa" {
			h = http.NotFound
		} else {
			mahasiswaIDStr := parts[1]
			switch len(parts) {
			case 2:
				h = func(writer http.ResponseWriter, request *http.Request) {
					handleMahasiswaIDEndpoint(writer, request, mahasiswaIDStr)
				}
			case 3:
				if parts[2] == "minat" {
					h = func(writer http.ResponseWriter, request *http.Request) {
						handleMahasiswaMinatBaseEndpoint(writer, request, mahasiswaIDStr)
					}
				} else {
					h = http.NotFound
				}
			case 4:
				if parts[2] == "minat" {
					minatIDStr := parts[3]
					h = func(writer http.ResponseWriter, request *http.Request) {
						handleMahasiswaMinatSpecificEndpoint(writer, request, mahasiswaIDStr, minatIDStr)
					}
				} else {
					h = http.NotFound
				}
			default:
				h = http.NotFound
			}
		}
		authAndLoggingStack(h).ServeHTTP(w, r)
	})

	mux.HandleFunc("/minat", func(w http.ResponseWriter, r *http.Request) {
		var h http.HandlerFunc
		if r.URL.Path != "/minat" {
			h = http.NotFound
		} else {
			h = func(writer http.ResponseWriter, request *http.Request) {
				switch request.Method {
				case http.MethodGet:
					controllers.GetAllMinat(writer, request)
				case http.MethodPost:
					controllers.CreateMinat(writer, request)
				default:
					utils.RespondWithError(writer, http.StatusMethodNotAllowed, "Method not allowed for /minat")
				}
			}
		}
		authAndLoggingStack(h).ServeHTTP(w, r)
	})

	mux.HandleFunc("/minat/", func(w http.ResponseWriter, r *http.Request) {
		var h http.HandlerFunc
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		if len(parts) != 2 || parts[0] != "minat" {
			h = http.NotFound
		} else {
			idStr := parts[1]
			h = func(writer http.ResponseWriter, request *http.Request) {
				switch request.Method {
				case http.MethodGet:
					controllers.GetMinatByID(writer, request, idStr)
				case http.MethodPatch:
					controllers.UpdateMinat(writer, request, idStr)
				case http.MethodDelete:
					controllers.DeleteMinat(writer, request, idStr)
				default:
					utils.RespondWithError(writer, http.StatusMethodNotAllowed, "Method not allowed for /minat/{id}")
				}
			}
		}
		authAndLoggingStack(h).ServeHTTP(w, r)
	})

	http.Handle("/", mux)
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
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed for /mahasiswa/{id} or /mahasiswa/me")
	}
}

func handleMahasiswaMinatBaseEndpoint(w http.ResponseWriter, r *http.Request, mahasiswaIDStr string) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetMinatByMahasiswaID(w, r, mahasiswaIDStr)
	case http.MethodPost:
		controllers.AddMinatToMahasiswa(w, r, mahasiswaIDStr)
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed for /mahasiswa/{id}/minat or /mahasiswa/me/minat")
	}
}

func handleMahasiswaMinatSpecificEndpoint(w http.ResponseWriter, r *http.Request, mahasiswaIDStr string, minatIDStr string) {
	switch r.Method {
	case http.MethodDelete:
		controllers.RemoveMinatFromMahasiswa(w, r, mahasiswaIDStr, minatIDStr)
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed for /mahasiswa/{id}/minat/{minat_id} or /mahasiswa/me/minat/{minat_id}")
	}
}