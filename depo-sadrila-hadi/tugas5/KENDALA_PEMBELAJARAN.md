# Kendala, Pembelajaran, dan Desain Sistem PMM

Dokumen ini menjelaskan beberapa kendala yang dihadapi selama pengembangan, pembelajaran yang didapat, serta alasan di balik beberapa pilihan desain dalam sistem Pendataan Minat Mahasiswa (PMM).

## 1. Kendala yang Dihadapi

* **Routing Kompleks dengan `net/http`:** Menggunakan `http.ServeMux` bawaan Go untuk routing, terutama untuk path dengan parameter ganda seperti `/mahasiswa/{id}/minat/{minat_id}`, memerlukan parsing path manual di dalam handler (seperti di `routes/routes.go`). Ini sedikit lebih rumit dibandingkan menggunakan library router pihak ketiga (misalnya Gorilla Mux, Gin) yang menyediakan fitur parsing parameter secara otomatis.
* **Penanganan Nilai NULL dari Database:** Saat mengambil data dari database, kolom yang memperbolehkan NULL (seperti `deskripsi` pada tabel `minat` atau `deleted_at` pada semua tabel) memerlukan penanganan khusus di Go. Menggunakan tipe data seperti `sql.NullString` atau pointer (`*time.Time`) diperlukan agar proses `Scan` tidak error jika nilai di database adalah NULL.
* **Manajemen Error Detail:** Membedakan antara error `sql.ErrNoRows` (data tidak ditemukan) dengan error database lainnya penting untuk memberikan response HTTP yang tepat (misalnya `404 Not Found` vs `500 Internal Server Error`). Ini memerlukan pengecekan error secara eksplisit setelah query database.
* **Logika Autentikasi & Otorisasi:** Meskipun sederhana, implementasi Basic Auth dan pengecekan kepemilikan data memerlukan beberapa langkah: membaca header, decoding base64, query ke database, perbandingan data, dan penentuan response code yang sesuai (`401 Unauthorized`, `403 Forbidden`). Memastikan alur ini benar dan aman (walaupun password plaintext) butuh ketelitian.
* **Struktur Kode Awal:** Mungkin pada awalnya ada kesulitan dalam menentukan struktur folder yang baik sebelum akhirnya memutuskan pola `models`, `controllers`, `routes`, `utils`, dll.

## 2. Ilmu yang Dipelajari

* **Dasar Pengembangan Web Backend dengan Go:** Penggunaan standard library `net/http` untuk membuat server, menangani request, routing dasar, dan mengirim response.
* **Interaksi Database (MySQL) di Go:** Menggunakan package `database/sql` dan driver `go-sql-driver/mysql` untuk koneksi, eksekusi query (SELECT, INSERT, UPDATE, DELETE), dan memproses hasil query (termasuk handling NULL dan error).
* **Desain API RESTful Sederhana:** Merancang endpoint API yang logis untuk operasi CRUD dan relasi data, serta menggunakan metode HTTP yang sesuai (GET, POST, PATCH, DELETE).
* **Pentingnya Struktur Proyek Modular:** Memahami manfaat memisahkan kode berdasarkan fungsinya (models, controllers, routes, utils) untuk keterbacaan, pemeliharaan, dan skalabilitas.
* **Manajemen Konfigurasi:** Menggunakan file `.env` untuk menyimpan konfigurasi sensitif (seperti kredensial DB) dan memisahkannya dari kode sumber.
* **Implementasi Autentikasi Dasar:** Memahami cara kerja HTTP Basic Authentication dan alur validasinya di sisi server.
* **Penanganan JSON:** Proses encoding data Go ke JSON untuk response API dan decoding JSON dari request body ke struct Go.
* **Konsep Soft Delete:** Memahami tujuan dan cara implementasi soft delete menggunakan kolom timestamp `deleted_at` untuk menjaga histori data.

## 3. Alasan Desain Sistem

* **Penggunaan `net/http` Standard Library:** Dipilih untuk memenuhi persyaratan tugas dan untuk memahami dasar-dasar web di Go sebelum beralih ke framework yang lebih kompleks. Ini memberikan pemahaman yang lebih baik tentang apa yang terjadi di balik layar framework.
* **Struktur Modular:** Mengadopsi struktur folder yang umum digunakan dalam pengembangan aplikasi Go untuk mempromosikan kode yang bersih, terorganisir, dan mudah dikelola.
* **Autentikasi Basic (Plaintext):** Dipilih secara spesifik karena instruksi tugas meminta "auth simpel" dengan pengecekan variabel (nama/password) ke database. Disadari bahwa ini **bukan praktik yang aman** untuk aplikasi produksi, di mana hashing password wajib digunakan.
* **Soft Delete:** Diimplementasikan untuk menghindari kehilangan data permanen. Data yang "dihapus" hanya ditandai dan dapat dipulihkan atau dianalisis jika perlu di kemudian hari.
* **Fungsi Utilitas (`utils`):** Dibuat untuk menerapkan prinsip DRY (Don't Repeat Yourself). Fungsi seperti `RespondWithJSON`, `RespondWithError`, dan `CheckAuthAndRespond` digunakan berulang kali di banyak controller, sehingga memisahkannya ke package `utils` membuat kode lebih bersih.
* **Migrasi Skema Sederhana:** Proses migrasi (`CREATE TABLE IF NOT EXISTS`) yang dijalankan saat startup sudah cukup untuk skala proyek tugas ini. Untuk proyek yang lebih besar, library migrasi yang lebih canggih mungkin diperlukan.