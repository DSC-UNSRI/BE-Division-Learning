# API CRUD Film & Director

## Deskripsi Sistem

Proyek ini adalah sebuah API yang dibuat menggunakan package standar `net/http`. API ini menyediakan
fungsi dasar CRUD (Create, Read, Update, Delete) untuk mengelola data film dan director, dilengkapi dengan sistem
autentikasi dan otorisasi berbasis peran untuk director.

## Fitur Utama

- **Manajemen Director (CRUD):**

  - Pendaftaran Director Baru (`POST /register/director`) dengan hashing password menggunakan bcrypt.
  - Login Director (`POST /login/director`) untuk mendapatkan token autentikasi.
  - Menampilkan semua Director (`GET /directors`).
  - Menampilkan Director tertentu berdasarkan ID (`GET /directors/{id}`).
  - Membuat data Director baru (khusus admin, `POST /directors`).
  - Memperbarui data Director (`PATCH /directors/{id}`) dengan batasan akses (director hanya bisa mengubah profilnya sendiri kecuali admin).
  - Menghapus data Director (Soft Delete, `DELETE /directors/{id}`) dengan batasan akses.

- **Manajemen Film (CRUD):**

  - Menampilkan semua film (`GET /films`) - Data dipersonalisasi: hanya menampilkan film yang dibuat oleh director yang sedang login.
  - Menampilkan film tertentu berdasarkan ID (`GET /films/{id}`).
  - Membuat data film baru (`POST /films`) - Film otomatis terhubung dengan director yang sedang login.
  - Memperbarui data film (`PATCH /films/{id}`) - Hanya director pembuat atau admin yang dapat mengubah.
  - Menghapus data film (Soft Delete, `DELETE /films/{id}`) - Hanya director pembuat atau admin yang dapat menghapus.

- **Autentikasi & Otorisasi Lanjutan:**
  - **AuthMiddleware:** Memeriksa token di header `Authorization` (Bearer Token) dan mengautentikasi director.
  - **AdminMiddleware:** Membatasi akses hanya untuk director dengan peran 'admin'.
  - **LoggingMiddleware:** Middleware tambahan untuk mencatat setiap permintaan masuk (dapat diaktifkan di konfigurasi rute).
  - **Personalisasi Data:** Pada beberapa endpoint `GET`, data yang ditampilkan disesuaikan dengan director yang sedang login (misalnya `GET /films` dan `GET /getfilmsbydirector/{id}`).

**Cara Menggunakan (Testing API dengan Postman):**

Pastikan aplikasi Go Anda berjalan di `http://localhost:8080/`. Untuk semua permintaan yang memerlukan autentikasi, sertakan header `Authorization: Bearer <your_token_here>`. Ganti `<your_token_here>` dengan token yang Anda dapatkan setelah login.

- **1. Registrasi Director Baru:**

  - Metode: `POST`
  - URL: `http://localhost:8080/register/director`
  - Body: Pilih `x-www-form-urlencoded`
    - KEY `name`, VALUE `admin_director`
    - KEY `password`, VALUE `admin_password`
    - KEY `role`, VALUE `admin` (atau `user`)
  - Catatan: Buat setidaknya satu director dengan `role: admin` dan satu dengan `role: user`.

- **2. Login Director:**

  - Metode: `POST`
  - URL: `http://localhost:8080/login/director`
  - Body: Pilih `x-www-form-urlencoded`
    - KEY `name`, VALUE `admin_director` (atau `user_director`)
    - KEY `password`, VALUE `admin_password` (atau `user_password`)
  - Klik Send. Respons akan mengembalikan `token` yang akan Anda gunakan untuk permintaan selanjutnya.

- **3. Membuat Film Baru:**

  - Metode: `POST`
  - URL: `http://localhost:8080/films`
  - Headers: `Authorization: Bearer <your_token>`
  - Body: Pilih `x-www-form-urlencoded`
    - KEY `title`, VALUE `Judul Film Baru`
  - Klik Send. Film akan otomatis terhubung dengan director yang login.

- **4. Melihat Semua Film (Dipersonalisasi):**

  - Metode: `GET`
  - URL: `http://localhost:8080/films`
  - Headers: `Authorization: Bearer <your_token>`
  - Klik Send. Hanya film yang dibuat oleh director yang login yang akan ditampilkan.

- **5. Melihat Film Berdasarkan ID:**

  - Metode: `GET`
  - URL: `http://localhost:8080/films/1` (ganti `1` dengan ID film yang ada)
  - Headers: `Authorization: Bearer <your_token>`
  - Klik Send.

- **6. Mengupdate Film:**

  - Metode: `PATCH`
  - URL: `http://localhost:8080/films/1` (ganti `1` dengan ID film yang ada)
  - Headers: `Authorization: Bearer <your_token>`
  - Body: Pilih `x-www-form-urlencoded`
    - KEY `title`, VALUE `Judul Film Update`
  - Klik Send. Hanya director pembuat film atau admin yang dapat memperbarui.

- **7. Menghapus Film (Soft Delete):**

  - Metode: `DELETE`
  - URL: `http://localhost:8080/films/1` (ganti `1` dengan ID film yang ada)
  - Headers: `Authorization: Bearer <your_token>`
  - Klik Send. Hanya director pembuat film atau admin yang dapat menghapus.

- **8. Membuat Director Baru (Khusus Admin):**

  - Metode: `POST`
  - URL: `http://localhost:8080/directors`
  - Headers: `Authorization: Bearer <admin_token>`
  - Body: Pilih `x-www-form-urlencoded`
    - KEY `name`, VALUE `nama_director_baru`
    * KEY `password`, VALUE `password_awal`
    - KEY `role`, VALUE `user` (atau `admin`)
  - Klik Send.

- **9. Melihat Semua Director:**

  - Metode: `GET`
  - URL: `http://localhost:8080/directors`
  - Headers: `Authorization: Bearer <your_token>`
  - Klik Send.

- **10. Mengupdate Director:**

  - Metode: `PATCH`
  - URL: `http://localhost:8080/directors/1` (ganti `1` dengan ID director yang ada)
  - Headers: `Authorization: Bearer <your_token>`
  - Body: Pilih `x-www-form-urlencoded`
    - KEY `name`, VALUE `Nama Director Update` (opsional)
    - KEY `password`, VALUE `password_baru` (opsional)
    - KEY `role`, VALUE `admin` (opsional, hanya admin yang bisa mengubah role)
  - Klik Send. Director non-admin hanya bisa mengubah profilnya sendiri.

- **11. Menghapus Director (Soft Delete):**

  - Metode: `DELETE`
  - URL: `http://localhost:8080/directors/1` (ganti `1` dengan ID director yang ada)
  - Headers: `Authorization: Bearer <your_token>`
  - Klik Send. Director non-admin hanya bisa menghapus profilnya sendiri.

- **12. Melihat Film Berdasarkan ID Director (Dipersonalisasi):**

  - Metode: `GET`
  - URL: `http://localhost:8080/getfilmsbydirector/1` (ganti `1` dengan ID director)
  - Headers: `Authorization: Bearer <your_token>`
  - Klik Send. Jika Anda admin, Anda bisa melihat film dari director mana pun. Jika Anda user, hanya film dari ID director Anda sendiri yang akan ditampilkan.

- **13. Melihat Director Berdasarkan ID Film:**
  - Metode: `GET`
  - URL: `http://localhost:8080/getdirectorbyfilm/1` (ganti `1` dengan ID film)
  - Headers: `Authorization: Bearer <your_token>`
  - Klik Send.
