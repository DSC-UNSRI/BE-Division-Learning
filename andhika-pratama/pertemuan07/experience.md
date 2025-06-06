# Kendala, Pembelajaran, dan Alasan Pengembangan Sistem

## Kendala yang Dihadapi

1. **Validasi Foreign Key (`lecturer_id`) saat membuat Course**  
    Jadi saat kita membuat row baru di suatu tabel yang punya foreign key dari tabel lain, kita harus buat validasi untuk mastiin apakah foreign key yang kita input itu udah ada dan sesuai dengan prinmary key si induk
   
2. **Soft-delete memperbolehkan penghapusan key yang menjadi foreign key**  
    Kalo kita coba hapus sebuah row yang primary key nya ter-assign ke suatu tabel lain yang berhubungan gitu, biasanya (pake hard-delete) sql akan memberi error ga bisa dihapus. Nah, karena soft-delete ini enggak benar benar "ngehapus', jadinya key key tadi jadi bisa dihapus, sehingga harus ditammbah validasi tambahan. 

---

## Ilmu yang Dipelajari

1. **Penggunaan `net/http` untuk REST API**
   Aku belajar bagaimana sih menggunakan library bawaan Go (`net/http`) untuk membuat endpoint RESTful secara manual, meng-handle berbagai method HTTP (GET, POST, PATCH, DELETE), serta mengatur response header dan status code tanpa menggunakan library tambahan.
2. **Handling Relasi dan Validasi Manual**  
   Akujuga mempelajari kek mana menangani relasi data antar tabel secara manual (tanpa GORM), termasuk validasi ID yang berelasi dan bagaimana menangani error atau kendala lainnya secara eksplisit.

---

## Alasan Sistem Dibangun Demikian

Sebenernya enggak ada alasan begimana begimana buat sistemnya dengan cara "ini", aku cuman buat sistemnya dengan cara "ini" karena segitulah batas pengetahuan dan kemampuan aku,  dan menurut aku, ini udah cukup clear dan rapi (sebenernya ngikutin arsitektur si safar si hehe).