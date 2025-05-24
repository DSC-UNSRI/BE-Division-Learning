# Development Notes

## Alasan Arsitektur
Kami menggunakan clean folder separation agar memudahkan maintainability. `controllers`, `models`, dan `routes` dibuat terpisah sesuai fungsinya.

## Kendala
- Awalnya kesulitan dalam mengatur struktur clean code dan modularitas antar file
- Belajar cara meng-handle error dan validasi agar tidak ada panic saat request salah

## Ilmu Baru
- Menggunakan MySQL driver dan GoDotENV untuk mengatur koneksi database
- Mengerti routing dengan Gorilla Mux
- Belajar membuat autentikasi sederhana menggunakan validasi nama & password

## Hal Khusus
- Menghindari komentar dalam kode untuk latihan clean code
- Menggunakan commit per fitur (modular commit)
