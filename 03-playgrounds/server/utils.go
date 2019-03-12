package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"googlemaps.github.io/maps"
)

var (
	GEO_API_KEY = os.Getenv("GEO_API_KEY")
)

func getGeoCode(incomingRequest IncomingRequest) StructData {
	// Use request to find geocode data
	c, err := maps.NewClient(maps.WithAPIKey(GEO_API_KEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	rawCode := &maps.GeocodingRequest{
		Address: incomingRequest.Address,
	}

	geocode, err := c.Geocode(context.Background(), rawCode)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	geoData := StructData{
		Name:      incomingRequest.Name,
		Address:   geocode[0].FormattedAddress,
		Latitude:  int(geocode[0].Geometry.Viewport.SouthWest.Lat),
		Longitude: int(geocode[0].Geometry.Viewport.SouthWest.Lng),
	}

	return geoData
}

func dbConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/biba")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to DB!")
	return db
}
