# Tugas Pertemuan 12 – Fullstack (React + Fiber + GORM)

Proyek ini berisi frontend (Vite + React + TS) dan backend (Go Fiber + GORM, MySQL/MariaDB) yang saling terhubung melalui proxy Vite ke `/api`.

## Struktur

```
pert12/
├─ frontend/      # React + Vite + TypeScript
└─ backend8/      # Go Fiber + GORM (MySQL)
```

## Menjalankan Backend

```
cd backend8
go mod tidy
go run main.go
```

Server berjalan di `http://127.0.0.1:3000`.

## Menjalankan Frontend

```
cd frontend
npm install
npm run dev
```

Vite akan berjalan di `http://localhost:5173` dan mem-proxy `/api` ke backend `http://localhost:3000`.

## Autentikasi

- JWT disimpan sebagai cookie httpOnly (`token`).
- Default avatar saat signup dipilih acak dari `https://i.pravatar.cc/150?img={1..70}`.

## Endpoint API (ringkas)

Base URL: `/api`

- Auth

  - `POST /login` (multipart form: `email`, `password`)
  - `POST /register` (multipart form: `name`, `email`, `password`, `role`="user|admin")
  - `POST /logout`
  - `GET /me` (butuh cookie `token`)

- Event

  - `GET /event`
  - `POST /event` (admin) body JSON: `{ location: string }`
  - `PATCH /event/:id` (admin) multipart: `location?`, `start?`, `cover?` (file)
  - `DELETE /event/:id` (admin)

- User
  - `PATCH /profile/:id` multipart: `name?`, `profile_picture?` (file)

## Catatan

- Jika gambar profil/cover tidak ada, backend akan mengisi secara default:
  - Profil: avatar acak pravatar
