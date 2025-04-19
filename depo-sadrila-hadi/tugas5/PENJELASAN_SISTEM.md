# Penjelasan Sistem Pendataan Minat Mahasiswa (PMM)

## 1. Tujuan Sistem

Sistem ini bertujuan untuk mengelola data mahasiswa beserta minat yang mereka miliki. Fungsi utamanya meliputi:
* Registrasi dan pengelolaan data mahasiswa (CRUD).
* Pengelolaan data minat yang tersedia (CRUD).
* Pencatatan dan pengelolaan relasi antara mahasiswa dan minat yang mereka pilih.
* Menyediakan API backend sederhana untuk operasi-operasi tersebut.

## 2. Teknologi yang Digunakan

* **Bahasa Pemrograman:** Go (Golang)
* **Web Server & Routing:** Standard library Go (`net/http`, `http.ServeMux`)
* **Database:** MySQL
* **Driver Database Go:** `github.com/go-sql-driver/mysql`
* **Manajemen Konfigurasi:** File `.env` dan library `github.com/joho/godotenv`

## 3. Struktur Proyek

Proyek ini diorganisir ke dalam beberapa package untuk memisahkan tanggung jawab (separation of concerns):

* **`main.go`**: Entry point aplikasi, menginisialisasi konfigurasi, database, migrasi, dan router.
* **`/config`**: Mengelola pemuatan variabel lingkungan dari file `.env`.
* **`/database`**: Menangani koneksi ke database MySQL dan menjalankan migrasi skema awal.
* **`/models`**: Mendefinisikan struct Go yang merepresentasikan tabel database (`Mahasiswa`, `Minat`, `MahasiswaMinat`) beserta query SQL untuk pembuatan tabel.
* **`/controllers`**: Berisi logika bisnis untuk menangani request HTTP, berinteraksi dengan database, dan mempersiapkan response. Terdapat controller untuk `Mahasiswa`, `Minat`, dan relasi `MahasiswaMinat`.
* **`/routes`**: Mendefinisikan endpoint API, memetakan path URL dan metode HTTP ke fungsi controller yang sesuai. Menggunakan `http.ServeMux` bawaan Go.
* **`/utils`**: Berisi fungsi-fungsi pembantu (helper) seperti response JSON, response error, konversi tipe data, dan logika autentikasi/otorisasi.

## 4. Model Data

Sistem ini menggunakan 3 tabel utama di database:

1.  **`mahasiswa`**: Menyimpan data pengguna sistem.
    * `id` (INT, PK, AI): ID unik mahasiswa.
    * `nama` (VARCHAR, UNIQUE): Nama mahasiswa, digunakan untuk login.
    * `password` (VARCHAR): Password mahasiswa (disimpan sebagai plaintext untuk tugas ini).
    * `created_at`, `updated_at`, `deleted_at` (TIMESTAMP): Metadata waktu dan untuk soft delete.
2.  **`minat`**: Menyimpan daftar minat yang bisa dipilih.
    * `id` (INT, PK, AI): ID unik minat.
    * `nama_minat` (VARCHAR, UNIQUE): Nama minat.
    * `deskripsi` (TEXT, NULL): Deskripsi opsional tentang minat.
    * `created_at`, `updated_at`, `deleted_at` (TIMESTAMP): Metadata waktu dan untuk soft delete.
3.  **`mahasiswa_minat`**: Tabel junction (many-to-many) untuk menghubungkan mahasiswa dengan minat yang dipilih.
    * `mahasiswa_id` (INT, FK): Merujuk ke `mahasiswa.id`.
    * `minat_id` (INT, FK): Merujuk ke `minat.id`.
    * Primary Key komposit (`mahasiswa_id`, `minat_id`).

## 5. Fitur Utama / Endpoints API

### 5.1. Mahasiswa

* **`POST /mahasiswa`**: Registrasi mahasiswa baru. (Tidak perlu Autentikasi)
    * Request Body: `{"nama": "...", "password": "..."}`
    * Response: Data mahasiswa yang baru dibuat (tanpa password).
* **`GET /mahasiswa`**: Mendapatkan daftar semua mahasiswa (yang belum dihapus). (Autentikasi Diperlukan)
    * Response: Array objek mahasiswa.
