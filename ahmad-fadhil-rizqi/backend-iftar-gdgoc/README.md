Penjelasan Sistem Dashboard Iftar GDGoC

1. Deskripsi Sistem
   Sistem ini adalahdashboard berbasis termina yang digunakan untuk mencatat data peserta Iftar GDGoC.  
   Sistem akanmemeriksa autentikasi user terlebih dahul sebelum bisa mengakses menu utama.

Pengguna dapat melakukan beberapa operasi utama:

1.Memilih kendaraa (hanya bisa memilih satu dari Bus Kaleng, Mobil Pribadi, Travel, atau Nebeng).
2.Mencatat barang yang akan dibaw ke iftar.
3.Menambahkan rekomendasi hiburan atau makana untuk acara iftar.
4.Mendata teman yang ikut ifta beserta divisinya.
5.Melihat seluruh data yang telah dimasukka.
6.Keluar dari siste.

Setiap aktivitas pengguna akan dicatat dalam file`data/log_iftar.txt untuk keperluan pencatatan.

---

2. Langkah-langkah Penggunaan

2.1. Persiapan Awal
Sebelum menjalankan aplikasi, pastikan:

- Golang sudah terinstal (`go version` untuk memeriksa).
- Buat file `.env` dengan format berikut:
