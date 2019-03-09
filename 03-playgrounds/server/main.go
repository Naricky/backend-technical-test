package main

import (
	"biba/backend-technical-test/03-playgrounds/server/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes for all the requests
	r.HandleFunc("/playground/GET", routes.GetPlayground)           // Retrieves playground information
	r.HandleFunc("/playground/UPDATE/:id", routes.UpdatePlayground) //Update specific Playground
	r.HandleFunc("/playground/CREATE", routes.CreatePlayground)     // Create playground
	r.HandleFunc("/playground/DELETE", routes.DeletePlayground)     // DELETE playground
	r.HandleFunc("/playground/LIST", routes.ListPlayground)
	r.HandleFunc("/weather/LIST", routes.ListWeather)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
