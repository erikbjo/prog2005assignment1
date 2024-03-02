package util

import (
	"log"
	"net/http"
	"prog2005assignment1/server/shared"
	"time"
	"unicode"
)

var client = &http.Client{
	Timeout: 1 * time.Second,
}

// LanguageCodeChecker
/*
Check if the language code is valid. Returns a boolean indicating if the language code is valid.
Can also return an error to the client if there's an error with the request to the external API.
*/
func LanguageCodeChecker(languageCode string, responseWriter http.ResponseWriter) bool {
	if len(languageCode) != 2 || !unicode.IsLetter(rune(languageCode[0])) || !unicode.IsLetter(rune(languageCode[1])) {
		log.Println("Invalid request. Invalid language code.")
		return false
	}

	// Make request to Language2Country API, if 204 is returned, the language is not valid
	res, err := client.Get(shared.LanguageApi + "/" + languageCode)
	if err != nil {
		log.Println("Error when checking Language2Country API:", err.Error())
		http.Error(responseWriter, "Error when checking external API", http.StatusServiceUnavailable)
		return false
	}
	if res.StatusCode == 204 {
		// Service is available, but the language code is not valid
		return false
	} else if res.StatusCode == 200 {
		// Service is available and the language code is valid
		return true
	}

	// If the status code is not 204 or 200, return false
	return false
}
