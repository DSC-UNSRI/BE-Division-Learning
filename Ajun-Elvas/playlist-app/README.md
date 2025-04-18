# Playlist App

A simple CRUD system for managing artists and songs using Go.

## Features

- CRUD for Artist
- CRUD for Song
- Simple User Auth with Header
- Clean architecture

## How to Run

1. Setup MySQL with database `playlist_db`
2. Update `.env` with your DB credentials
3. Run:
    ```bash
    go mod tidy
    go run main.go
    ```
