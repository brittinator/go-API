// From https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/
// Also https://thenewstack.io/make-a-restful-json-api-go/
package main

import (
	"log"
	"net/http"
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

func main() {
	logger := NewLogger()
	router := NewRouter(logger)
	seedPeople()
	seedCars()

	// bind to port 12345 and pass router in
	log.Fatal(http.ListenAndServe(":12345", router))
}
