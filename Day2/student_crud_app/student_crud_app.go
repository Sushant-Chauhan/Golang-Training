package student_crud_app

import (
	"errors"
	"fmt"
	"time"
)

// DOB
type DOB struct {
	Day   int
	Month int
	Year  int
}

// Student
type Student struct {
	RollNo           int
	FirstName        string
	LastName         string
	DateOfBirth      DOB
	SemesterCGPA     []float64
	YearOfEnrollment int
	YearOfPassing    int

	// Private derived attributes
	fullName            string
	age                 int
	finalCGPA           float64
	semesterGrades      []string
	finalGrade          string
	numberOfYearsToGrad int
}

// all students
var allStudents []Student

func NewStudent(rollNo int, firstName, lastName, dateOfBirth string, yearOfEnrollment, yearOfPassing int, semesterCGPA []float64) (*Student, error) {
	//  date of birth
	dobParts := parseDOB(dateOfBirth)
	dob := DOB{Day: dobParts[0], Month: dobParts[1], Year: dobParts[2]}

	// Validate the inputs
	if err := validateStudentInput(firstName, lastName, dob, yearOfEnrollment, yearOfPassing, semesterCGPA); err != nil {
		return nil, err
	}

	// Calculate derived attributes
	fullName := firstName + " " + lastName
	age := calculateAge(dob)
	finalCGPA := calculateFinalCGPA(semesterCGPA)
	semesterGrades := calculateSemesterGrades(semesterCGPA)
	finalGrade := calculateFinalGrade(finalCGPA)
	numberOfYearsToGrad := yearOfPassing - yearOfEnrollment

	student := &Student{
		RollNo:              rollNo,
		FirstName:           firstName,
		LastName:            lastName,
		DateOfBirth:         dob,
		SemesterCGPA:        semesterCGPA,
		YearOfEnrollment:    yearOfEnrollment,
		YearOfPassing:       yearOfPassing,
		fullName:            fullName,
		age:                 age,
		finalCGPA:           finalCGPA,
		semesterGrades:      semesterGrades,
		finalGrade:          finalGrade,
		numberOfYearsToGrad: numberOfYearsToGrad,
	}

	allStudents = append(allStudents, *student) // Adding studentt to allStudents slice
	return student, nil
}

// dob string parse
func parseDOB(dobStr string) []int {
	var day, month, year int
	fmt.Sscanf(dobStr, "%d-%d-%d", &day, &month, &year)
	return []int{day, month, year}
}

// age based on DOB
func calculateAge(dob DOB) int {
	currentYear := time.Now().Year()
	age := currentYear - dob.Year

	if time.Now().Month() < time.Month(dob.Month) || (time.Now().Month() == time.Month(dob.Month) && time.Now().Day() < dob.Day) {
		age--
	}
	return age
}

// final CGPA
func calculateFinalCGPA(cgpa []float64) float64 {
	var total float64
	for _, grade := range cgpa {
		total += grade
	}
	return total / float64(len(cgpa))
}

// semester grades
func calculateSemesterGrades(cgpa []float64) []string {
	var grades []string
	for _, gpa := range cgpa {
		if gpa >= 9 {
			grades = append(grades, "A+")
		} else if gpa >= 8 {
			grades = append(grades, "A")
		} else if gpa >= 7 {
			grades = append(grades, "B")
		} else {
			grades = append(grades, "C")
		}
	}
	return grades
}

// final grade
func calculateFinalGrade(finalCGPA float64) string {
	if finalCGPA >= 9 {
		return "A+"
	} else if finalCGPA >= 8 {
		return "A"
	} else if finalCGPA >= 7 {
		return "B"
	}
	return "C"
}

// Validate student input
func validateStudentInput(firstName, lastName string, dob DOB, yearOfEnrollment, yearOfPassing int, semesterCGPA []float64) error {
	if err := validateName(firstName); err != nil {
		return err
	}
	if err := validateName(lastName); err != nil {
		return err
	}
	if err := validateDOB(dob); err != nil {
		return err
	}
	if yearOfEnrollment >= yearOfPassing {
		return errors.New("Year of Enrollment must be less than Year of Passing")
	}
	if err := validateSemesterCGPA(semesterCGPA); err != nil {
		return err
	}
	return nil
}

// Validate name
func validateName(name string) error {
	if len(name) == 0 || !isAlphabetic(name) {
		return errors.New("Name must be non-empty and contain only alphabets")
	}
	return nil
}

// Check if string contains only alphabets
func isAlphabetic(s string) bool {
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return false
		}
	}
	return true
}

// Validate date of birth
func validateDOB(dob DOB) error {
	age := calculateAge(dob)
	if age < 15 {
		return errors.New("Age must be greater than or equal to 15 for graduate students")
	}
	return nil
}

// Validate semester CGPA
func validateSemesterCGPA(cgpas []float64) error {
	for _, cgpa := range cgpas {
		if cgpa < 0 || cgpa > 10 {
			return errors.New("CGPA must be between 0 and 10")
		}
	}
	return nil
}

// Get all students
func GetAllStudents() []Student {
	return allStudents
}

// Update student information based on RollNo
func UpdateStudent(rollNo int, field, newValue string) error {
	for i, student := range allStudents {
		if student.RollNo == rollNo {
			switch field {
			case "firstName":
				student.FirstName = newValue
			case "lastName":
				student.LastName = newValue
			default:
				return errors.New("Invalid field for update")
			}
			allStudents[i] = student                                      // Update the student in the slice
			student.fullName = student.FirstName + " " + student.LastName // Update fullName
			return nil
		}
	}
	return errors.New("Student not found")
}

// Delete a student based on RollNo
func DeleteStudent(rollNo int) error {
	for i, student := range allStudents {
		if student.RollNo == rollNo {
			allStudents = append(allStudents[:i], allStudents[i+1:]...) // Remove the student from the slice
			return nil
		}
	}
	return errors.New("Student not found")
}

// Delete all students
func DeleteAllStudents() {
	allStudents = []Student{} // Clear the student slice
}

// Getter methods for the derived attributes
func (s *Student) FullName() string {
	return s.fullName
}

func (s *Student) Age() int {
	return s.age
}

func (s *Student) FinalCGPA() float64 {
	return s.finalCGPA
}
