# Sistem Manajemen Restoran - Insight

Dokumen ini memberikan wawasan mengenai **Sistem Manajemen Restoran** yang dikembangkan menggunakan **Go (Golang)** untuk backend. Dalam bagian ini, kita akan membahas tantangan yang dihadapi selama pengembangan, pelajaran yang dipelajari, dan alasan di balik desain dan arsitektur sistem ini.

---

## Kendala yang Dihadapi

### 1. **Desain Database**
   - Salah satu tantangan utama adalah merancang skema database yang skalabel dan efisien untuk menangani banyak relasi, terutama antara chef dan menu.
   - Memastikan **integritas referensial** melalui foreign key (yaitu, menghubungkan chef dengan menu) sambil menghindari data redundan atau query yang kompleks cukup sulit. Pembatasan foreign key membantu menjaga konsistensi data, namun juga memerlukan perhatian ekstra terhadap pembaruan dan penghapusan data.

### 2. **Penanganan Kesalahan dan Validasi**
   - Awalnya, mekanisme penanganan kesalahan belum cukup kuat, yang mengarah pada situasi di mana kesalahan tidak ditangani dengan cara yang ramah pengguna. Sebagai contoh, menangani input yang tidak valid seperti field yang hilang atau ID yang tidak valid saat POST dan GET request membutuhkan waktu untuk disempurnakan.
   - Validasi khususnya menjadi tantangan saat berurusan dengan data form. Memastikan bahwa semua field ada dan diformat dengan benar (misalnya, pengalaman sebagai integer, harga sebagai float) memerlukan pengecekan yang mendalam di backend untuk menghindari masalah di kemudian hari.

### 3. **Autentikasi dan Keamanan**
   - Pada awalnya, tidak ada mekanisme keamanan untuk melindungi proses login. Autentikasi diimplementasikan dengan **pengecekan password dalam bentuk plaintext**, yang tidak ideal. Pada akhirnya, saya menyadari bahwa saya perlu mengimplementasikan hash password (menggunakan bcrypt) untuk mengamankan kredensial pengguna.
   - Memastikan **kontrol akses yang tepat** untuk berbagai endpoint berdasarkan peran (chef, admin, dll.) adalah tantangan lain. Meskipun sistem ini relatif sederhana pada titik ini, implementasi jalur aman sangat penting untuk pengembangan lebih lanjut.

### 4. **Menangani Data yang Hilang atau Tidak Valid**
   - Untuk request yang memerlukan data spesifik (seperti `ID`, `name`, `category`, dll.), memastikan bahwa data yang hilang atau tidak valid tidak merusak aplikasi menjadi hal yang sulit. Banyak kesalahan dapat terjadi jika, misalnya, menu diminta dengan chef ID yang tidak ada atau jika kategori tidak disediakan. Validasi yang tepat dan pesan kesalahan yang bermakna diperlukan untuk meningkatkan pengalaman pengguna.

---

## Pelajaran yang Dipelajari

### 1. **Relasi Database Sangat Penting**
   - Merancang database relasional dengan **pembatasan foreign key** dan **integritas referensial** membuat pengelolaan data antara **chefs**, **menus**, dan **courses** menjadi lebih mudah. Hal ini membantu menghindari redundansi data dan menjaga hubungan yang konsisten.
   - Memahami model **E/R** (Entity-Relationship model) sangat penting ketika bekerja dengan database relasional, terutama ketika menangani banyak relasi antara tabel.

### 2. **Penggunaan yang Efektif dari Paket `http` dan `json` Go**
   - Paket **http** dan **json** di Go sangat penting dalam membangun backend yang tangguh dan efisien untuk menangani request dan mengirimkan respons. Belajar menggunakan paket ini untuk menangani request HTTP, membuat endpoint dinamis, dan menyusun respons JSON sangat penting dalam membangun API yang fungsional.
   - Selain itu, belajar bagaimana cara menangani metode HTTP dengan benar (GET, POST, PATCH, DELETE) dan bagaimana cara mem-parsing data form dan parameter URL di Go adalah pengalaman belajar yang besar.

