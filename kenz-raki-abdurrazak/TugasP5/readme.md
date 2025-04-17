# Seminar CRUD API

Sistem ini merupakan RESTful API untuk mendata pembicara dan acara seminar. Dibangun dengan menggunakan bahasa Go, sistem ini mendukung operasi CRUD serta autentikasi sederhana menggunakan `auth_key`.

## Models

### 1. Speaker
- id: int
- name: string
- expertise: string
- auth_key: string (digunakan untuk autentikasi)

### 2. Event
- id: int
- title: string
- description: string
- speaker_id: int (relasi ke speaker)

## Endpoints

### Speaker Endpoints
- `GET /speakers` - Get all speakers
- `GET /speakers/{id}` - Get speaker by ID
- `POST /speakers` - Create speaker (tidak perlu Auth-Key)
- `PUT /speakers/{id}` - Update speaker
- `DELETE /speakers/{id}` - Delete speaker

### Event Endpoints
- `GET /events` - Get all events (**butuh Auth-Key**)
- `GET /events/{id}` - Get event by ID (**butuh Auth-Key**)
- `POST /events` - Create event (**butuh Auth-Key**)
- `PUT /events/{id}` - Update event (**butuh Auth-Key**)
- `DELETE /events/{id}` - Delete event (**butuh Auth-Key**)

### Gabungan
- `GET /speakers/{id}/events` - Get semua event dari speaker tertentu
- `POST /events/full` - Buat speaker dan event sekaligus

## Auth

Sistem autentikasi dilakukan menggunakan header:

Auth-Key: <auth_key dari speaker>

contoh
Auth-Key: seminarkey123