
/*
Contact

Attributes: 
Relationships: Each contact belongs to a user, and each contact can have multiple contact details.
Features :  CRUD on Contact and Contact Details
*/ 


package contact

import (
	"contactapp/contactinfo"
	"errors"
	"fmt"
)

// Contact represents an individual contact, which may have multiple contact details
type Contact struct {
	ContactID    int                        
	Firstname    string                     
	Lastname     string                     
	IsActive     bool                       
	ContactInfos []*contactinfo.ContactInfo // List of associated contact details
}

// Constructor for creating a new contact
func NewContact(contactID int, firstname, lastname string, isActive bool, contactInfos []*contactinfo.ContactInfo) (*Contact, error) {
	// Validation
	if firstname == "" || lastname == "" {
		return nil, errors.New("first name and last name cannot be empty")
	}

	newContact := &Contact{
		ContactID:    contactID,
		Firstname:    firstname,
		Lastname:     lastname,
		IsActive:     isActive,
		ContactInfos: contactInfos,
	}

	return newContact, nil
}

// Update the contact's details - first name, last name, or IsActive status
func (c *Contact) UpdateContact(parameter string, newValue interface{}) error {
 	if !c.IsActive {
		return errors.New("cannot update inactive contact")
	}

	// Update based on the parameter
	switch parameter {
	case "firstname":
		if newFirstname, ok := newValue.(string); ok && newFirstname != "" {
			c.Firstname = newFirstname
		} else {
			return errors.New("invalid value for first name")
		}
	case "lastname":
		if newLastname, ok := newValue.(string); ok && newLastname != "" {
			c.Lastname = newLastname
		} else {
			return errors.New("invalid value for last name")
		}
	case "isActive":
		if newIsActive, ok := newValue.(bool); ok {
			c.IsActive = newIsActive
		} else {
			return errors.New("invalid value for isActive status")
		}
	default:
		return errors.New("invalid parameter to update")
	}

	return nil
}

// Deactivate a contact (soft delete)
func (c *Contact) DeactivateContact() {
	c.IsActive = false
	for _, info := range c.ContactInfos {
		info.DeactivateContactInfo() // Deactivate all associated contact details
	}
}

// Add new contact info to the contact
func (c *Contact) AddContactInfo(info *contactinfo.ContactInfo) error {
	if !c.IsActive {
		return errors.New("cannot add contact info to inactive contact")
	}

	// Append the new contact info to the list
	c.ContactInfos = append(c.ContactInfos, info)
	return nil
}

// Remove (deactivate) a contact info from the contact
func (c *Contact) RemoveContactInfo(contactInfoID int) error {
	for _, info := range c.ContactInfos {
		if info.ContactInfoID == contactInfoID && info.IsActive {
			info.DeactivateContactInfo()
			return nil
		}
	}

	return errors.New("contact info not found or already inactive")
}

// Print contact details  
func (c *Contact) PrintContactDetails() {
	fmt.Printf("Contact ID: %d\nName: %s %s\nActive: %t\n", c.ContactID, c.Firstname, c.Lastname, c.IsActive)
	for _, info := range c.ContactInfos {
		fmt.Printf("  Contact Info: [%d] %s: %s (Active: %t)\n", info.ContactInfoID, info.ContactInfoType, info.ContactInfoValue, info.IsActive)
	}
}
