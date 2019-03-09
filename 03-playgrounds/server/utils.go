package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"googlemaps.github.io/maps"
)

var (
	GEO_API_KEY = os.Getenv("GEO_API_KEY")
)

func getGeoCode() {
	// Use request to find geocode data
	c, err := maps.NewClient(maps.WithAPIKey(GEO_API_KEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	rawCode := &maps.GeocodingRequest{
		Address: "4132 Halifax Street, Vancouver, Canada",
	}

	geocode, err := c.Geocode(context.Background(), rawCode)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	var structData StructData
	for _, component := range geocode {

		structData = StructData{
			AddressComponent: component.AddressComponents,
			FormattedAddress: component.FormattedAddress,
			Geometry:         component.Geometry,
			PartialMatch:     component.PartialMatch,
			PlaceID:          component.PlaceID,
			PlusCode:         component.PlusCode,
			Types:            component.Types,
		}

	}
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "root"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO Playground(name, city) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	defer db.Close()
}
