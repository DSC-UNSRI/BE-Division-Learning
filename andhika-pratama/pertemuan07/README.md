# Course Data Management

## System Overview

This is a backend system designed for the management of lecturers and courses, allowing for the creation, retrieval, updating, and deletion of both lecturers and courses. The system also includes advanced functionality to manage lecturer data, including filtering lecturers based on their city of residence.

### Key Models

- **Lecturer**: Contains information about lecturers, such as their name and credentials.
- **Course**: Represents a course, including the course name and related lecturer.
- **Address**: Stores the address details of lecturers, including the city of residence.

---

## Features

### CRUD Operations

- **Lecturer**:
  - Create, retrieve, update, and delete lecturer data.
  - Retrieve lecturer by ID.
  - Retrieve lecturers by city.
  
- **Course**:
  - Create, retrieve, update, and delete courses.
  - Retrive course by course ID.
  - Retrieve courses by lecturer ID.

---

## Endpoints

### Authentication

- **POST** `/login`: Authenticates a user based on the provided credentials.

### Lecturer Endpoints

- **GET** `/lecturers`: Retrieve a list of all lecturers.
- **POST** `/lecturers`: Create a new lecturer.
- **GET** `/lecturers/{id}`: Retrieve a lecturer by ID.
- **PATCH** `/lecturers/{id}`: Update a lecturer's details.
- **DELETE** `/lecturers/{id}`: Delete a lecturer by ID.
- **GET** `/lecturersbycity/{city}`: Retrieve lecturers living in a specific city.

### Course Endpoints

- **GET** `/courses`: Retrieve a list of all courses.
- **POST** `/courses`: Create a new course.
- **GET** `/courses/{id}`: Retrieve a course by ID.
- **PATCH** `/courses/{id}`: Update a course's details.
- **DELETE** `/courses/{id}`: Delete a course by ID.
- **GET** `/coursesbylecturer/{lecturer_id}`: Retrieve courses taught by a specific lecturer.

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