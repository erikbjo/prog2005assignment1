package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

// ReadershipHandler
/*
Handle requests for /readership, only GET requests are supported.
*/
func ReadershipHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleReadershipGetRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			" is supported.", http.StatusNotImplemented)
		return
	}
}

/*
Handle GET request for /readership
*/
func handleReadershipGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	// Get two_letter_language_code from request, cut off .../readership/
	cutQuery := strings.TrimPrefix(r.URL.Path, ReadershipPath)

	// cutQuery should now be {two_letter_language_code}/... OR {two_letter_language_code}...
	// Split by / and take first part
	twoLetterLanguageCode := strings.Split(cutQuery, "/")[0]

	// Check if twoLetterLanguageCode is valid, i.e. two letters
	if len(twoLetterLanguageCode) != 2 || !unicode.IsLetter(rune(twoLetterLanguageCode[0])) || !unicode.IsLetter(rune(twoLetterLanguageCode[1])) {
		log.Println("Invalid request. Invalid language code.")
		http.Error(w, "Invalid request. Invalid language code.", http.StatusBadRequest)
		return
	}

	// Get limit from request, if not set, limit is 0. Has to be a positive integer.
	var limit int
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limit = 0
	} else {
		err := error(nil)
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			log.Println("Invalid limit specified.")
			http.Error(w, "Invalid limit specified. Please specify a positive integer.", http.StatusBadRequest)
			return
		}
	}

	// Get authors and books from bookCountHandler
	authors, books := GetAuthorsAndBooks(w, twoLetterLanguageCode)

	// Get countries with language with two letter language code
	countries := getCountriesWithLanguageWithTwoLetterLanguageCode(w, twoLetterLanguageCode)
	if countries == nil {
		log.Println("Error when trying to get countries with language.")
		// Error message already sent
		return
	} else if len(countries) == 0 {
		log.Println("No countries found with language: " + twoLetterLanguageCode)
		http.Error(w, "No countries found with language.", http.StatusNotFound)
		return
	}

	// Loop through countries and get readership (inhabitants) from API
	var readerships []Readership
	for i, country := range countries {
		// If limit is set and reached, break
		if limit > 0 && i >= limit {
			break
		}

		// Create new readership struct
		newReadership := Readership{
			Country:    country.OfficialName,
			Isocode:    country.Iso31661Alpha2,
			Books:      books,
			Authors:    authors,
			Readership: getReadership(w, country),
		}

		// Append to readerships
		readerships = append(readerships, newReadership)
	}

	// Return JSON
	marshaledReaderships, err := json.MarshalIndent(readerships, "", "\t")
	if err != nil {
		log.Println("Error during JSON encoding: " + err.Error())
		http.Error(w, "Error during JSON encoding.", http.StatusInternalServerError)
		return
	}

	// Write to response
	_, err = w.Write(marshaledReaderships)
	if err != nil {
		log.Println("Failed to write response: " + err.Error())
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return // Not necessary, but for clarity
	}
}

/*
Get population of a country from RestCountries API
*/
func getReadership(w http.ResponseWriter, country Country) int {
	defer client.CloseIdleConnections()

	// Get response from RestCountries API
	response, err := client.Get(CurrentRestCountriesApi + "/alpha/" + country.Iso31661Alpha3)
	if err != nil {
		log.Println("Error when trying to get readership: " + err.Error())
		http.Error(w, "Error when trying to get readership", http.StatusInternalServerError)
	}

	// Decode JSON, could return multiple countries, but since we use alpha3 code, in reality we only get one
	var countries []CountryFromRestCountries
	err = json.NewDecoder(response.Body).Decode(&countries)
	if err != nil {
		log.Println("Error when decoding JSON: " + err.Error())
		http.Error(w, "Error when decoding JSON", http.StatusInternalServerError)
	}

	// Assume we only get one country since alpha3 codes are unique, return population of first country
	return countries[0].Population
}

/*
Get all countries that uses a given language by two letter language code
*/
func getCountriesWithLanguageWithTwoLetterLanguageCode(w http.ResponseWriter, code string) []Country {
	defer client.CloseIdleConnections()

	// Get response from Language2Countries API
	response, err := client.Get(LanguageApi + code)
	if err != nil {
		log.Println("Error when trying to get countries with language: " + err.Error())
		http.Error(w, "Error when trying to get countries with language", http.StatusServiceUnavailable)
		return nil
	}

	// Decode JSON
	var countries []Country
	err = json.NewDecoder(response.Body).Decode(&countries)
	if err != nil {
		log.Println("Error when decoding JSON: " + err.Error())
		http.Error(w, "Error when decoding JSON", http.StatusInternalServerError)
		return nil
	}

	return countries
}
