from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from fastapi.middleware.cors import CORSMiddleware
import uvicorn

class Student(BaseModel):
    id: int
    name: str
    email: str
    student_id: str
    major: str
    year: str
    gpa: float
    enrollment_date: str

students = [
    Student(id=1, name="John Doe", email="john.doe@university.edu", student_id="S1001", major="Computer Science", year="Sophomore", gpa=3.7, enrollment_date="2024-08-15"),
    Student(id=2, name="Jane Doe", email="jane.doe@university.edu", student_id="S1002", major="Biology", year="Junior", gpa=3.9, enrollment_date="2023-08-15"),
    Student(id=3, name="Jim Doe", email="jim.doe@university.edu", student_id="S1003", major="Mathematics", year="Senior", gpa=3.5, enrollment_date="2022-08-15"),
]

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/")
def read_root():
    return {"message": "Here is the Python FastAPI for the students database"}

@app.get("/students")
def get_students():
    return students

@app.get("/students/{id}")
def get_student(id: int):
    for student in students:
        if student.id == id:
            return student
    raise HTTPException(status_code=404, detail="Student not found")

@app.post("/students", status_code=201)
def create_student(student: Student):
    if any(s.id == student.id for s in students):
        raise HTTPException(status_code=400, detail="Student with this id already exists")
    students.append(student)
    return student

@app.put("/students/{id}")
def update_student(id: int, student: Student):
    for index, existing in enumerate(students):
        if existing.id == id:
            if student.id != id:
                raise HTTPException(status_code=400, detail="Student id in body must match path id")
            students[index] = student
            return student
    raise HTTPException(status_code=404, detail="Student not found")

@app.delete("/students/{id}")
def delete_student(id: int):
    for index, student in enumerate(students):
        if student.id == id:
            deleted = students.pop(index)
            return {
                "message": "Student deleted successfully",
                "student": deleted
            }
    raise HTTPException(status_code=404, detail="Student not found")
