# Sistem Organisasi Mahasiswa

API untuk mengelola data mahasiswa dan organisasi sederhana

## Endpoint Students

### GET /students

- **Description**: Get all students
- **Response**:
  ```json
  [
    {
      "id": 1,
      "name": "Naufal",
      "email": "naufal@gmail.com",
      "major": "Information Systems",
      "year": 2024,
      "org_id": 1
    }
  ]
  ```

### GET /students/{id}

- **Description**: Get student by ID
- **Response**:
  ```json
  {
    "id": 1,
    "name": "Naufal",
    "email": "naufal@gmail.com",
    "major": "Information Systems",
    "year": 2024,
    "org_id": 1
  }
  ```

### POST /students

- **Description**: Create new student
- **Request Body**:
  ```json
  {
    "id": 1,
    "name": "Naufal",
    "email": "naufal@gmail.com",
    "major": "Information Systems",
    "year": 2024,
    "org_id": 1
  }
  ```

### PUT /students/{id}

- **Description**: Update student
- **Request Body**:
  ```json
  {
    "name": "M. Naufal",
    "email": "naufal@gmail.com",
    "major": "Computer Science",
    "year": 2024
  }
  ```

### DELETE /students/{id}

- **Description**: Delete student

## Endpoint Organizations

### GET /organizations

- **Description**: Get all organizations
- **Response**:
  ```json
  [
    {
      "id": 1,
      "name": "GDSC",
      "type": "UKM"
    }
  ]
  ```

### GET /organizations/{id}

- **Description**: Get organization by ID
- **Response**:
  ```json
  {
    "id": 1,
    "name": "GDSC",
    "type": "UKM"
  }
  ```

### POST /organizations

- **Description**: Create new organization
- **Request Body**:
  ```json
  {
    "name": "GDSC",
    "type": "UKM"
  }
  ```

### PUT /organizations/{id}

- **Description**: Update organization
- **Request Body**:
  ```json
  {
    "name": "GDGoC",
    "type": "UKM"
  }
  ```

### DELETE /organizations/{id}

- **Description**: Delete organization

## Endpoint Gabungan

### GET /organizations/{orgID}/members

- **Description**: Get all members of an organization
- **Response**:
  ```json
  [
    {
      "id": 1,
      "name": "Naufal",
      "email": "naufal@gmail.com",
      "major": "Information Systems",
      "year": 2024
    }
  ]
  ```

## Struktur Database

### Tabel `students`

| Column   | Type         |
| -------- | ------------ |
| id       | INT (PK)     |
| name     | VARCHAR(100) |
| email    | VARCHAR(100) |
| password | VARCHAR(100) |
| major    | VARCHAR(100) |
| year     | INT          |
| org_id   | INT (FK)     |

### Tabel `organizations`

| Column | Type         |
| ------ | ------------ |
| id     | INT (PK)     |
| name   | VARCHAR(100) |
| type   | VARCHAR(100) |
