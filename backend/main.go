package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// User struct for holding name and last name
type User struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

// Database connection string
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "your-dbname"
)

// InsertUser function inserts the given user to the database
func InsertUser(user User) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := `INSERT INTO users (name, last_name) VALUES ($1, $2)`

	_, err = db.Exec(sqlStatement, user.Name, user.LastName)
	if err != nil {
		return err
	}

	return nil
}

// LoginHandler handles the login request
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the user to the database
	err = InsertUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User added successfully")
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Register the login route
	r.HandleFunc("/login", LoginHandler).Methods("POST")

	// Start the server
	http.ListenAndServe(":8080", r)
}
