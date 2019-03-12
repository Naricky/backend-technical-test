package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// Routes for all the requests
	r.HandleFunc("/playground/GET", GetPlayground).Methods(http.MethodPost, http.MethodOptions)         // Get and create playground information
	r.HandleFunc("/playground/UPDATE", UpdatePlayground).Methods(http.MethodPost)    // Update specific Playground
	r.HandleFunc("/playground/DELETE/", DeletePlayground).Methods(http.MethodDelete) // Delete playground
	r.HandleFunc("/playground/LIST", ListPlayground).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/playground/CLOSEST", ClosestPlayground).Methods(http.MethodPost, http.MethodOptions)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
