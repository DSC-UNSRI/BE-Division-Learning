# Dashboard CLI System

This is a simple CLI-based dashboard system for managing users, friends, and recommendations for an Iftar event.

## Features

- **Authentication**: Simple authentication using credentials stored in a `.env` file.
- **Dashboard Management**: Allows users to enter data related to their participation in the event.
- **Friend Management**: Users can add friends along with their division.
- **Recommendation System**: Users can suggest categories like movies, books, etc.
- **CRUD Operations**: Basic create and update operations for dashboard data.

## Project Structure

```
tugas-4/
├── controllers/
│   ├── auth.go         # Handles authentication logic
│   ├── dashboard.go    # Manages dashboard-related features
│   ├── friend.go       # Handles friend-related operations
│
├── models/
│   ├── dashboard.go    # Data structure for dashboard
│   ├── friend.go       # Data structure for friend management
│   ├── user.go         # User model for authentication
│
├── .env.example        # Example environment variable file
├── .gitignore          # Git ignore file
├── go.mod              # Go module file
├── go.sum              # Dependency lock file
├── main.go             # Entry point of the application
├── README.md           # Documentation
```

## How to Run

1. Clone the repository.
2. Copy `.env.example` to `.env` and fill in your credentials.
3. Run the following commands:
   ```sh
   go mod tidy
   go run main.go
   ```

## Future Improvements

- Implement full CRUD support for all entities.
- Improve error handling and input validation.
- Implement persistent storage (e.g., database) instead of in-memory storage.

---

Developed as part of the "tugas-4" assignment.
