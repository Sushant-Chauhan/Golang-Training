/* 
User

Attributes:  UserID, Firstname , Lastname, IsAdmin, IsActive , Contacts []*contact.Contact
Roles: Admin, Staff
Relationships: A user can have multiple contacts, and contacts can have multiple contact details.
Features :  CRUD on users [ IsAdmin: true]

Admin Features:
Create User: An admin can add new users to the system (Admin or Staff).
Read Users: An admin can Read All Users in the system.
Update User: An admin can Edit the details of any user (Can Edit- First Name, Last Name )
Delete User: An admin can delete any user from the system. (IsActive = Flase)


Staff Features:  Create, read, update, or delete details of contacts, & contact details like phone number or email.
Create Contact: A staff user can add new contacts associated with their account.
Read Contacts: A staff user can view their contacts.
Update Contact: A staff user can edit their contacts.
Delete Contact: A staff user can delete their contacts.
  
*/

 

package user

import (
	"contactapp/contactapp"
	"contactapp/contactinfo"
	"errors"
	"fmt")

//user struct
type User struct {
	UserID    int
	Firstname string
	Lastname  string
	IsAdmin   bool
	IsActive  bool
	Contacts  []*contact.Contact
}
// ------------ Admin Features  - CRUD on users  -----------
// Create User: An admin can add new users to the system (Admin or Staff).
// Read Users: An admin can Read All Users in the system.
// Update User: An admin can Edit the details of any user (Can Edit- First Name, Last Name )
// Delete User: An admin can delete any user from the system. (IsActive = Flase)

//______ 1. CREATE - Admin , Staff ______  

// FACTORY FOR NEW ADMIN CREATION  - by Admin another admin
var allUsers []*User   //admin and staff both
var userID = 0

func CreateNewAdmin(fname, lname string) (*User, error) {
	// Validation
	if fname == "" || lname == "" {
		return nil, errors.New("first name or last name cannot be empty")
	}

	var adminUser = &User{
		UserID : userID,
		Firstname: fname,
		Lastname:  lname,
		IsAdmin:true,
		IsActive:true,
		Contact:nil
	}
	userID++
	// allAdmin = append(allAdmin, adminUser)
	allUsers = append(allUsers, adminUser)

	return adminUser, nil
}

 
// FACTORY FOR NEW STAFF CREATION BY ADMIN
// var addAdmin []*User
func (u *User) CreateNewStaff(fname, lname string) (*User, error) {
	if fname == "" || lname == "" {
		return nil, errors.New("first name or last name cannot be empty")
	}

    //check is he admin (for staff creation) or acive
	if !u.IsAdmin || !u.IsActive {
		return nil, errors.New("only active Admins can create new users")
	}

	//we will create a new staff
	staffUser := &User{
		UserID:    userID,
		Firstname: fname,
		Lastname:  lname,
		IsAdmin:   false,
		IsActive:  true,
		Contacts:  nil,
	}

	userID++
	allUsers = append(allUsers, staffUser)
	return staffUser, nil
	
}


// 2. ______ READ USERS by Admin ______
func (u *User) ReadUsers() ([]*User, error) {
	if !u.IsAdmin || !u.IsActive {
		return nil, errors.New("only active Admins can read all users")
	}
	return allUsers, nil
}

//3. _____ UPDATE USER by Admin ______
func (u *User) UpdateUser(targetUserID int, field, newValue string) error {
	//code
   //- check if he admin or not - if yes then he can update otherwise not 
   //-find target user
   //-then using switch case - see what needs to be updated and update it (fname,lname,role change-like make him admin,isActive status change. Othewise default case)
    if !u.IsAdmin || !u.IsActive {
		return errors.New("only active Admins can update users")
	}

	// Find user to update
	var targetUser *User
	for _, usr := range allUsers {
		if usr.UserID == targetUserID {
			targetUser = usr
			break
		}
	}

	if targetUser == nil {
		return errors.New("user not found")
	}

	// Update based on field
	switch field {
	case "Firstname":
		targetUser.Firstname = newValue
	case "Lastname":
		targetUser.Lastname = newValue
	case "IsActive":
		if newValue == "true" {
			targetUser.IsActive = true
		} else {
			targetUser.IsActive = false
		}
	default:
		return errors.New("invalid field to update")
	}
	return nil
}

//4. _______ DELETE USER by Admin ___________
func (u *User) DeleteUser(targetUserID int) error {
	//code
	//- check if he admin or not - only admin can delete user
	if !u.IsAdmin || !u.IsActive {
		return errors.New("only active Admins can delete users")
	}

	// Find user to deactivate (soft delete)
	for _, usr := range allUsers {
		if usr.UserID == targetUserID {
			usr.IsActive = false // soft delete by deactivating user
			return nil
		}
	}
	return errors.New("user not found")
}


// ------------ Staff features :  CRUD on Contact  &&  Contact Details ------------


// 1. CRUD on Contact Details:
// Staff Features:  Create, read, update, or delete details of contacts, & contact details like phone number or email.
// Create Contact: A staff user can add new contacts associated with their account.
// Read Contacts: A staff user can view their contacts.
// Update Contact: A staff user can edit their contacts.
// Delete Contact: A staff user can delete their contacts.
  
