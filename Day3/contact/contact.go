
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


func (contact *Contact) CreateContactInfo(infoType, value string) {
	infoID := 0

	if len(contact.ContactInfos) != 0 {
		infoID = contact.ContactInfos[len(contact.ContactInfos)-1].ContactInfoID
		infoID++
	}
	tempContactInfo := contactinfo.NewContactInfo(infoType, value, infoID)
	contact.ContactInfos = append(contact.ContactInfos, tempContactInfo)
}


func (contact *Contact) GetContactID() int {	return contact.ContactID }

func (contact *Contact) GetActivityStatus() bool { return contact.IsActive }
 
func (contact *Contact) GetContactInfo(infoID int) (*contactinfo.ContactInfo, error) {
	if !contact.IsActive {
		return nil, errors.New("contact is inactive")
	}
	for _, info := range contact.ContactInfos {
		if info.ContactInfoID == infoID && info.IsActive {
			return info, nil
		}
	}
	return nil, errors.New("contact info not found")
}

// UPDATION
func (contact *Contact) UpdateContact(parameter string, newValue interface{}) error {
	switch parameter {
	case "Firstname":
		if value, ok := newValue.(string); ok {
			if value == "" {
				return errors.New("first name cannot be empty")
			}
			contact.Firstname = value
			return nil
		}
		return errors.New("expected a string for first name")

	case "Lastname":
		if value, ok := newValue.(string); ok {
			if value == "" {
				return errors.New("last name cannot be empty")
			}
			contact.Lastname = value
			return nil
		}
		return errors.New("expected a string for last name")

	default:
		return errors.New("parameter not recognized")
	}
}

func (contact *Contact) UpdateContactInfo(infoID int, parameter string, newValue interface{}) error {
	for _, info := range contact.ContactInfos {
		if info.ContactInfoID == infoID && info.IsActive {
			return info.UpdateContactInfo(parameter, newValue)
		}
	}
	return errors.New("contact info not found")
}

// DELETION
func (contact *Contact) DeleteContactInfo(infoID int) error {
	if !contact.IsActive {
		return errors.New("contact is inactive")
	}
	for _, info := range contact.ContactInfos {
		if info.ContactInfoID == infoID && info.IsActive {
			info.IsActive = false
			return nil
		}
	}
	return errors.New("contact info not found")
}

func validateUserInfo(firstname, lastname string) error {
	if firstname == "" || lastname == "" {
		return errors.New("first name or last name cannot be empty")
	}
	return nil
}

