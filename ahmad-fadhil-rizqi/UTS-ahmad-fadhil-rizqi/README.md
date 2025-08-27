# API Forum Tanya Jawab (Quora Lite)

## Deskripsi Sistem

Proyek ini adalah sebuah API backend untuk aplikasi forum tanya jawab sederhana yang dibuat menggunakan package standar `net/http` dari Go, tanpa menggunakan _framework_ eksternal. API ini menyediakan fungsionalitas bagi pengguna untuk mendaftar, login, bertanya, dan menjawab pertanyaan.

Sistem ini menerapkan autentikasi berbasis Bearer Token dan memiliki pembagian hak akses antara pengguna `free` dan `premium` untuk beberapa fiturnya, sesuai dengan brief tugas dari "PT Backend ABADI".

## Fitur Utama

- **Manajemen Pengguna:**

  - Registrasi pengguna baru (`/register`).
  - Login pengguna (`/login`) yang menghasilkan Bearer Token.
  - Pengambilan data profil pengguna yang sedang login (`/profile`).

- **Manajemen Pertanyaan (CRUD):**

  - Membuat pertanyaan baru (`/questions`).
  - Mendapatkan daftar semua pertanyaan (`/questions`).
  - Mendapatkan detail satu pertanyaan beserta semua jawabannya (`/questions/{id}`).
  - Memperbarui pertanyaan (hanya oleh pemilik dan pengguna premium).
  - Menghapus pertanyaan (hanya oleh pemilik, _soft delete_).

- **Manajemen Jawaban (CRUD):**

  - Membuat jawaban untuk pertanyaan tertentu (`/questions/{id}/answers`).
  - Memperbarui jawaban (hanya oleh pemilik).
  - Menghapus jawaban (hanya oleh pemilik, _soft delete_).

- **Fitur Lanjutan:**
  - **Autentikasi & Otorisasi**: Sistem token Bearer untuk melindungi _endpoint_ yang membutuhkan login.
  - **Tier Pengguna (Free vs. Premium)**: Pengguna `free` memiliki batasan waktu 30 detik antar pembuatan pertanyaan.
  - **Lupa Password**: Mekanisme lupa password tanpa email menggunakan pertanyaan keamanan.
  - **Token Kadaluwarsa**: Token yang dihasilkan saat login memiliki masa berlaku 24 jam untuk meningkatkan keamanan.

## Cara Penggunaan (API Endpoint)

Pastikan aplikasi Go Anda berjalan di `http://localhost:8080`. Untuk semua permintaan yang memerlukan autentikasi, sertakan header `Authorization: Bearer <TOKEN_ANDA>`. Ganti `<TOKEN_ANDA>` dengan token yang Anda dapatkan setelah login.

---

### 1. Autentikasi & Profil

#### **Registrasi Pengguna Baru**

- **Metode**: `POST`
- **URL**: `http://localhost:8080/register`
- **Body**: `x-www-form-urlencoded`
  - `KEY`: `username`, `VALUE`: `nama_user_baru`
  - `KEY`: `password`, `VALUE`: `password_rahasia`
- **Hasil**: Pesan sukses pendaftaran.

#### **Login Pengguna**

- **Metode**: `POST`
- **URL**: `http://localhost:8080/login`
- **Body**: `x-www-form-urlencoded`
  - `KEY`: `username`, `VALUE`: `nama_user_baru`
  - `KEY`: `password`, `VALUE`: `password_rahasia`
- **Hasil**: Mendapatkan token Bearer. **Simpan token ini untuk digunakan pada _request_ selanjutnya.**
  ```json
  {
    "message": "Login berhasil",
    "token": "..."
  }
  ```

#### **Melihat Profil**

- **Metode**: `GET`
- **URL**: `http://localhost:8080/profile`
- **Headers**: `Authorization: Bearer <TOKEN_ANDA>`
- **Hasil**: Data profil pengguna yang sedang login.

---

### 2. Pertanyaan (Questions)

#### **Membuat Pertanyaan Baru**

