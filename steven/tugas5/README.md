CRUD Marketplace API


Fitur Utama :
1. Autentikasi Toko Sederhana
2. CRUD Product
3. CRUD Store
4. Soft Delete
5. Load Environment dari .env


Jalankan project :
go run main.go


API Endpoint:
1. Auth Store
POST /store/login
Form: owner, password

2. Store
Method	Endpoint	Deskripsi
GET	    /store      Ambil semua toko
POST	/store	    Buat toko baru
GET	    /store/{id}	Detail toko berdasarkan ID
PATCH	/store/{id}	Update toko
DELETE	/store/{id}	Hapus toko (soft delete)

3. Product

Method	Endpoint	    Deskripsi
GET	    /products	    Ambil semua produk
POST	/products	    Tambah produk baru
GET	    /products/{id}	Detail produk berdasarkan ID
PATCH	/products/{id}	Update produk
DELETE	/products/{id}	Hapus produk (soft delete)

