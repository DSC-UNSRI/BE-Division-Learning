# UTS_BE - Backend API Sistem Tanya Jawab

Backend RESTful API ini dibuat untuk sistem tanya jawab dengan fitur:
- Registrasi & Login user
- Role-based access (`free` & `premium`)
- Token authentication dengan waktu kadaluarsa
- CRUD untuk pertanyaan dan jawaban
- Highlight khusus untuk user `premium`
- Pembatasan pertanyaan untuk user `free`

---

## 🔧 Teknologi

- Go (Golang)
- MySQL (melalui Laragon)
- bcrypt untuk hash password dan kode rahasia
- Token generator internal (bukan JWT)
- Struktur modular (controllers, middleware, routes, utils, models)

---

