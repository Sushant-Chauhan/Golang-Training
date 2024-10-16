package contact
import "contactapp/contactInfo"

type Contact struct {
	ContactID int
	Firstname string
	Lastname string
	IsActive  bool  
	ContactInfos []*contactinfo.ContactInfo
}

function NewContact(ContactID int, 
	Firstname string, 
	Lastname string, 
	IsActive bool,
	ContactInfos []*contactinfo.ContactInfo) (*Contact, error){
		//here
}

function UpdateContact(){
	//validation and switch
}