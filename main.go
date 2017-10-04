// From https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/
package main

import (
	"encoding/json"
	"log"
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

// Fleet b/c no database yet
var Fleet []Car

func GetPersonHandler(w http.ResponseWriter, req *http.Request) {
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

func GetPeopleHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	// Decoding JSON and mapping it into the person variable
	json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["ID"]
	people = append(people, person)

	// encode all people
	json.NewEncoder(w).Encode(people)
}

func DeletePersonHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for i, item := range people {
		if item.ID == params["ID"] {
			// recreate the slice of people w/out that person
			people = append(people[:i], people[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(people)
}

func seedPeople() {
	// Seed people
	people = append(people, Person{ID: "1", Firstname: "Julia", Lastname: "Childs", Address: &Address{City: "Spokane", State: "WA"}})
	people = append(people, Person{ID: "2", Firstname: "Jessica", Lastname: "Rabbit", Address: &Address{City: "Toon Town", State: "OO"}})
	people = append(people, Person{ID: "3", Firstname: "Ada", Lastname: "Lovelace", Address: &Address{City: "Somewhere", State: "WA"}})
}

func seedCars() {
	// Seed cars
	Fleet = append(Fleet, Car{ID: "1", Make: "Ford", Model: "Mustang"})
	Fleet = append(Fleet, Car{ID: "2", Make: "Hyundai", Model: "Veloster"})
	Fleet = append(Fleet, Car{ID: "3", Make: "Ford", Model: "Taurus"})
}

// Car is a Car
type Car struct {
	// omitempty means if empty it's not included in the JSON
	ID    string `json:"id,omitempty"`
	Make  string `json:"make,omitempty"`
	Model string `json:"model,omitempty"`
}

func GetCarHandler(w http.ResponseWriter, req *http.Request) {
	// grabbing params passed in with the request
	params := mux.Vars(req)
	// b/c there's no db
	for _, item := range Fleet {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetFleetHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Fleet)
}

func CreateCarHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var car Car
	// Decoding JSON and mapping it into the person variable
	json.NewDecoder(req.Body).Decode(&car)
	car.ID = params["ID"]
	Fleet = append(Fleet, car)

	// encode all people
	json.NewEncoder(w).Encode(people)
}

func DeleteCarHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for i, item := range Fleet {
		if item.ID == params["ID"] {
			// recreate the slice of Fleet w/out that person
			Fleet = append(Fleet[:i], Fleet[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(Fleet)
}

func main() {
	// mux.Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.
	// Requests can be matched based on URL host, path, path prefix, schemes, header and query values, HTTP methods or using custom matchers.
	router := mux.NewRouter()
	seedPeople()
	seedCars()

	// handle car routes
	carsSubrouter := router.PathPrefix("/cars").Subrouter()

	// these are the routes we handle with the method handlers
	// Routes are tested in the order they were added to the router, 1st one wins
	router.HandleFunc("/people", GetPeopleHandler).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonHandler).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonHandler).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonHandler).Methods("DELETE")

	carsSubrouter.HandleFunc("/", GetFleetHandler).Methods("GET")
	carsSubrouter.HandleFunc("/car/{id}", GetCarHandler).Methods("GET")
	carsSubrouter.HandleFunc("/car/{id}", CreateCarHandler).Methods("POST")
	carsSubrouter.HandleFunc("/car/{id}", DeleteCarHandler).Methods("DELETE")

	// bind to port 12345 and pass router in
	log.Fatal(http.ListenAndServe(":12345", router))
}
