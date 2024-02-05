package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
Entry point handler for Location information
*/
func LocationHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetRequest(w, r)
	case http.MethodPost:
		handlePostRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' and '"+http.MethodPost+"' are supported.", http.StatusNotImplemented)
		return
	}

}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	loc := Location{Name: "Gjøvik", Postcode: "2815", Geolocation: Coordinates{Latitude: 12.5, Longitude: 56.4}}

	w.Header().Add("content-type", "application/json")

	// Encode JSON
	encoder := json.NewEncoder(w)
	err := encoder.Encode(loc)
	if err != nil {
		http.Error(w, "Error during JSON encoding.", http.StatusInternalServerError)
	}
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	location := Location{}

	err1 := decoder.Decode(&location)
	if err1 != nil {
		log.Println("Failed to decode request body: " + err1.Error())
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	fmt.Println(location)

	if location.Name == "" {
		log.Println("Error while processing name of received location. Name is empty.")
		http.Error(w, "Error while processing name", http.StatusBadRequest)
		return
	}
}
