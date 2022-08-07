package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Customer struct
// Phone number is a string to handle leading zeroes, etc.
// https://stackoverflow.com/questions/3483156/whats-the-right-way-to-represent-phone-numbers
// Keys are capitalized so json encoder can see them
type CustomerInfo struct {
	Id        uint
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

// Database Type
// The demo has a map as the database.
// Of course, this should change in a full implementation.
// ID needs to be unique, and is managed through map key
type CustomerDatabase = map[uint]CustomerInfo

// Global var to emulate the databse for now.
// Should move it to something more... sophisticated down the line.
var customerDatabase = CustomerDatabase{
	0: CustomerInfo{
		Id:        0,
		Name:      "Peppa Pig",
		Role:      "Cheeky Piggy",
		Email:     "peppa.pig@somewhere.in.uk",
		Phone:     "+44-00-98765-23",
		Contacted: false,
	},
	1: CustomerInfo{
		Id:        1,
		Name:      "Suzie Sheep",
		Role:      "Peppa's BFF",
		Email:     "suzie.sheep@somewhere.in.uk",
		Phone:     "+44-00-987432-23",
		Contacted: false,
	},
	2: CustomerInfo{
		Id:        2,
		Name:      "Mandy Mouse",
		Role:      "Peppa's playmate",
		Email:     "mandy.mouse@somewhere.in.uk",
		Phone:     "+44-00-98325-23",
		Contacted: true,
	},
}

// In a full implementation, this would return a handle to an
// external database. For now, a watered down version works.
func GetCustomerDatabase() *CustomerDatabase {
	db := CustomerDatabase{}

	return &db
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	j, err := json.Marshal(customerDatabase)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(j))
	}

	json.NewEncoder(w).Encode(customerDatabase)
}

func getSingleCustomer(w http.ResponseWriter, r *http.Request) {
	// Reading query parameters, as pointed in docs
	// https://github.com/gorilla/mux
	vars := mux.Vars(r)

	// https://stackoverflow.com/questions/35154875/convert-string-to-uint-in-go-lang
	u64, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	id := uint(u64)
	fmt.Println(id)

	w.WriteHeader(http.StatusOK)
	for i := range customerDatabase {
		if customerDatabase[i].Id == id {
			json.NewEncoder(w).Encode(customerDatabase[i])
			break
		}
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	// REST API implementation prefers JSON
	w.Header().Set("Content-Type", "application/json")

	var newEntry CustomerInfo

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)

	// If the entry already exists, flag a conflict.
	// Otherwise merge it in the current database.
	// Check if key already exists
	// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
	if _, ok := customerDatabase[newEntry.Id]; ok {
		w.WriteHeader(http.StatusConflict)
	} else {
		customerDatabase[newEntry.Id] = newEntry
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(customerDatabase)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {

}

func removeCustomer(w http.ResponseWriter, r *http.Request) {

}

func main() {
	/*
		Set up a router to handle the following:

		- Getting a single customer through a /customers/{id} path
		- Getting all customers through a the /customers path
		- Creating a customer through a /customers path
		- Updating a customer through a /customers/{id} path
		- Deleting a customer through a /customers/{id} path
	*/
	router := mux.NewRouter()
	router.HandleFunc("/customers/{id}", getSingleCustomer).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", createCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")
	router.HandleFunc("/customers/{id}", removeCustomer).Methods("DELETE")

	// Make it accessible at localhost:8000
	fmt.Println("Started at http://localhost:8000!")
	http.ListenAndServe(":8000", router)
}
