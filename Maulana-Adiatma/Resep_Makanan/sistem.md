Sistem ini adalah aplikasi backend sederhana yang dibuat menggunakan Golang, dengan penyimpanan data menggunakan MySQL (via XAMPP), dan dapat diakses serta di  uji menggunakan Postman.
Tema dari aplikasi ini adalah pendataan resep makanan berdasarkan negara asalnya. Data disimpan dalam satu tabel saja: data_resep.

- Sistem Resep makanan ini maul gunakan MySQL dari XAMPP dimana untuk 2 model (resep dan makanan) aku hubungkan dari database yang aku buat db_makanan dengan 2 tabel
    1. data_resep
    2. data_negara
- setelah model lanjut dalam pembuatan function alias controlnya agar data di database tersebut bisa melakukam CRUD dan karena kebetulan ini masih awal awal aku ngebuat SELECT * yang artinya
  aku ngambil semua elemen tabel dari tabel yang aku tuju
- disini aku juga ada nambah get baru yang namanya gorilla, yang berguna untuk memenuhi kekurangan http yang dimiliki oleh golang (sepemahaman aku saja ini)
- dan setelah sudah nanti akan lanjut pembuatan 5 endpoint dengan menggunakan handle func dengan nambah nama path. controlnya dan method nya untuk mungkin bisa kita akses
  di postman
- lanjut ngebuat main.go untuk menjalankan sistem nya dan disini aku ngebuat fungsi InitDB() agar dapat membaca file .env dan membuat koneksi ke database MySQL.

# Resep Makanan

## Models
### Country
- id (int)
- namaNegara (string)

### resep
- id (int)
- namaResep (string)
- DeskripsiResep (string)

## Endpoints
### Country
- /negara = Methods("GET")
- /neagara = Methods("POST")
- /negara/{id} = Methods("GET")
- /negara/{id} = Methods("PUT")
- /negara/{id} = Methods("DELETE")

### resep
- /resep = Methods("GET")
- /resep = Methods("POST")
- /resep/{id} = Methods("GET")
- /resep/{id} = Methods("PUT")
- /resep/{id} = Methods("DELETE")
 
