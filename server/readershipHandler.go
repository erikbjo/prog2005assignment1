package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

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

func handleReadershipGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	/*
		Pseude code:
		Get two_letter_language_code from request
		Get authors and books from bookCountHandler
		Get country name from API
		Get readership (inhabitants) from API
	*/

	// Get two_letter_language_code from request
	cutQuery, _ := strings.CutPrefix(r.URL.Path, ReadershipPath)

	// Split and get first
	twoLetterLanguageCode := strings.Split(cutQuery, "/")[0]

	if len(twoLetterLanguageCode) != 2 || !unicode.IsLetter(rune(twoLetterLanguageCode[0])) || !unicode.IsLetter(rune(twoLetterLanguageCode[1])) {
		log.Println("Invalid request.")
		http.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}

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
		log.Println("Limit: " + strconv.Itoa(limit))
	}

	// Get authors and books from bookCountHandler
	authors, books := GetAuthorsAndBooks(w, twoLetterLanguageCode)

	// Get country name from API
	countries := getCountriesWithLanguageWithTwoLetterLanguageCode(w, twoLetterLanguageCode)

	// Loop through countries and get readership (inhabitants) from API
	var readerships []Readership

	for i, country := range countries {
		// If limit is set and reached, break
		if limit > 0 && i >= limit {
			break
		}

		newReadership := Readership{
			Country:    country.OfficialName,
			Isocode:    country.Iso31661Alpha2,
			Books:      books,
			Authors:    authors,
			Readership: getReadership(w, country),
		}

		readerships = append(readerships, newReadership)
	}

	// Return JSON
	marshaledReaderships, err := json.MarshalIndent(readerships, "", "    ")
	if err != nil {
		log.Println("Error during JSON encoding: " + err.Error())
		http.Error(w, "Error during JSON encoding.", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(marshaledReaderships)
	if err != nil {
		log.Println("Failed to write response: " + err.Error())
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}

}

func getReadership(w http.ResponseWriter, country Country) int {
	defer client.CloseIdleConnections()
	response, err := client.Get(CurrentRestCountriesApi + "/alpha/" + country.Iso31661Alpha3)
	if err != nil {
		log.Println("Error when trying to get readership: " + err.Error())
		http.Error(w, "Error when trying to get readership", http.StatusInternalServerError)
	}

	var countries []CountryFromRestCountries
	err = json.NewDecoder(response.Body).Decode(&countries)
	if err != nil {
		log.Println("Error when decoding JSON: " + err.Error())
		http.Error(w, "Error when decoding JSON", http.StatusInternalServerError)
	}

	return countries[0].Population
}

func getCountriesWithLanguageWithTwoLetterLanguageCode(w http.ResponseWriter, code string) []Country {
	defer client.CloseIdleConnections()
	response, err := client.Get(LanguageApi + code)
	if err != nil {
		log.Println("Error when trying to get countries with language: " + err.Error())
		http.Error(w, "Error when trying to get countries with language", http.StatusInternalServerError)
	}

	var countries []Country
	err = json.NewDecoder(response.Body).Decode(&countries)
	if err != nil {
		log.Println("Error when decoding JSON: " + err.Error())
		http.Error(w, "Error when decoding JSON", http.StatusInternalServerError)
	}

	return countries
}
