// contactinfo.go

package contactinfo

import "errors"

// ContactInfo represents the details like phone number, email, etc.
type ContactInfo struct {
	IsActive        bool   // Tracks if the contact info is active
	ContactInfoID   int    // Unique ID for each contact info
	ContactInfoType string // Type of contact info (e.g., "Phone", "Email")
	ContactInfoValue string // The actual value (e.g., "123-456-7890", "email@example.com")
}

// Constructor for creating new ContactInfo
func NewContactInfo(contactInfoID int, infoType, value string) (*ContactInfo, error) {
	// Validation of input
	if infoType == "" || value == "" {
		return nil, errors.New("contact info type or value cannot be empty")
	}

	// Create the ContactInfo object
	newContactInfo := &ContactInfo{
		IsActive:        true,
		ContactInfoID:   contactInfoID,
		ContactInfoType: infoType,
		ContactInfoValue: value,
	}

	return newContactInfo, nil
}

// Function to update contact info
func (ci *ContactInfo) UpdateContactInfo(parameter string, newValue interface{}) error {
	// Check if contact info is active
	if !ci.IsActive {
		return errors.New("cannot update inactive contact info")
	}

	switch parameter {
	case "type":
		// Update type of contact info
		if newValueStr, ok := newValue.(string); ok && newValueStr != "" {
			ci.ContactInfoType = newValueStr
		} else {
			return errors.New("invalid value for contact info type")
		}
	case "value":
		// Update value of contact info
		if newValueStr, ok := newValue.(string); ok && newValueStr != "" {
			ci.ContactInfoValue = newValueStr
		} else {
			return errors.New("invalid value for contact info value")
		}
	default:
		return errors.New("invalid parameter to update")
	}

	return nil
}

// Mark contact info as inactive (soft delete)
func (ci *ContactInfo) DeactivateContactInfo() {
	ci.IsActive = false
}
