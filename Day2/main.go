package main

import (
	"fmt"
	"log"
	"projectcrud/student_crud_app"
)

func main() {
	student1, err := student_crud_app.NewStudent(1, "Sushant", "Chauhan", "24-02-2001", 2020, 2024, []float64{9.0, 8.5, 9.2, 9.0, 8.5, 9.2, 8.4, 9.2})
	if err != nil {
		log.Fatalf("Error creating student 1: %v\n", err)
	}
	fmt.Printf("Created Student: %s, Age: %d, Final CGPA: %.2f\n", student1.FullName(), student1.Age(), student1.FinalCGPA())

	student2, err := student_crud_app.NewStudent(2, "Priya", "Sharma", "15-05-2000", 2019, 2023, []float64{8.5, 9.0, 9.5})
	if err != nil {
		log.Fatalf("Error creating student 2: %v\n", err)
	}
	fmt.Printf("Created Student: %s, Age: %d, Final CGPA: %.2f\n", student2.FullName(), student2.Age(), student2.FinalCGPA())

	// Read all students
	allStudents := student_crud_app.GetAllStudents()
	fmt.Println("\nAll Students:")
	for _, s := range allStudents {
		fmt.Printf("RollNo: %d, Name: %s, Age: %d, Final CGPA: %.2f\n", s.RollNo, s.FullName(), s.Age(), s.FinalCGPA())
	}

	// Update a student
	if err := student_crud_app.UpdateStudent(1, "firstName", "Sushant Kumar"); err != nil {
		log.Fatalf("Error updating student: %v\n", err)
	}

	// Read updated students
	allStudents = student_crud_app.GetAllStudents()
	fmt.Println("\nAfter Update:")
	for _, s := range allStudents {
		fmt.Printf("RollNo: %d, Name: %s, Age: %d, Final CGPA: %.2f\n", s.RollNo, s.FullName(), s.Age(), s.FinalCGPA())
	}

	// Delete student by RollNo
	if err := student_crud_app.DeleteStudent(2); err != nil {
		log.Fatalf("Error deleting student: %v\n", err)
	}

	// Read all students after deletion
	allStudents = student_crud_app.GetAllStudents()
	fmt.Println("\nAfter Deletion:")
	for _, s := range allStudents {
		fmt.Printf("RollNo: %d, Name: %s, Age: %d, Final CGPA: %.2f\n", s.RollNo, s.FullName(), s.Age(), s.FinalCGPA())
	}

	// Delete all students
	student_crud_app.DeleteAllStudents() // New method to delete all students
	fmt.Println("\nAll Students Deleted. Current List:")
	allStudents = student_crud_app.GetAllStudents()
	if len(allStudents) == 0 {
		fmt.Println("No students left.")
	}
}
