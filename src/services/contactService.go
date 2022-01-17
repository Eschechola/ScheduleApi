package services

import (
	"fmt"

	"github.com/eschechola/ScheduleApi/src/entities"
)

var contacts []entities.Contact

func SeedContacts() {
	for i := 0; i < 20; i++ {
		newContact := entities.Contact{
			Id:    i,
			Name:  fmt.Sprintf("Lucas %d", i),
			Email: fmt.Sprintf("lucas@eu%d.com", i),
		}

		contacts = append(contacts, newContact)
	}
}

func GetContacts() []entities.Contact {
	return contacts
}

func GetContact(id int) entities.Contact {
	for _, contact := range contacts {
		if contact.Id == id {
			return contact
		}
	}

	return entities.Contact{}
}

func CreateContact(contact entities.Contact) entities.Contact {
	var newContact = contact

	newContact.Id = len(contacts) + 1
	contacts = append(contacts, newContact)

	return newContact
}

func DeleteContact(id int) {
	for index, contact := range contacts {
		if contact.Id == id {
			contacts = append(contacts[:index], contacts[index+1:]...)
			break
		}
	}
}
