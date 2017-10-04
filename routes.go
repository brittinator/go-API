package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// NewRouter creates a router for all routes, and wraps information about the route request up in a log format.
func NewRouter(logger *log.Logger) *mux.Router {
	// mux.Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.
	// Requests can be matched based on URL host, path, path prefix, schemes, header and query values, HTTP methods or using custom matchers.
	router := mux.NewRouter()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Log(logger, handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	// TODO: how to work in subrouter for /cars?
	return router
}

// handle car routes
// var carsSubrouter = router.PathPrefix("/cars").Subrouter()

// these are the routes we handle with the method handlers
// Routes are tested in the order they were added to the router, 1st one wins
// router.HandleFunc("/people", GetPeopleHandler).Methods("GET")
// router.HandleFunc("/people/{id}", GetPersonHandler).Methods("GET")
// router.HandleFunc("/people/{id}", CreatePersonHandler).Methods("POST")
// router.HandleFunc("/people/{id}", DeletePersonHandler).Methods("DELETE")

// carsSubrouter.HandleFunc("/", GetFleetHandler).Methods("GET")
// carsSubrouter.HandleFunc("/car/{id}", GetCarHandler).Methods("GET")
// carsSubrouter.HandleFunc("/car/{id}", CreateCarHandler).Methods("POST")
// carsSubrouter.HandleFunc("/car/{id}", DeleteCarHandler).Methods("DELETE")

// // bind to port 12345 and pass router in
// log.Fatal(http.ListenAndServe(":12345", router))

var routes = Routes{
	// Route{
	// 	"Index",
	// 	"GET",
	// 	"/",
	// 	Index,
	// },
	Route{
		"PeopleIndex",
		"GET",
		"/people",
		PeopleIndexHandler,
	},
	Route{
		"PeopleShow",
		"GET",
		"/people/{id}",
		PeopleShowHandler,
	},
	Route{
		"PeopleCreate",
		"POST",
		"/people/{id}",
		PeopleCreateHandler,
	},
	Route{
		"PeopleDelete",
		"DELETE",
		"/People/{id}",
		PeopleDeleteHandler,
	},
	Route{
		"CarIndex",
		"GET",
		"/car",
		CarIndexHandler,
	},
	Route{
		"PersonShow",
		"GET",
		"/car/{id}",
		CarShowHandler,
	},
	Route{
		"CarCreate",
		"POST",
		"/car/{id}",
		CarCreateHandler,
	},
	Route{
		"CarDelete",
		"DELETE",
		"/car/{id}",
		CarDeleteHandler,
	},
}
