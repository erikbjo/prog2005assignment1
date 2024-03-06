package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"prog2005assignment1/server/shared"
	"testing"
)

func TestBookCountHandler(t *testing.T) {
	// Testing the book count handler, with a "fake" endpoint
	bookCountURL := "http://localhost:" + shared.DefaultPort + shared.BookCountPath + "?language=no,la,invalid"

	req, err := http.NewRequest("GET", bookCountURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleBookCountGetRequest)

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

	// Expecting an array of two bookCount structs
	// The first bookCount struct should have the language code "no" and the count above 15, authorCount above 10
	// The second bookCount struct should have the language code "la" and the count above 120, authorCount above 80
	// This is due to the Gutendex is ever-changing

	// Decode the JSON
	decoder := json.NewDecoder(rr.Body)
	var bookCounts []shared.BookCount
	err = decoder.Decode(&bookCounts)
	if err != nil {
		t.Errorf("Failed to decode JSON: %v", err)
	}

	// Test the fields
	if len(bookCounts) != 2 {
		t.Errorf("Expected 2 bookCount structs, got: %v", len(bookCounts))
	}

	// Test the first bookCount struct
	if bookCounts[0].Language != "no" {
		t.Errorf("Expected language code 'no', got: %v", bookCounts[0].Language)
	}

	if bookCounts[0].Books < 15 {
		t.Errorf("Expected count above 15, got: %v", bookCounts[0].Books)
	}

	if bookCounts[0].Authors < 10 {
		t.Errorf("Expected authorCount above 10, got: %v", bookCounts[0].Authors)
	}

	// Test the second bookCount struct
	if bookCounts[1].Language != "la" {
		t.Errorf("Expected language code 'la', got: %v", bookCounts[1].Language)
	}

	if bookCounts[1].Books < 120 {
		t.Errorf("Expected count above 120, got: %v", bookCounts[1].Books)
	}

	if bookCounts[1].Authors < 80 {
		t.Errorf("Expected authorCount above 80, got: %v", bookCounts[1].Authors)
	}

}
