package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	entities "github.com/eschechola/ScheduleApi/src/entities"
	contactService "github.com/eschechola/ScheduleApi/src/services"
	"github.com/gorilla/mux"
)

func main() {
	contactService.SeedContacts()

	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/contacts", GetContacts).Methods("GET")
	router.HandleFunc("/contact/{id}", GetContact).Methods("GET")
	router.HandleFunc("/contact/create", CreateContact).Methods("POST")
	router.HandleFunc("/contact/delete/{id}", DeleteContact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func Home(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("Hello World!")
}

func GetContacts(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(contactService.GetContacts())
}

func GetContact(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	contactId, _ := strconv.Atoi(params["id"])
	contact := contactService.GetContact(contactId)

	json.NewEncoder(writer).Encode(contact)
}

func CreateContact(writer http.ResponseWriter, request *http.Request) {
	var newContact entities.Contact
	_ = json.NewDecoder(request.Body).Decode(&newContact)

	contactCreated := contactService.CreateContact(newContact)

	json.NewEncoder(writer).Encode(contactCreated)
}

func DeleteContact(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	contactId, _ := strconv.Atoi(params["id"])

	contactService.DeleteContact(contactId)

	json.NewEncoder(writer).Encode("Contact deleted with success")
}
