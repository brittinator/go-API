package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CarShowHandler(w http.ResponseWriter, req *http.Request) {
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

func CarIndexHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Fleet)
}

func CarCreateHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var car Car
	// Decoding JSON and mapping it into the person variable
	json.NewDecoder(req.Body).Decode(&car)
	car.ID = params["ID"]
	Fleet = append(Fleet, car)

	// encode all people
	json.NewEncoder(w).Encode(people)
}

func CarDeleteHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for i, item := range Fleet {
		if item.ID == params["ID"] {
			// recreate the slice of Fleet w/out that person
			Fleet = append(Fleet[:i], Fleet[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(Fleet)
}
