package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Contact struct {
	Id    int    `json:"id"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

var contacts []Contact

func main() {
	SeedContacts()

	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/contacts", GetContacts).Methods("GET")
	router.HandleFunc("/contact/{id}", GetContact).Methods("GET")
	router.HandleFunc("/contact/create", CreateContact).Methods("POST")
	router.HandleFunc("/contact/delete/{id}", DeleteContact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func SeedContacts() {
	for i := 0; i < 20; i++ {
		newContact := Contact{
			Id:    i,
			Name:  fmt.Sprintf("Lucas %d", i),
			Email: fmt.Sprintf("lucas@eu%d.com", i),
		}

		contacts = append(contacts, newContact)
	}
}

func Home(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("Hello World!")
}

func GetContacts(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(contacts)
}

func GetContact(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	contactId, _ := strconv.Atoi(params["id"])

	for _, contact := range contacts {
		if contact.Id == contactId {
			json.NewEncoder(writer).Encode(contact)
			return
		}
	}
}

func CreateContact(writer http.ResponseWriter, request *http.Request) {
	var newContact Contact
	_ = json.NewDecoder(request.Body).Decode(&newContact)

	newContact.Id = len(contacts) + 1
	contacts = append(contacts, newContact)

	json.NewEncoder(writer).Encode(newContact)
}

func DeleteContact(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	contactId, _ := strconv.Atoi(params["id"])

	for index, contact := range contacts {
		if contact.Id == contactId {
			contacts = append(contacts[:index], contacts[index+1:]...)
			break
		}
	}

	json.NewEncoder(writer).Encode("Contact deleted with success")
}
