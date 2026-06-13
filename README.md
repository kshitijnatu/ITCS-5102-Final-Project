# Student CRUD API — Final Project

A multi-language comparison of RESTful CRUD APIs for managing student records. This repository implements the same student resource API across three languages and frameworks, making it easy to compare syntax, structure, and developer experience.

## Project Overview

Each implementation exposes a simple in-memory student database with the same REST endpoints:

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/` | API welcome message |
| `GET` | `/students` | List all students |
| `GET` | `/students/{id}` | Get a student by ID |
| `POST` | `/students` | Create a new student |
| `PUT` | `/students/{id}` | Update an existing student |
| `DELETE` | `/students/{id}` | Delete a student |

### Student Model

All implementations use the same student fields:

| Field | Type | Description |
|-------|------|-------------|
| `id` | integer | Internal record ID |
| `name` | string | Student's full name |
| `email` | string | University email address |
| `student_id` | string | School-assigned ID (e.g. `S1001`) |
| `major` | string | Field of study |
| `year` | string | Academic year (e.g. Freshman, Sophomore) |
| `gpa` | float | Grade point average |
| `enrollment_date` | string | Enrollment date (`YYYY-MM-DD`) |

### Repository Structure

```
Final_Project/
├── python-fastapi-CRUD/     # Python + FastAPI implementation
├── typescript-express-CRUD/ # TypeScript + Express implementation
├── go-CRUD/                 # Go implementation
└── README.md
```

---

## Python (FastAPI)

**Directory:** `python-fastapi-CRUD/`

### Tech Stack

| Technology | Purpose |
|------------|---------|
| [Python 3.13](https://www.python.org/) | Runtime |
| [FastAPI](https://fastapi.tiangolo.com/) | Web framework |
| [Pydantic](https://docs.pydantic.dev/) | Data validation and serialization |
| [Uvicorn](https://www.uvicorn.org/) | ASGI server |

### Setup

```bash
cd python-fastapi-CRUD
python3 -m venv venv
source venv/bin/activate
pip install fastapi uvicorn
```

### Run

**Development (with hot reload):**

```bash
uvicorn main:app --reload
```

**Alternative:**

```bash
python main.py
```

The server starts at **http://localhost:8000**.

### API Documentation

FastAPI auto-generates interactive docs:

- Swagger UI: http://localhost:8000/docs
- ReDoc: http://localhost:8000/redoc

### Example Requests

**Create a student:**

```bash
curl -X POST http://localhost:8000/students \
  -H "Content-Type: application/json" \
  -d '{
    "id": 4,
    "name": "Alice Smith",
    "email": "alice.smith@university.edu",
    "student_id": "S1004",
    "major": "Physics",
    "year": "Freshman",
    "gpa": 3.8,
    "enrollment_date": "2025-08-15"
  }'
```

**Get all students:**

```bash
curl http://localhost:8000/students
```

**Get one student:**

```bash
curl http://localhost:8000/students/1
```

**Update a student:**

```bash
curl -X PUT http://localhost:8000/students/1 \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@university.edu",
    "student_id": "S1001",
    "major": "Computer Science",
    "year": "Junior",
    "gpa": 3.8,
    "enrollment_date": "2024-08-15"
  }'
```

**Delete a student:**

```bash
curl -X DELETE http://localhost:8000/students/1
```

---

## TypeScript (Express)

**Directory:** `typescript-express-CRUD/`

### Tech Stack

| Technology | Purpose |
|------------|---------|
| [Node.js](https://nodejs.org/) | Runtime |
| [TypeScript](https://www.typescriptlang.org/) | Typed JavaScript |
| [Express](https://expressjs.com/) | Web framework |

### Planned Setup

```bash
cd typescript-express-CRUD
npm init -y
npm install express
npm install -D typescript @types/express @types/node ts-node nodemon
npx tsc --init
```

### Planned Run

```bash
npm run dev
```

The server will run at **http://localhost:3000** (or another configured port).

### Status

> **Not yet implemented.** This directory is reserved for the TypeScript + Express version of the student CRUD API.

---

## Go

**Directory:** `go-CRUD/`

### Tech Stack

| Technology | Purpose |
|------------|---------|
| [Go](https://go.dev/) | Language and runtime |
| [net/http](https://pkg.go.dev/net/http) | Standard library HTTP server |
| [encoding/json](https://pkg.go.dev/encoding/json) | JSON encoding/decoding |

### Planned Setup

```bash
cd go-CRUD
go mod init go-crud
```

### Planned Run

```bash
go run main.go
```

The server will run at **http://localhost:8080** (or another configured port).

### Status

> **Not yet implemented.** This directory is reserved for the Go version of the student CRUD API.

---

## Notes

- All implementations use an **in-memory** data store. Data is lost when the server restarts.
- CORS is enabled on the Python implementation to allow cross-origin requests from front-end clients.
- Each service should run on a **different port** if you start more than one at the same time.