* **`GET /mahasiswa/{id}`**: Mendapatkan detail mahasiswa berdasarkan ID. (Autentikasi Diperlukan, Otorisasi: Sebaiknya hanya user sendiri)
    * Response: Objek mahasiswa tunggal.
* **`PATCH /mahasiswa/{id}`**: Memperbarui data mahasiswa (nama dan/atau password). (Autentikasi & Otorisasi Diperlukan: Hanya user pemilik ID)
    * Request Body: `{"nama": "...", "password": "..."}` (field opsional)
    * Response: Data mahasiswa yang sudah diperbarui (tanpa password).
* **`DELETE /mahasiswa/{id}`**: Melakukan soft delete pada mahasiswa. (Autentikasi & Otorisasi Diperlukan: Hanya user pemilik ID)
    * Response: Pesan sukses.

### 5.2. Minat

* **`POST /minat`**: Menambahkan minat baru. (Autentikasi Diperlukan)
    * Request Body: `{"nama_minat": "...", "deskripsi": "..."}` (`deskripsi` opsional)
    * Response: Data minat yang baru dibuat.
* **`GET /minat`**: Mendapatkan daftar semua minat (yang belum dihapus). (Autentikasi Diperlukan)
    * Response: Array objek minat.
* **`GET /minat/{id}`**: Mendapatkan detail minat berdasarkan ID. (Autentikasi Diperlukan)
    * Response: Objek minat tunggal.
* **`PATCH /minat/{id}`**: Memperbarui data minat. (Autentikasi Diperlukan)
    * Request Body: `{"nama_minat": "...", "deskripsi": "..."}` (field opsional)
    * Response: Data minat yang sudah diperbarui.
* **`DELETE /minat/{id}`**: Melakukan soft delete pada minat. (Autentikasi Diperlukan)
    * Response: Pesan sukses.

### 5.3. Relasi Mahasiswa-Minat

* **`POST /mahasiswa/{id}/minat`**: Menambahkan relasi minat ke mahasiswa tertentu. (Autentikasi & Otorisasi Diperlukan: Hanya user pemilik ID Mahasiswa)
    * Request Body: `{"minat_id": ...}`
    * Response: Pesan sukses.
* **`GET /mahasiswa/{id}/minat`**: Mendapatkan daftar minat yang dimiliki oleh mahasiswa tertentu. (Autentikasi Diperlukan)
    * Response: Array objek minat milik mahasiswa tersebut.
* **`DELETE /mahasiswa/{id}/minat/{minat_id}`**: Menghapus relasi minat dari mahasiswa tertentu. (Autentikasi & Otorisasi Diperlukan: Hanya user pemilik ID Mahasiswa)
    * Response: Pesan sukses.

## 6. Autentikasi & Otorisasi

* **Autentikasi:** Sistem menggunakan HTTP Basic Authentication. Klien harus mengirimkan header `Authorization: Basic <base64(nama:password)>`. Backend akan mendekode kredensial ini dan membandingkannya dengan data di tabel `mahasiswa` (nama dan password plaintext).
* **Otorisasi:** Implementasi otorisasi sederhana berbasis kepemilikan. Hanya mahasiswa yang terautentikasi yang dapat mengubah atau menghapus profilnya sendiri, serta menambah atau menghapus minat dari profilnya. Endpoint untuk mengelola data `Minat` (CRUD) hanya memerlukan autentikasi (semua user terautentikasi bisa melakukannya).

## 7. Database & Migrasi

* Menggunakan database MySQL. Konfigurasi koneksi (host, port, user, password, dbname) diambil dari file `.env`.
* Migrasi skema database dilakukan secara otomatis saat aplikasi pertama kali dijalankan melalui fungsi `database.Migrate()` yang mengeksekusi query `CREATE TABLE IF NOT EXISTS` dari package `models`.

## 8. Cara Menjalankan

1.  Pastikan Go dan MySQL sudah terinstall.
2.  Clone repository proyek.
3.  Buat file `.env` di root proyek dan isi variabel berikut sesuai konfigurasi database Anda:
    ```dotenv
    DB_USER=root
    DB_PASSWORD=
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME=db_pmm
    PORT=8080
    ```
4.  Buat database di MySQL sesuai dengan `DB_NAME`.
5.  Jalankan aplikasi: `go run main.go`
6.  Server akan berjalan di `http://localhost:8080` (atau port yang didefinisikan di `.env`).