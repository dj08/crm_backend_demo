package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Customer struct
type Customer struct {
	id        uint
	name      string
	role      string
	email     string
	phone     uint
	contacted bool
}

// The demo has a slice as the database.
// Of course, this should change in a full implementation.
type CustomerDatabase = []Customer

// Placeholder code to encapsulate database op down the line.
func GetDatabaseInstance() *CustomerDatabase {

}

// In a full implementation, this would return a handle to an
// external database. For now, a slice works.
func initializeDatabase() CustomerDatabase {

}

func getCustomers(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Create a mock database
	customerData := initializeDatabase()

	/*
		Set up a router to handle the following:

		- Getting a single customer through a /customers/{id} path
		- Getting all customers through a the /customers path
		- Creating a customer through a /customers path
		- Updating a customer through a /customers/{id} path
		- Deleting a customer through a /customers/{id} path
	*/
	router := mux.NewRouter()
	router.HandleFunc("/customers", getCustomers).Methods("GET")

	// Make it accessible at localhost:8000
	fmt.Println("Started at http://localhost:8000!")
	http.ListenAndServe(":8000", router)
}
