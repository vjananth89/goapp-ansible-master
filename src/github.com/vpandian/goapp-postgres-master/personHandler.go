package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define the structure in the form
type Person struct {
	Name       string `json:"name"`
	FavoriteColor   string `json:"color"`
	PetPref string `json:"pet_pref"`
}

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve people from postgresql database 
	personList, err := store.GetPerson()

	// Convert the `personList` variable to JSON
	personListBytes, err := json.Marshal(personList)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write JSON list of persons to response
	w.Write(personListBytes)
}

func createPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML form data received in the request
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Extract the field information about the person from the form info
	person := Person{}
	person.Name = r.Form.Get("name")
	person.FavoriteColor = r.Form.Get("color")
	person.PetPref = r.Form.Get("pet_pref")

	// enter the details in the db
	err = store.CreatePerson(&person)
	if err != nil {
		fmt.Println(err)
	}

	//Redirect to the originating HTML page
	http.Redirect(w, r, "/", http.StatusFound)
}
