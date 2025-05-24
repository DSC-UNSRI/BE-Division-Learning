package routes

import (
	"net/http"
	"Task_5/controllers"
)

func RegisterNilaiMataUangRoutes() {
	// Buat data mata uang baru
	http.HandleFunc("/nilai-matauang/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.CreateNilaiMataUang(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Ambil semua data mata uang
	http.HandleFunc("/nilai-matauang/all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetAllNilaiMataUang(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Ambil data berdasarkan ID
	http.HandleFunc("/nilai-matauang/detail", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetNilaiMataUangByID(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Update data mata uang
	http.HandleFunc("/nilai-matauang/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			controllers.UpdateNilaiMataUang(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Hapus data mata uang
	http.HandleFunc("/nilai-matauang/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			controllers.DeleteNilaiMataUang(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
