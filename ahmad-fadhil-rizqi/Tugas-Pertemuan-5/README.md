# API CRUD Film

## Deskripsi Sistem

Proyek ini adalah sebuah API yang dibuat menggunakan package standar `net/http`. API ini menyediakan
fungsi dasar CRUD (Create, Read, Update, Delete) untuk mengelola data film dan juga dilengkapi mekanisme
login pengguna yang simpel.

## Fitur Utama

- **Manajemen Film (CRUD):**

  - Menampilkan semua film (`GET /films`)
  - Menampilkan film tertentu berdasarkan ID (`GET /films/{id}`)
  - Membuat data film baru (`POST /films`)
  - Memperbarui data film (`PATCH /films/{id}`)
  - Menghapus data film (Soft Delete) (`DELETE /films/{id}`)

- **Autentikasi Sederhana:**
  - Endpoint login (`POST /login`) yang mengecek `username` dan `auth_key` dari pengguna ke database.

**Cara Menggunakan (Testing API dengan Postman):**

- **Membuat Film Baru:**

  - Metode: `POST`
  - URL: `/films`
  - Body: Pilih `x-www-form-urlencoded`
    - KEY `title`, VALUE `Judul Film Baru`
    - KEY `director`, VALUE `Nama Sutradara`
  - Klik Send.

- **Melihat Semua Film:**

  - Metode: `GET`
  - URL: `/films`
  - Klik Send.

- **Melihat Film Berdasarkan ID:**

  - Metode: `GET`
  - URL: `/films/1` (ganti `1` dengan ID film yang ada)
  - Klik Send.

- **Mengupdate Film:**

  - Metode: `PATCH`
  - URL: `/films/1` (ganti `1` dengan ID film yang ada)
  - Body: Pilih `x-www-form-urlencoded`
    - KEY `title`, VALUE `Judul Film Update` (atau field lain)
  - Klik Send.

- **Menghapus Film (Soft Delete):**

  - Metode: `DELETE`
  - URL: `/films/1` (ganti `1` dengan ID film yang ada)
  - Klik Send.

- **Membuat User (Untuk Login):**

  - Metode: `POST`
  - URL: `/users`
  - Body: Pilih `x-www-form-urlencoded`
    - KEY `username`, VALUE `namauser`
    - KEY `auth`, VALUE `kunciRahasia`

- **Login:**
  - Metode: `POST`
  - URL: `/login`
  - Body: Pilih `x-www-form-urlencoded`
    - KEY `username`, VALUE `namauser` (yang sudah dibuat)
    - KEY `auth`, VALUE `kunciRahasia` (yang sudah dibuat)
  - Klik Send. Jika berhasil, akan ada pesan sukses. Jika salah, akan muncul error Unauthorized.
