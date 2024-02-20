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

	/*
		return the number of potential readers for books in a given language, i.e., the population per country in which that language is official
	*/

	// Get two_letter_language_code from request
	cutQuery, _ := strings.CutPrefix(r.URL.Path, ReadershipPath)
	if len(cutQuery) != 2 || !unicode.IsLetter(rune(cutQuery[0])) || !unicode.IsLetter(rune(cutQuery[1])) {
		log.Println("Invalid request.")
		http.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}

	// Split and get first
	twoLetterLanguageCode := strings.Split(cutQuery, "/")[0]
	log.Println("two_letter_language_code: " + twoLetterLanguageCode)

	limit := r.URL.Query().Get("limit")
	if limit == "" {
		log.Println("No limit specified.")
	}

	// Get authors and books from bookCountHandler
	authors, books := GetAuthorsAndBooks(w, twoLetterLanguageCode)
	log.Println("Authors: " + strconv.Itoa(authors))
	log.Println("Books: " + strconv.Itoa(books))

	// Get country name from API
	countries := getCountriesWithLanguageWithTwoLetterLanguageCode(w, twoLetterLanguageCode)
	log.Println("Countries length: " + strconv.Itoa(len(countries)))

	// Loop through countries and get readership (inhabitants) from API
	var readerships []Readership

	for _, country := range countries {
		log.Println("Country: " + country.OfficialName)

		readership := Readership{
			Country:    country.OfficialName,
			Isocode:    country.Iso31661Alpha2,
			Books:      books,
			Authors:    authors,
			Readership: getReadership(w, country),
		}

		readerships = append(readerships, readership)
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
	response, err := client.Get(RestCountriesApi + "/alpha/" + country.Iso31661Alpha3)
	if err != nil {
		log.Println("Error when trying to get readership: " + err.Error())
		http.Error(w, "Error when trying to get readership", http.StatusInternalServerError)
	}

	log.Println("Response from countries API: " + response.Status)

	var countries []CountryFromRestCountries
	err = json.NewDecoder(response.Body).Decode(&countries)
	if err != nil {
		log.Println("Error when decoding JSON: " + err.Error())
		http.Error(w, "Error when decoding JSON", http.StatusInternalServerError)
	}

	if len(countries) > 1 {
		log.Println("\t--- More than one country found ---")
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

	log.Println("Response from language API: " + response.Status)
	log.Println(response.Body)

	var countries []Country
	err = json.NewDecoder(response.Body).Decode(&countries)
	if err != nil {
		log.Println("Error when decoding JSON: " + err.Error())
		http.Error(w, "Error when decoding JSON", http.StatusInternalServerError)
	}

	return countries
}
