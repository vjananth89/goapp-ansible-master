package main

import (
	"database/sql"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	// The `pq` package is a GO Postgres driver for the `database/sql` package.
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	// Declare a router
	r := mux.NewRouter()
	// Declare static file directory
	staticFileDirectory := http.Dir("./static/")
	// Create static file server for our static files, i.e., .html, .css, etc
	staticFileServer := http.FileServer(staticFileDirectory)
	// Create file handler
	staticFileHandler := http.StripPrefix("/", staticFileServer)
	// Add static file handler to our router
	r.Handle("/", staticFileHandler).Methods("GET")
	// Add handler for `get` and `post` people functions
	r.HandleFunc("/person", getPersonHandler).Methods("GET")
	r.HandleFunc("/person", createPersonHandler).Methods("POST")

	return r
}

func main() {
	// Setup connection to our postgresql database
	connString := `user=postgres
				   password=password
				   host=ec2-34-239-126-166.compute-1.amazonaws.com
				   port=5432
				   dbname=sampledb
				   sslmode=disable`
	db, err := sql.Open("postgres", connString)
	 if err != nil {
		panic(err)
		fmt.Println("Error")
	 }

	// Check if accessible by pinging
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// implement the `Store` interface. The `store` variable was
	store = &dbStore{db: db}

	// Create router
	r := newRouter()

	// Listen to the port
	http.ListenAndServe(":8080", r)
}
