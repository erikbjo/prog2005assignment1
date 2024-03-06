package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"prog2005assignment1/server/shared"
	"reflect"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	statusURL := "http://localhost:" + shared.DefaultPort + shared.StatusPath

	// Test the status handler
	// Expect a 200 OK response, with a JSON body
	// The JSON body should contain the status of the Gutendex API, Language2Country API, RestCountries API,
	// the version of the server and the uptime of the server

	req, err := http.NewRequest("GET", statusURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleStatusGetRequest)

	handler.ServeHTTP(rr, req)

	// Test the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Test the content type
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("content-type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expectedContentType)
	}

	// Test the JSON body
	// Cant really test the remote APIs, but we can test the structure of the JSON, and the value types
	// The JSON should contain the following fields: GutendexAPI, LanguageAPI, CountriesAPI, Version, Uptime
	// The GutendexAPI, LanguageAPI and CountriesAPI should be integers
	// The Version should be a string
	// The Uptime should be a float64

	// Decode the JSON
	decoder := json.NewDecoder(rr.Body)
	var status shared.Status
	err = decoder.Decode(&status)
	if err != nil {
		t.Errorf("Failed to decode JSON: %v", err)
	}

	// Test the fields
	if reflect.TypeOf(status.GutendexAPI).Kind() != reflect.Int {
		t.Errorf("GutendexAPI is not an integer: got %v", status.GutendexAPI)
	}

	if reflect.TypeOf(status.LanguageAPI).Kind() != reflect.Int {
		t.Errorf("LanguageAPI is not an integer: got %v", status.LanguageAPI)
	}

	if reflect.TypeOf(status.CountriesAPI).Kind() != reflect.Int {
		t.Errorf("CountriesAPI is not an integer: got %v", status.CountriesAPI)
	}

	if reflect.TypeOf(status.Version).Kind() != reflect.String {
		t.Errorf("Version is not a string: got %v", status.Version)
	} else if status.Version != shared.Version {
		t.Errorf("Version is not the expected version: got %v", status.Version)
	}

	if reflect.TypeOf(status.Uptime).Kind() != reflect.Float64 {
		t.Errorf("Uptime is not a float64: got %v", status.Uptime)
	}
}
