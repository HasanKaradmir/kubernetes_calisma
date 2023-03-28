package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type loginData struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

type response struct {
	Message string `json:"message"`
}

func main() {
	fmt.Println("Port listening on...")
	http.HandleFunc("/backend", func(w http.ResponseWriter, r *http.Request) {
		// check if the request method is POST
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// decode the request body into a loginData struct
		var data loginData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// construct the response object
		resp := response{Message: "Welcome, " + data.Name + " " + data.LastName + "!"}
		// encode the response as JSON and send it back to the client
		json.NewEncoder(w).Encode(resp)
	})

	// start the server on localhost:5500
	log.Fatal(http.ListenAndServe("127.0.0.1:5500", nil))
}
