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
│   └── benchmark.py         # Performance load-test client
├── typescript-express-CRUD/ # TypeScript + Express implementation
│   └── src/benchmark.ts     # Performance load-test client
├── go-CRUD/                 # Go implementation
│   └── benchmark/           # Performance load-test client
└── README.md
```

---

## Performance Comparison

Each language has its own benchmark client that sends concurrent `GET /students` requests to measure throughput and latency. Run **one server at a time**, then run its matching benchmark in a second terminal.

| Language   | Port | Benchmark file              |
|------------|------|-----------------------------|
| Go         | 8080 | `go-CRUD/benchmark/main.go` |
| Python     | 8000 | `python-fastapi-CRUD/benchmark.py` |
| TypeScript | 3000 | `typescript-express-CRUD/src/benchmark.ts` |

**Defaults:** 10,000 requests, 50 concurrent workers, endpoint `GET /students`

**Metrics to compare:** Requests/sec (higher is faster), avg latency, p95 latency (lower is faster)

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

### Performance Benchmark

```bash
# Terminal 1
uvicorn main:app --host 0.0.0.0 --port 8000

# Terminal 2
python benchmark.py
```

**Results** (10,000 requests, 50 concurrent workers, `GET /students`):

```
Language:     Python (FastAPI client)
URL:          http://localhost:8000/students
Duration:     5.361s
Requests/sec: 1865.35
Successful:   10000
Failed:       0
Avg latency:  ~26ms
```

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
| [CORS](https://www.npmjs.com/package/cors) | Cross-origin request handling |
| [ts-node](https://typestrong.org/ts-node/) | Run TypeScript directly in development |
| [Nodemon](https://nodemon.io/) | Auto-restart on file changes |

### Project Structure

```
typescript-express-CRUD/
├── src/
│   ├── main.ts      # Express app and route handlers
│   └── Student.ts   # Student interface
├── dist/            # Compiled JavaScript (generated by tsc)
├── package.json
└── tsconfig.json
```

### Setup

```bash
cd typescript-express-CRUD
npm install
```

### Run

**Development (with hot reload):**

```bash
npm run dev
```

The server starts at **http://localhost:3000**.

> If port 3000 is already in use, stop the other process (`lsof -ti :3000 | xargs kill`) or change the `PORT` constant in `src/main.ts`.

### Performance Benchmark

```bash
# Terminal 1
npm run build && npm start

# Terminal 2
npm run benchmark
```

**Results** (10,000 requests, 50 concurrent workers, `GET /students`):

```
Language:     TypeScript (Express client)
URL:          http://localhost:3000/students
Duration:     0.857s
Requests/sec: 11667.65
Successful:   10000
Failed:       0
Avg latency:  ~2ms
```

### Example Requests

**Create a student:**

```bash
curl -X POST http://localhost:3000/students \
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
curl http://localhost:3000/students
```

**Get one student:**

```bash
curl http://localhost:3000/students/1
```

**Update a student:**

```bash
curl -X PUT http://localhost:3000/students/1 \
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
curl -X DELETE http://localhost:3000/students/1
```

---

## Go

**Directory:** `go-CRUD/`

### Tech Stack

| Technology | Purpose |
|------------|---------|
| [Go 1.22+](https://go.dev/) | Language and runtime |
| [net/http](https://pkg.go.dev/net/http) | Standard library HTTP server |
| [encoding/json](https://pkg.go.dev/encoding/json) | JSON encoding/decoding |

### Project Structure

```
go-CRUD/
├── main.go
├── benchmark/
│   └── main.go
└── go.mod
```

### Setup

Requires [Go 1.22 or later](https://go.dev/dl/) (uses Go 1.22+ route patterns like `GET /students/{id}`).

```bash
cd go-CRUD
go mod download
```

### Run

```bash
go run main.go
```

The server starts at **http://localhost:8080**.

> If port 8080 is already in use, stop the other process (`lsof -ti :8080 | xargs kill`) or change the port in `main.go`.

### Performance Benchmark

```bash
# Terminal 1
go run main.go

# Terminal 2
cd benchmark && go run .
```

**Results** (10,000 requests, 50 concurrent workers, `GET /students`):

```
Language:     Go
URL:          http://localhost:8080/students
Duration:     295ms
Requests/sec: 33937.99
Successful:   10000
Failed:       0
Avg latency:  1.42ms
```

### Example Requests

**Create a student:**

```bash
curl -X POST http://localhost:8080/students \
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
curl http://localhost:8080/students
```

**Get one student:**

```bash
curl http://localhost:8080/students/1
```

**Update a student:**

```bash
curl -X PUT http://localhost:8080/students/1 \
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
curl -X DELETE http://localhost:8080/students/1
```

---

## Notes

- All implementations use an **in-memory** data store. Data is lost when the server restarts.
- CORS is enabled on the Python, TypeScript, and Go implementations to allow cross-origin requests from front-end clients.
- Each service should run on a **different port** if you start more than one at the same time.
