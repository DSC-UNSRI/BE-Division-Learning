# Sistem Dashboard Iftar GDGoC

## Deskripsi
Sistem ini merupakan dashboard berbasis CLI (Command-Line Interface) yang digunakan untuk mendata peserta yang akan mengikuti Iftar GDGoC.  
Pengguna harus login terlebih dahulu menggunakan email dan password yang tersimpan dalam file `.env`.  
Setelah berhasil login, pengguna dapat menggunakan berbagai fitur yang tersedia dalam sistem ini.

---

## Fitur-Fitur yang Tersedia

### 1. Login dengan Sistem Authentication Sederhana
- Sistem akan membaca email dan password dari file `.env`.
- Jika email dan password cocok, pengguna dapat mengakses dashboard.
- Jika tidak cocok, pengguna tidak bisa masuk.

### 2. Pilih Kendaraan (1 Opsi)
- Pengguna dapat memilih satu dari empat kendaraan untuk pergi ke iftar:
  - Kendaraan Pribadi
  - Bus Kaleng
  - Nebeng
  - Travel
- Pilihan kendaraan akan disimpan dan dapat ditampilkan kembali di dashboard.

### 3. Tambah Barang yang Akan Dibawa
- Pengguna bisa menambahkan barang yang akan dibawa ke iftar.
- Barang yang ditambahkan akan tersimpan dan bisa dilihat kembali.

### 4. Tambah Rekomendasi Acara Iftar
- Pengguna dapat memberikan rekomendasi dalam berbagai kategori, seperti:
  - Film (contoh: Avatar, Naruto)
  - Buku (contoh: Laskar Pelangi)
  - Lagu (contoh: Imagine Dragons - Believer)
- Setiap rekomendasi akan disimpan dengan format kategori dan isi rekomendasi.

### 5. Tambah Teman yang Akan Ikut Iftar
- Pengguna dapat mendata teman-teman yang ikut iftar.
- Setiap teman memiliki nama dan divisi (misalnya: "Aldi - BackEnd", "Rina - UI/UX").

### 6. Lihat Data yang Sudah Dimasukkan
- Sistem akan menampilkan semua data yang telah disimpan:
  - Kendaraan yang dipilih.
  - Barang yang akan dibawa.
  - Rekomendasi yang telah diberikan.
  - Daftar teman yang ikut iftar.
- Jika belum ada data, sistem akan menampilkan pesan "Tidak ada data" di setiap kategori.

### 7. Keluar dari Sistem
- Pengguna dapat memilih opsi "Exit" untuk keluar dari dashboard.

---

## Teknologi yang Digunakan
- Golang sebagai bahasa pemrograman utama.
- Godotenv untuk membaca data dari file `.env`.
- CLI Input Handling menggunakan `bufio` untuk membaca input dari user.
