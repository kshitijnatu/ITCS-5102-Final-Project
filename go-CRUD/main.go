package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	StudentID      string  `json:"student_id"`
	Major          string  `json:"major"`
	Year           string  `json:"year"`
	GPA            float64 `json:"gpa"`
	EnrollmentDate string  `json:"enrollment_date"`
}

var students = []Student{
	{ID: 1, Name: "John Doe", Email: "john.doe@university.edu", StudentID: "S1001", Major: "Computer Science", Year: "Sophomore", GPA: 3.7, EnrollmentDate: "2024-08-15"},
	{ID: 2, Name: "Jane Doe", Email: "jane.doe@university.edu", StudentID: "S1002", Major: "Biology", Year: "Junior", GPA: 3.9, EnrollmentDate: "2023-08-15"},
	{ID: 3, Name: "Jim Doe", Email: "jim.doe@university.edu", StudentID: "S1003", Major: "Mathematics", Year: "Senior", GPA: 3.5, EnrollmentDate: "2022-08-15"},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", readRoot)
	mux.HandleFunc("GET /students", listStudents)
	mux.HandleFunc("GET /students/{id}", getStudent)
	mux.HandleFunc("POST /students", createStudent)
	mux.HandleFunc("PUT /students/{id}", updateStudent)
	mux.HandleFunc("DELETE /students/{id}", deleteStudent)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", withCORS(mux)))
}

func readRoot(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, map[string]string{
		"message": "Here is the Go API for the students database",
	})
}

func listStudents(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid student id")
		return
	}

	for _, student := range students {
		if student.ID == id {
			jsonResponse(w, http.StatusOK, student)
			return
		}
	}
	jsonError(w, http.StatusNotFound, "Student not found")
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	for _, existing := range students {
		if existing.ID == student.ID {
			jsonError(w, http.StatusBadRequest, "Student with this id already exists")
			return
		}
	}

	students = append(students, student)
	jsonResponse(w, http.StatusCreated, student)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid student id")
		return
	}

	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if student.ID != id {
		jsonError(w, http.StatusBadRequest, "Student id in body must match path id")
		return
	}

	for i, existing := range students {
		if existing.ID == id {
			students[i] = student
			jsonResponse(w, http.StatusOK, student)
			return
		}
	}
	jsonError(w, http.StatusNotFound, "Student not found")
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Invalid student id")
		return
	}

	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			jsonResponse(w, http.StatusOK, map[string]any{
				"message": "Student deleted successfully",
				"student": student,
			})
			return
		}
	}
	jsonError(w, http.StatusNotFound, "Student not found")
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func jsonResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func jsonError(w http.ResponseWriter, status int, detail string) {
	jsonResponse(w, status, map[string]string{"detail": detail})
}
