package contactinfo

import (
	"errors"
)

// CONTACTINFORMATION STRUCTURE
type ContactInformation struct {
	ContactInfoID    int
	ContactInfoType  string
	ContactInfoValue string
	IsActive         bool
}

// GET CONTACT INFO ID
func (contactinfo *ContactInformation) GetContactInfoID() int {
	return contactinfo.ContactInfoID
}

// GET CONTACT INFO STATUS
func (contactinfo *ContactInformation) GetContactInfoStatus() bool {
	return contactinfo.IsActive
}

// CREATE NEW CONTACT INFORMATION
func CreateContactInfo(contactInfoType, contactInfoValue string, contactInfoID int) (*ContactInformation, error) {
	err := validateContactInfo(contactInfoType, contactInfoValue)
	if err != nil {
		return nil, err
	}
	tempContactInfo := &ContactInformation{
		ContactInfoID:    contactInfoID,
		ContactInfoType:  contactInfoType,
		ContactInfoValue: contactInfoValue,
		IsActive:         true,
	}
	return tempContactInfo, nil
}

// READ CONTACT INFORMATION
func GetContactInfo(contactInfoID int, contactInfos []*ContactInformation) (*ContactInformation, error) {
	for _, contactInfo := range contactInfos {
		if contactInfo.ContactInfoID == contactInfoID && contactInfo.IsActive {
			return contactInfo, nil
		}
	}
	return nil, errors.New("no such contact information id found")
}

// UPDATE CONTACT INFORMATION
func (contactInfo *ContactInformation) UpdateContactInfo(parameter string, newValue interface{}) error {
	switch parameter {
	case "Contact Information Type":
		return contactInfo.updateContactInfoType(newValue)
	case "Contact Information Value":
		return contactInfo.updateContactInfoValue(newValue)
	default:
		return errors.New("no such parameter found")
	}
}

// UPDATE CONTACT INFO TYPE
func (contactInfo *ContactInformation) updateContactInfoType(newValue interface{}) error {
	if value, ok := newValue.(string); ok {
		if value == "" {
			return errors.New("contact information type cannot be empty")
		}
		if value == "phone" && len(contactInfo.ContactInfoValue) != 10 {
			return errors.New("phone number must be 10 digits")
		}
		contactInfo.ContactInfoType = value
		return nil
	}
	return errors.New("invalid contact information type, expected a string")
}

// UPDATE CONTACT INFO VALUE
func (contactInfo *ContactInformation) updateContactInfoValue(newValue interface{}) error {
	if value, ok := newValue.(string); ok {
		if value == "" {
			return errors.New("contact information value cannot be empty")
		}
		if contactInfo.ContactInfoType == "phone" && len(value) != 10 {
			return errors.New("phone number must be 10 digits")
		}
		contactInfo.ContactInfoValue = value
		return nil
	}
	return errors.New("invalid contact information value, expected a string")
}

// VALIDATE CONTACT INFORMATION
func validateContactInfo(contactInfoType, contactInfoValue string) error {
	if contactInfoType == "" || contactInfoValue == "" {
		return errors.New("contact information type and value cannot be empty")
	}
	if contactInfoType == "phone" && len(contactInfoValue) != 10 {
		return errors.New("phone number must be 10 digits")
	}
	return nil
}
