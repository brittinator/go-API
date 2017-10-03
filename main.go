// From https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)


// Person is a person
type Person struct {
	// omitempty means if empty it's not included in the JSON
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address is an address struct
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// b/c no database yet
var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
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

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	// Decoding JSON and mapping it into the person variable
	json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["ID"]
	people = append(people, person)

	// encode all people
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)

  for i, item := range people {
    if item.ID == params["ID"]
    // recreate the slice of people w/out that person
    people = append(people[:i], people[i+1]...)
  }

  json.NewEncoder(w).Encode(people)
}

func main() {
	// mux.Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.
	// Requests can be matched based on URL host, path, path prefix, schemes, header and query values, HTTP methods or using custom matchers.
  router := mux.NewRouter()
  // Seed people
  people = append(people, Person{ID:"1", Firstname:"Julia", Lastname:"Childs", Address: &Address{City: "Spokane", State: "WA"}})
  people = append(people, Person{ID:"2", Firstname:"Jessica", Lastname:"Rabbit", Address: &Address{City: "Toon Town", State: "OO"}})
  people = append(people, Person{ID:"3", Firstname:"Ada", Lastname:"Lovelace", Address: &Address{City: "Somewhere", State: "WA"}})

  // these are the routes we handle with the method handlers
  router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":12345", router))
}
