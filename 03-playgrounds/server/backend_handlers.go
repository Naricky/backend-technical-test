package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	id        int
	name      string
	address   string
	latitude  int
	longitude int
)

func GetPlayground(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("this is request", r.Body)
	// When request to get information about playground comes in, first double check if it exist on database
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("this is body",body)

	var incomingRequest IncomingRequest

	json.Unmarshal(body, &incomingRequest)

	db := dbConn()

	rows, err := db.Query(`
	Select * FROM data
	`)
	
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Getting existingData from mySql
	var existingData []StructData
	var singularData StructData
	for rows.Next() {

		err := rows.Scan(&singularData.Id, &singularData.Name, &singularData.Address, &singularData.Latitude, &singularData.Longitude)
		if err != nil {
			log.Fatal(err)
		}
		existingData = append(existingData, singularData)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	//Comparing to request Name or Address. If exist on db, return. If not exist, append then return
	existingChecker := 0

	for _, data := range existingData {
		if data.Address == incomingRequest.Address || data.Name == incomingRequest.Name {
			existingChecker++
		}
	}

	geoData := getGeoCode(incomingRequest)

	if existingChecker == 0 {

		insertDB, err := db.Query("INSERT INTO data(name,address,latitude,longitude) VALUES(?,?,?,?)", geoData.Name, geoData.Address, geoData.Latitude, geoData.Longitude)
		if err != nil {
			log.Fatal(err)
		}
		defer insertDB.Close()
	}
	PayloadInByte, err := json.Marshal(geoData)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(PayloadInByte)
	
	defer db.Close()

}

// To update playground, I was not sure which specifics params : For sake of the assignment, I've hardcoded the paramters for id = 1
func UpdatePlayground(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	sqlStatement := `
	UPDATE data
	SET name = ?
	WHERE id = ?;`
	_, err := db.Exec(sqlStatement, "ChangedName!", 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated Playground info!")
	http.Redirect(w, r, "/", 301)

}

// To delete playground, I was not sure which specific params : For sake of the assignment, I've hardcoded the parameters for id = 1
func DeletePlayground(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	sqlStatement := `
	DELETE FROM data
	WHERE id = ?;`
	_, err := db.Exec(sqlStatement, 1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Deleted Playground info!")
		http.Redirect(w, r, "/", 301)
	}
}

// Request to receive all data from db
func ListPlayground(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	db := dbConn()

	rows, err := db.Query(`
	Select * FROM data
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Getting existingData from mySql
	var existingData []StructData
	var singularData StructData
	for rows.Next() {

		err := rows.Scan(&singularData.Id, &singularData.Name, &singularData.Address, &singularData.Latitude, &singularData.Longitude)
		if err != nil {
			log.Fatal(err)
		}
		existingData = append(existingData, singularData)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	PayloadInByte, err := json.Marshal(existingData)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(PayloadInByte)

}

// Get Closest playground in db comparing to geocodes given by incoming request
func ClosestPlayground(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body, err := ioutil.ReadAll(r.Body)
	var incomingRequest IncomingRequest

	json.Unmarshal(body, &incomingRequest)
	db := dbConn()
	geoData := getGeoCode(incomingRequest)
	rows, err := db.Query(`
	SELECT *
	FROM data
	ORDER BY ((latitude-?)*(latitude-?)) + ((longitude - ?)*(longitude - ?)) ASC
	`, geoData.Latitude, geoData.Latitude, geoData.Longitude, geoData.Longitude)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// Getting existingData from mySql
	var existingData []StructData
	var singularData StructData
	for rows.Next() {

		err := rows.Scan(&singularData.Id, &singularData.Name, &singularData.Address, &singularData.Latitude, &singularData.Longitude)
		if err != nil {
			log.Fatal(err)
		}
		existingData = append(existingData, singularData)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	PayloadInByte, err := json.Marshal(existingData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(PayloadInByte)

}
