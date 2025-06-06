# Restaurant Management System

This is a **Restaurant Management System** built using **Go (Golang)** for backend, providing endpoints to manage **chefs** and **menus**. The system allows adding, updating, deleting, and viewing chefs and menu items for the restaurant. It also includes authentication for login functionality.

---

## Features

- **Chef Management**: Add, update, delete, and view chefs.
- **Menu Management**: Add, update, delete, and view menu items.
- **Login Authentication**: Authenticate chefs using their credentials (username and password).
- **Menu Retrieval by Chef**: Retrieve menus assigned to a specific chef.
- **Menu Retrieval by Category**: Retrieve menus categorized by their types (e.g., Main Course, Dessert).
- **API Responses**: All API responses are returned in JSON format for easy integration.

---

## Technologies Used

- **Go (Golang)**: Backend programming language.
- **MySQL**: Database for storing chefs and menu data.
- **Postman**: Used for testing API endpoints.

---

## Setup and Installation

### Prerequisites

Before setting up this project, ensure you have the following installed:

- **Go** (Version 1.16+): [Download Go](https://golang.org/dl/)
- **MySQL**: [Download MySQL](https://dev.mysql.com/downloads/)
- **Postman**: [Download Postman](https://www.postman.com/downloads/)

## Endpoints

### Authentication

- **POST** `/login`: Authenticates a user based on the provided credentials.

### MENUS Endpoints

- **GET** `/menus`: Retrieve a list of all menus.
- **POST** `/menus`: Create a new menus.
- **GET** `/menus/{id}`: Retrieve a menus by ID.
- **PATCH** `/menus/{id}`: Update a menus's details.
- **DELETE** `/menus/{id}`: Delete a menus by ID.
- **GET** `/getmenusbychef/{chef_id}`: Retrieve menus by chef id.
- **GET** `/getmenusbycategory/{category}`: Retrieve menus by category.


### CHEFS Endpoints

- **GET** `/chefs`: Retrieve a list of all chefs.
- **POST** `/chefs`: Create a new chefs.
- **GET** `/chefs/{id}`: Retrieve a chefs by ID.
- **PATCH** `/chefs/{id}`: Update a chefs's details.
- **DELETE** `/chefs/{id}`: Delete a chefs by ID.

---

## Installation

To set up and run the system locally, follow the steps below:

1. Clone the repository.
2. Install required dependencies:
   ```bash
   go get github.com/go-sql-driver/mysql
   go get github.com/joho/godotenv
   ```
3. Configure the .env file with your MySQL database connection details:
   ```bash
   DB_HOST=your_db_host
   DB_PORT=your_db_port
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   ```
4. Run the application:
   ```bash
   go run main.go
   ```