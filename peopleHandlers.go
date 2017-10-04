package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func PeopleShowHandler(w http.ResponseWriter, req *http.Request) {
	// grabbing params passed in with the request
	params := mux.Vars(req)
	// b/c there's no db
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func PeopleIndexHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func PeopleCreateHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	// Decoding JSON and mapping it into the person variable
	json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["ID"]
	people = append(people, person)

	// encode all people
	json.NewEncoder(w).Encode(people)
}

func PeopleDeleteHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for i, item := range people {
		if item.ID == params["ID"] {
			// recreate the slice of people w/out that person
			people = append(people[:i], people[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(people)
}
