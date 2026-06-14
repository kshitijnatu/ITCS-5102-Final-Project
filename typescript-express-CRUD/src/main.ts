import express, { Request, Response } from "express";
import cors from "cors";
import { Student } from "./Student";

const app = express();
const PORT = 3000;

app.use(cors());
app.use(express.json());

const students: Student[] = [
  { id: 1, name: "John Doe", email: "john.doe@university.edu", student_id: "S1001", major: "Computer Science", year: "Sophomore", gpa: 3.7, enrollment_date: "2024-08-15" },
  { id: 2, name: "Jane Doe", email: "jane.doe@university.edu", student_id: "S1002", major: "Biology", year: "Junior", gpa: 3.9, enrollment_date: "2023-08-15" },
  { id: 3, name: "Jim Doe", email: "jim.doe@university.edu", student_id: "S1003", major: "Mathematics", year: "Senior", gpa: 3.5, enrollment_date: "2022-08-15" },
];

app.get("/", (req: Request, res: Response) => {
  res.json({ message: "Here is the TypeScript Express API for the students database" });
});

app.get("/students", (req: Request, res: Response) => {
  res.json(students);
});

app.get("/students/:id", (req: Request, res: Response) => {
  const id = Number(req.params.id);
  const student = students.find((s) => s.id === id);
  if (!student) return res.status(404).json({ detail: "Student not found" });
  res.json(student);
});

app.post("/students", (req: Request, res: Response) => {
  const student: Student = req.body;
  if (students.some((s) => s.id === student.id)) {
    return res.status(400).json({ detail: "Student with this id already exists" });
  }
  students.push(student);
  res.status(201).json(student);
});

app.put("/students/:id", (req: Request, res: Response) => {
  const id = Number(req.params.id);
  const student: Student = req.body;
  const index = students.findIndex((s) => s.id === id);
  if (index === -1) return res.status(404).json({ detail: "Student not found" });
  if (student.id !== id) return res.status(400).json({ detail: "Student id in body must match path id" });
  students[index] = student;
  res.json(student);
});

app.delete("/students/:id", (req: Request, res: Response) => {
  const id = Number(req.params.id);
  const index = students.findIndex((s) => s.id === id);
  if (index === -1) return res.status(404).json({ detail: "Student not found" });
  const deleted = students.splice(index, 1)[0];
  res.json({ message: "Student deleted successfully", student: deleted });
});

app.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
});