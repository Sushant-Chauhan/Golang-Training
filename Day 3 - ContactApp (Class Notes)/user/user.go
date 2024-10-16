
package user

import (
	"contactapp/contactapp"
	"contactapp/contactinfo"
	"errors"
	"fmt")

//user struct
type User struct {
	UserID  int
	Firstname string
	Lastname stringx
	IsAdmin bool
	IsActive bool
	Contacts []*contact.Contact
}

// ------------ Admin Features  - CRUD on users  -----------
// Create User: An admin can add new users to the system (Admin or Staff).
// Read Users: An admin can Read All Users in the system.
// Update User: An admin can Edit the details of any user (Can Edit- First Name, Last Name )
// Delete User: An admin can delete any user from the system. (IsActive = Flase)

//______ 1. CREATE - Admin , Staff ______  

// FACTORY FOR NEW ADMIN CREATION BY ADMIN
var addAdmin []*User
var userID = 0
func CreateNewAdmin(fname, lname string) *User{
	//validation
	if fname == "" || lname == "" {
		fmt.Println(("First name or last name cannot be empty"))
		return nil
	}

	var tempAdminObject = &User{
		UserID : userID,
		Firstname: fname,
		Lastname:  lname,
		IsAdmin:true,
		IsActive:true,
		Contact:nil
	}
	userID++
	allAdmin = append(allAdmin, tempAdminObject)
	return tempAdminObject,nil
}

 
// FACTORY FOR NEW STAFF CREATION BY ADMIN
func (u *User) NewStaff(fname, lname string) (*User, error) {
	//validate name
	if fname == "" || lname == "" {
		fmt.Println(("First name or last name cannot be empty"))
		return nil
	}
    //check is he admin (for staff creation) or acive
	if !u.IsAdmin || !u.IsActive {
		return nil, errors.New("only active Admins can create users")

	} else {  //we will create a new staff
		staffUser := &User{
			UserID:    userID,
			Firstname: fname,
			Lastname:  lname,
			IsAdmin:   false,
			IsActive:  true,
			Contacts:  nil,
		}
		userID++
		allStaff = append(allStaff, staffUser)
		return staffUser, nil
	}
}


// 2. ______ READ USERS by Admin ______
func (u *User) ReadUsers() ([]*User, error) {
		// code
}

//3. _____ UPDATE USER by Admin ______
func (u *User) UpdateUser(userID int, parameter string, newValue ) error {
   //code
   //- check if he admin or not - if yes then he can update otherwise not 
   //-find target user
   //-then using switch case - see what needs to be updated and update it (fname,lname,role change-like make him admin,isActive status change. Othewise default case)
}

//4. _______ DELETE USER by Admin ___________
func (u *User) DeleteUser(userID int)  { 
	//code
	//- check if he admin or not - only admin can delete user
}


// ------------ Staff features :  CRUD on Contact  &&  Contact Details ------------


// 1. CRUD on Contact Details:
// Staff Features:  Create, read, update, or delete details of contacts, & contact details like phone number or email.
// Create Contact: A staff user can add new contacts associated with their account.
// Read Contacts: A staff user can view their contacts.
// Update Contact: A staff user can edit their contacts.
// Delete Contact: A staff user can delete their contacts.
  
//1._________ Create Contact _________ 
func (u *User) CreateContact(firstname, lastname string) error { }


//2.________ Read Contacts _________
func (u *User) ReadContacts() ([]*contact.Contact, error) { }

//3. _________ Update Contact _________
func (u *User) UpdateContact(contactID int, parameter string, newValue ) error { }


//4. _________ Delete Contact _________
func (u *User) DeleteContact(contactID int) error { }



// 2. CRUD on Contact Details:
// Create Contact Details: Staff can add new contact details (such as address, additional phone numbers, or notes) to the contacts they manage.
// Read Contact Details: Staff can view the details of their contacts (email, phonenumber , address)
// Update Contact Details: Staff can edit the details of their contacts. For example, they can change the address, update an additional phone number, or modify 
// Delete Contact Details: Staff can remove contact details from their contacts. This action could apply to any part of the contact's details like deleting an old address or removing extra phone numbers.

//1._________ Create Contact Details _________ 
func (u *User) CreateContactInfo(contactID int, infoType, value string) error {

//2.________ Read Contacts  Details _________
func (u *User) ReadContactInfo(contactID int, infoID int) (*contactinfo.ContactInfo, error) {

//3. _________ Update Contact Details _________
func (u *User) UpdateContactInfo(contactID int, infoID int, parameter string, newValue interface{}) error {

//4. _________ Delete Contact Details _________
func (u *User) DeleteContactInfo(contactID int, infoID int, parameter string, newValue interface{}) error {




/*
//lecture
func (u *User) NewStaff(...){
	}
	userID++
	allStaff = append(addStaff, tempStaffObject)
	return tempStaffObject,nil

func (u *User) createContact(firstName, lastname string){
	defer func(){

	}
	//validations
	//here
	//contact factory call
	contactObj,err : contact.NewContact()
	if err!=nil{
		panic()
	}
	//append the contact object to u.contacts
}

func (u *User) UpdateContact(parameter string, 
	newValue interface{}, contactID int){
	//validation
	//call to the update method of contact struct
}

func (u *User) CreateContactInfo(contactID int, ContactInfoType string, 
	ContactInfoValue string) {
	//validations
	//find contact based on ID
	//call contact.CreateContactInfo()

	//call factory
	//insert contactInfo object into contact.contactInfos
}
*/
// I want to update contact details : 






