SISTEM DASHBOARD PENDATAAN USER UNTUK IFTAR GDGoC

Deskripsi:
Sistem ini merupakan aplikasi backend sederhana yang dibuat menggunakan bahasa pemrograman Go. Tujuan utama sistem ini adalah untuk mendata informasi pengguna yang akan mengikuti acara iftar GDGoC. Sistem ini hanya diperuntukkan bagi satu pengguna dengan autentikasi sederhana menggunakan variabel lingkungan (.env).

Fitur Utama:

1. Autentikasi Sederhana:

   - Menggunakan file .env yang berisi variabel NAMA, EMAIL, dan PASSWORD.
   - User harus memasukkan email dan password yang sesuai dengan data di file .env untuk mengakses dashboard.

2. Dashboard Pendataan:
   Setelah berhasil login, pengguna dapat mengakses dashboard yang memiliki fitur berikut:

   - Memilih kendaraan yang digunakan (Kendaraan Pribadi, Bus Kaleng, Nebeng, Travel).
   - Menginput barang yang akan dibawa (jumlah tidak terbatas).
   - Mengisi rekomendasi untuk acara iftar dengan format kategori dan isi rekomendasi.
   - Menginput data teman berupa nama dan divisi.

3. Struktur Folder:
   Struktur folder proyek mengikuti standar clean code, yaitu:
   - controllers: berisi fungsi-fungsi yang mengatur logika dashboard.
   - models: berisi struktur data dan fungsi autentikasi.
   - main.go sebagai entry point aplikasi.

Struktur Folder:
tugas4/
├── .env
├── controllers/
│ └── dashboard.go
├── models/
│ ├── auth.go
│ └── data.go
├── main.go
└── README.txt

Cara Menjalankan:

1. Pastikan Go sudah terinstal di komputer.
2. Jalankan perintah berikut untuk menginstal dependensi:
   go get github.com/joho/godotenv
3. Jalankan aplikasi dengan perintah:
   go run main.go

Autentikasi:

- Sistem autentikasi sederhana menggunakan file .env yang berisi variabel NAMA, EMAIL, dan PASSWORD.
- Pengguna harus memasukkan email dan password yang sesuai dengan isi file .env untuk mengakses dashboard.

Error Handling dan Validasi:

- Sistem dilengkapi dengan validasi input sederhana untuk memastikan tidak terjadi error atau panic selama aplikasi berjalan.

Commit:

- Commit dilakukan secara bertahap menggunakan conventional commit dengan format:
  prefix(scope): pesan
  Contoh: feat(auth): menambahkan autentikasi sederhana menggunakan dotenv

Dibuat oleh:
[Depo Sadrila Hadi]
