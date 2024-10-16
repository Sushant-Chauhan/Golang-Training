// main.go

package main

import (
	"fmt"
	"contactapp/user"
)

func main() {
	// Create Admin user
	admin1 := user.CreateAdminUser("Sushant", "Chauhan")
	fmt.Println("Admin created:", admin1.Firstname, admin1.Lastname)

	// Create Staff users
	staff1 := admin1.CreateStaffUser("Varun", "Sharma")
	fmt.Println("Staff user created:", staff1.Firstname, staff1.Lastname)

	staff2 := admin1.CreateStaffUser("Arun", "Singh")
	fmt.Println("Staff user created:", staff2.Firstname, staff2.Lastname)

	// Create contacts for staff2
	fmt.Println("\nAdding contact info for staff2...")

	_, _ = staff2.CreateContactInfo(2, "email", "abcdef@gmail.com")
	_, _ = staff2.CreateContactInfo(2, "phone", "9282372838")
	_, _ = staff2.CreateContactInfo(2, "email", "pqrs@gmail.com")
	_, _ = staff2.CreateContactInfo(2, "phone", "9726152415")
	_, _ = staff2.CreateContactInfo(2, "email", "test123@gmail.com")
	_, _ = staff2.CreateContactInfo(2, "email", "contact1@gmail.com")
	_, _ = staff2.CreateContactInfo(2, "phone", "9172615442")

	// Print contact details for staff2
	fmt.Println("\nContact details for staff2:")
	staff2.PrintContactDetails()
}
