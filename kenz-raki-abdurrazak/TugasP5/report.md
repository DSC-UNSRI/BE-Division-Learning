# Laporan Pengembangan Sistem Seminar

## Kendala yang Dihadapi
1. Penyesuaian struktur modular dengan clean architecture sederhana
2. Implementasi middleware sederhana untuk autentikasi
3. Menyesuaikan struktur query SQL dengan relasi antara speaker dan event
4. integrasi dengan database yang ada
5. konflik pada routes

## Ilmu yang Dipelajari
- Cara membuat modular API dengan Go
- Penerapan autentikasi menggunakan `auth_key`
- Penggunaan middleware di Go dengan gorilla/mux
- Penanganan error dan validasi sederhana pada setiap endpoint
- Struktur dasar clean architecture (separated models, controllers, routes)

## Alasan Desain Sistem
- Menggunakan field `auth_key` untuk autentikasi agar mudah digunakan
- CRUD lengkap diterapkan 
- Endpoint gabungan 
- Middleware ditambahkan pada endpoint yang perlu proteksi

