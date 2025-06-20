# Programmer Hub Backend System

## Requirements

- **MySql** (Database)
- **Go** >= ^1.23.4

---

## How to Run

1. **Setup Environment**:

   - Copy `.env.example` to `.env` and configure the values (e.g., database credentials, API keys).

2. **Run the Server**:
   - Run in terminal using command "go run main.go" .

---

## The Interesting Parts

This backend system for a **Programmer Hub** is designed for **scalability** and **low latency**, optimized to handle collaborative coding environments, resource sharing, and developer interactions. Key highlights:

- **Clean Architecture** and **Port-Adapter Pattern** for maintainability and flexibility.
- **Domain-Driven Design (DDD)** to encapsulate core logic for developer workflows and resource management.
- Decoupled infrastructure layer (e.g., REST API, MySQL) for easy future integrations (e.g., Redis, gRPC).

_Note: Some advanced collaboration features remain in progress due to time constraints (2-days development window)._

---

## POSTMAN Link

[LINK](https://lunar-robot-614267.postman.co/workspace/My-Workspace~492001a6-6bb2-4a3f-83e0-682bf5145535/collection/28360257-14fb3cfc-64be-418d-975e-438517d17953?action=share&creator=28360257).

---

## Directory Structure

Follows **Clean Architecture** and **Port-Adapter Pattern** for dependency management.
programmer-hub
├── config # Environment configuration (env.go)
├── database # Database connections and migrations
│ ├── database.go
│ └── migrate.go
├── handlers # HTTP controllers
│ ├── programmerHandler.go # Programmer-specific logic
│ └── userHandler.go
├── models # Data models
│ ├── programmer.go # Programmer entity
│ └── user.go
├── repositories # Database interaction layer
│ ├── programmerRepository.go # Programmer data operations
│ └── userRepository.go
├── routes # API endpoint definitions
│ └── routes.go
├── services # Business logic layer
│ ├── programmerService.go # Core programmer workflows
│ └── userService.go
└── utils # Utilities (auth, conversions)
├── authenticate.go
└── conversion.go

### Key Components:

- **`handlers/programmerHandler.go`**: Manages API logic for programmer-specific features (e.g., project collaboration, skill management) (future development).
- **`services/programmerService.go`**: Implements domain logic for programmer workflows.
- **`repositories/programmerRepository.go`**: Handles database operations for programmer data.
- **`config/env.go`**: Centralizes environment variable management for scalability.

## Architecture Overview

### Domain Layer (`models/`, `services/`)

- Contains **business rules** and **core entities** (e.g., `programmer.go` defines programmer attributes and behaviors).
- Uses **DDD principles** to separate domain logic from infrastructure.

### Infrastructure Layer (`database/`, `handlers/`)

- Acts as a gateway to external systems (e.g., REST API, MySQL).
- Easily replaceable (e.g., switch databases without affecting domain logic).

### Adapter Layer (`repositories/`, `utils/`)

- Converts data between external formats and domain models.
- Example: `programmerRepository.go` maps database rows to `Programmer` structs.