### 3. **Penanganan Kesalahan dan Validasi di Go**
   - Penanganan kesalahan yang tepat dan validasi sangat penting dalam membangun sistem yang stabil. Proyek ini mengajarkan saya pentingnya menangani edge case dan memastikan data yang valid sebelum berinteraksi dengan database atau melakukan operasi lainnya.
   - Menggunakan **kode status HTTP** dan memastikan pesan yang tepat dikirim untuk setiap kesalahan membantu membuat sistem lebih tangguh.

### 4. **Memahami Desain RESTful API**
   - Merancang sistem berdasarkan prinsip **RESTful** adalah hal yang penting. Dengan mengikuti konvensi REST (menggunakan metode HTTP seperti GET, POST, PATCH, DELETE), endpoint menjadi intuitif dan sesuai dengan praktik standar API.
   - Belajar bagaimana cara menyusun endpoint dan merutekannya dengan tepat untuk memenuhi persyaratan fungsionalitas adalah pelajaran yang berharga dalam desain API.

---

## Alasan Sistem Dibangun Seperti Ini

### 1. **Pemisahan Tanggung Jawab (Separation of Concerns)**
   - Sistem ini dirancang agar modular, memisahkan berbagai tanggung jawab seperti **manajemen chef**, **manajemen menu**, dan **autentikasi** ke dalam paket dan file yang terpisah. Hal ini membuat kode lebih mudah untuk dipelihara dan diperluas, serta memastikan bahwa setiap bagian dari sistem menangani fungsionalitasnya sendiri.

### 2. **Arsitektur yang Sederhana dan Skalabel**
   - Sistem ini dibangun dengan **arsitektur RESTful** yang sederhana, membuatnya mudah untuk menambahkan fitur baru di masa depan. Sebagai contoh, jika entitas baru (misalnya, pesanan, ulasan) perlu ditambahkan, mereka dapat dengan mudah dimasukkan tanpa mengubah struktur yang ada.
   - Dengan menggunakan **Go** dan **MySQL**, sistem ini ringan, cepat, dan dapat dengan mudah diskalakan untuk kasus penggunaan yang lebih kompleks, seperti integrasi dengan sistem eksternal atau menangani sejumlah besar pengguna.

### 3. **Kemudahan Penggunaan**
   - Sistem ini dibangun dengan pengguna di pikiran. Endpoint yang sederhana seperti **GET /chefs**, **POST /menus**, dan **GET /menus/category** memberikan akses mudah ke data, sambil memastikan bahwa validasi dan penanganan kesalahan membuatnya jelas kepada pengguna ketika terjadi kesalahan.
   - Menggunakan **Postman** untuk menguji endpoint membantu memastikan bahwa sistem ini tidak hanya berfungsi tetapi juga ramah pengguna, memberikan umpan balik yang jelas ketika permintaan tidak memenuhi harapan.


### 4. **Skalabilitas untuk Fitur Masa Depan**
   - Sistem ini dirancang dengan mempertimbangkan skalabilitas. Dengan memisahkan fungsionalitas ke dalam berbagai modul (Chefs, Menus, dan Authentication), sistem ini mudah diperluas dengan fitur tambahan, seperti:
     - **Pesanan pelanggan**
     - **Sistem pembayaran**
     - **Peran pengguna (misalnya, Admin, Manager)**

---

## Kesimpulan

Membangun **Sistem Manajemen Restoran** adalah pengalaman belajar yang sangat berharga. Sistem ini membantu saya memperoleh pengalaman praktis dengan **Go**, **MySQL**, dan **RESTful API**. Saya juga belajar pentingnya **keamanan**, **manajemen database**, dan **penanganan kesalahan**. Sistem ini dirancang untuk skalabilitas, modularitas, dan kemudahan pengembangan fitur di masa depan. Proses merancang sistem ini dan mengatasi tantangan memberikan saya pemahaman yang lebih dalam tentang pengembangan backend dan desain API.

Jika ada pertanyaan lebih lanjut, atau jika Anda ingin berkontribusi atau menyarankan perbaikan, jangan ragu untuk menghubungi atau membuat pull request!

