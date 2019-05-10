package main

import (
	// The sql go library is needed to interact with the database
	"database/sql"
)

// Store will have two methods, to add a new person, and to get all existing people
type Store interface {
	CreatePerson(person *Person) error
	GetPerson() ([]*Person, error)
}

// `dbStore` struct implements the `Store` interface. Variable `db` takes the pointer
// to the SQL database connection object.
type dbStore struct {
	db *sql.DB
}

// Create a global `store` variable of type `Store` interface. It will be initialized
// in `func main()`.
var store Store

func (store *dbStore) CreatePerson(person *Person) error {
	// 'Person' is a struct which has "name", "color", and "pet_pref" attributes.
	// Type SQL query to insert new person into our database.
	// Note: `peopleinfo` is the name of the table within our `peopleDatabase` postgresql database.
	_, err := store.db.Query(
		"INSERT INTO person(name,color,pet_pref) VALUES ($1,$2,$3)",
		person.Name, person.FavoriteColor, person.PetPref)
	return err
}

func (store *dbStore) GetPerson() ([]*Person, error) {
	// Query the database for all persons, and return the result to the `rows` object.
	// Note: `peopleinfo` is the name of the table within our `peopleDatabase`
	rows, err := store.db.Query("SELECT name, color, pet_pref FROM person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	personList := []*Person{}
	for rows.Next() {
		// For each row returned from the database, create a pointer to a `Person` struct.
		person := &Person{}
		// Populate the `Name`, `FavoriteColor`, and `PetPref` attributes of the person
		if err := rows.Scan(&person.Name, &person.FavoriteColor, &person.PetPref); err != nil {
			return nil, err
		}
		// Finally, append the new person to the returned slice, and repeat for the next row
		personList = append(personList, person)
	}
	return personList, nil
}