//1._________ Create Contact _________ 
func (u *User) CreateContact(firstname, lastname string) error {
	if !u.IsActive {
		return errors.New("only active users can create contacts")
	}
	contact := &contact.Contact{
		ContactID: userID,
		Firstname: firstname,
		Lastname:  lastname,
		IsActive:  true,
	}
	u.Contacts = append(u.Contacts, contact)
	return nil
}

//2.________ Read Contacts _________
func (u *User) ReadContacts() ([]*contact.Contact, error) {
	if !u.IsActive {
		return nil, errors.New("only active users can read contacts")
	}
	return u.Contacts, nil
}


//3. _________ Update Contact _________
func (u *User) UpdateContact(contactID int, field, newValue string) error {
	if !u.IsActive {
		return errors.New("only active users can update contacts")
	}

	// Find contact to update
	var targetContact *contact.Contact
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			targetContact = c
			break
		}
	}

	if targetContact == nil {
		return errors.New("contact not found")
	}

	// Update based on field
	switch field {
	case "Firstname":
		targetContact.Firstname = newValue
	case "Lastname":
		targetContact.Lastname = newValue
	case "IsActive":
		if newValue == "true" {
			targetContact.IsActive = true
		} else {
			targetContact.IsActive = false
		}
	default:
		return errors.New("invalid field to update")
	}
	return nil
}


//4. _________ Delete Contact _________
func (u *User) DeleteContact(contactID int) error {
	if !u.IsActive {
		return errors.New("only active users can delete contacts")
	}

	// Soft delete by deactivating contact
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			c.IsActive = false
			return nil
		}
	}
	return errors.New("contact not found")
}



// 2. CRUD on Contact Details:
// Create Contact Details: Staff can add new contact details (such as address, additional phone numbers, or notes) to the contacts they manage.
// Read Contact Details: Staff can view the details of their contacts (email, phonenumber , address)
// Update Contact Details: Staff can edit the details of their contacts. For example, they can change the address, update an additional phone number, or modify 
// Delete Contact Details: Staff can remove contact details from their contacts. This action could apply to any part of the contact's details like deleting an old address or removing extra phone numbers.

// //1._________ Create Contact Details _________ 
// func (u *User) CreateContactInfo(contactID int, infoType, value string) error {

// //2.________ Read Contacts  Details _________
// func (u *User) ReadContactInfo(contactID int, infoID int) (*contactinfo.ContactInfo, error) {

// //3. _________ Update Contact Details _________
// func (u *User) UpdateContactInfo(contactID int, infoID int, parameter string, newValue interface{}) error {

// //4. _________ Delete Contact Details _________
// func (u *User) DeleteContactInfo(contactID int, infoID int, parameter string, newValue interface{}) error {

// In-memory storage for contact details, using a counter for unique IDs
var contactInfoIDCounter int = 1


// 1. Create Contact Details
func (u *User) CreateContactInfo(contactID int, infoType, value string) error {
	// Validate input
	if infoType == "" || value == "" {
		return errors.New("contact info type or value cannot be empty")
	}

	var targetContact *contact.Contact
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			targetContact = c
			break
		}
	}

	if targetContact == nil {
		return errors.New("contact not found")
	}

	// Create new contact info
	newContactInfo := &contactinfo.ContactInfo{
		ContactInfoID:   contactInfoIDCounter,
		ContactInfoType: infoType,
		ContactInfoValue: value,
		IsActive:        true,
	}

	contactInfoIDCounter++

	targetContact.ContactInfos = append(targetContact.ContactInfos, newContactInfo)

	return nil
}

// 2. Read Contact Details
func (u *User) ReadContactInfo(contactID int, infoID int) (*contactinfo.ContactInfo, error) {
	var targetContact *contact.Contact
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			targetContact = c
			break
		}
	}

	if targetContact == nil {
		return nil, errors.New("contact not found")
	}

	for _, info := range targetContact.ContactInfos {
		if info.ContactInfoID == infoID && info.IsActive {
			return info, nil
		}
	}

	return nil, errors.New("contact info not found or inactive")
}

// 3. Update Contact Details
func (u *User) UpdateContactInfo(contactID int, infoID int, parameter string, newValue interface{}) error {
	var targetContact *contact.Contact
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			targetContact = c
			break
		}
	}

	if targetContact == nil {
		return errors.New("contact not found")
	}

	for _, info := range targetContact.ContactInfos {
		if info.ContactInfoID == infoID && info.IsActive {
			switch parameter {
			case "type":
				if newValueStr, ok := newValue.(string); ok && newValueStr != "" {
					info.ContactInfoType = newValueStr
				} else {
					return errors.New("invalid value for contact info type")
				}
			case "value":
				if newValueStr, ok := newValue.(string); ok && newValueStr != "" {
					info.ContactInfoValue = newValueStr
				} else {
					return errors.New("invalid value for contact info value")
				}
			default:
				return errors.New("invalid parameter to update")
			}
			return nil
		}
	}

	return errors.New("contact info not found or inactive")
}

// 4. Delete Contact Details
func (u *User) DeleteContactInfo(contactID int, infoID int) error {
	var targetContact *contact.Contact
	for _, c := range u.Contacts {
		if c.ContactID == contactID {
			targetContact = c
			break
		}
	}

	if targetContact == nil {
		return errors.New("contact not found")
	}

	// Find the contact info by ID and set its IsActive flag to false
	for _, info := range targetContact.ContactInfos {
		if info.ContactInfoID == infoID && info.IsActive {
			info.IsActive = false
			return nil
		}
	}

	return errors.New("contact info not found or already inactive")
}
