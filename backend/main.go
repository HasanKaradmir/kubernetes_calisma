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

	mux := http.NewServeMux()

	mux.Handle("/backend", corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if the request method is POST
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		// decode the request body into a loginData struct
		var data loginData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		log.Printf("Request body: %+v", data)

		// construct the response object
		resp := response{Message: "Welcome, " + data.Name + " " + data.LastName + "!"}
		// encode the response as JSON and send it back to the client
		json.NewEncoder(w).Encode(resp)
		log.Printf("Response: %+v", resp)
	})))

	// start the server on localhost:80
	log.Fatal(http.ListenAndServe("0.0.0.0:444", mux))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