- **Metode**: `POST`
- **URL**: `http://localhost:8080/questions/`
- **Headers**: `Authorization: Bearer <TOKEN_ANDA>`
- **Body**: `x-www-form-urlencoded`
  - `KEY`: `question`, `VALUE`: `Apa itu Go-lang?`
- **Catatan**: Pengguna `free` harus menunggu 30 detik antar post.

#### **Melihat Semua Pertanyaan**

- **Metode**: `GET`
- **URL**: `http://localhost:8080/questions/`
- **Hasil**: Daftar semua pertanyaan yang ada.

#### **Melihat Detail Satu Pertanyaan**

- **Metode**: `GET`
- **URL**: `http://localhost:8080/questions/{id}` (Contoh: `/questions/1`)
- **Hasil**: Detail pertanyaan beserta semua jawabannya.

#### **Memperbarui Pertanyaan**

- **Metode**: `PUT` atau `PATCH`
- **URL**: `http://localhost:8080/questions/{id}`
- **Headers**: `Authorization: Bearer <TOKEN_ANDA>`
- **Body**: `x-www-form-urlencoded`
  - `KEY`: `question`, `VALUE`: `Teks pertanyaan yang baru.`
- **Catatan**: Hanya bisa dilakukan oleh pemilik pertanyaan dengan tier `premium`.

#### **Menghapus Pertanyaan**

- **Metode**: `DELETE`
- **URL**: `http://localhost:8080/questions/{id}`
- **Headers**: `Authorization: Bearer <TOKEN_ANDA>`
- **Catatan**: Hanya bisa dilakukan oleh pemilik pertanyaan.

---

### 3. Jawaban (Answers)

#### **Membuat Jawaban Baru**

- **Metode**: `POST`
- **URL**: `http://localhost:8080/questions/{id}/answers`
- **Headers**: `Authorization: Bearer <TOKEN_ANDA>`
- **Body**: `x-www-form-urlencoded`
  - `KEY`: `answer`, `VALUE`: `Ini adalah jawaban saya.`

#### **Memperbarui Jawaban**

- **Metode**: `PUT` atau `PATCH`
- **URL**: `http://localhost:8080/answers/{id}`
- **Headers**: `Authorization: Bearer <TOKEN_ANDA>`
- **Body**: `x-www-form-urlencoded`
  - `KEY`: `answer`, `VALUE`: `Teks jawaban yang baru.`
- **Catatan**: Hanya bisa dilakukan oleh pemilik jawaban.

#### **Menghapus Jawaban**

- **Metode**: `DELETE`
- **URL**: `http://localhost:8080/answers/{id}`
- **Headers**: `Authorization: Bearer <TOKEN_ANDA>`
- **Catatan**: Hanya bisa dilakukan oleh pemilik jawaban.

---

### 4. Lupa Password

#### **Langkah 1: Mengatur Pertanyaan Keamanan (Saat Login)**

- **Metode**: `POST`
- **URL**: `http://localhost:8080/security-question`
- **Headers**: `Authorization: Bearer <TOKEN_ANDA>`
- **Body**: `x-www-form-urlencoded`
  - `KEY`: `question`, `VALUE`: `Siapa nama hewan peliharaan pertama saya?`
  - `KEY`: `answer`, `VALUE`: `Milo`

#### **Langkah 2: Meminta Pertanyaan Keamanan (Saat Lupa Password)**

- **Metode**: `POST`
- **URL**: `http://localhost:8080/forgot-password/get-question`
- **Body**: `x-www-form-urlencoded`
  - `KEY`: `username`, `VALUE`: `nama_user_yang_lupa`
- **Hasil**: Mendapatkan pertanyaan keamanan pengguna tersebut.

#### **Langkah 3: Mereset Password**

- **Metode**: `POST`
- **URL**: `http://localhost:8080/forgot-password/reset`
- **Body**: `x-www-form-urlencoded`
  - `KEY`: `username`, `VALUE`: `nama_user_yang_lupa`
  - `KEY`: `answer`, `VALUE`: `Milo` (jawaban yang benar)
  - `KEY`: `new_password`, `VALUE`: `password_baru_yang_aman`
- **Hasil**: Pesan sukses jika jawaban benar.
