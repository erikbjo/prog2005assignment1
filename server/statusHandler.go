package server

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"time"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleStatusGetRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' are supported.", http.StatusNotImplemented)
		return
	}

}

func handleStatusGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	currentStatus := Status{
		GutendexAPI:  getStatusCode(GUTENDEX_API_REMOTE),
		LanguageAPI:  getStatusCode(LANGUAGE_API_REMOTE),
		CountriesAPI: getStatusCode(COUNTRIES_API_REMOTE),
		Version:      "v1",
		Uptime:       math.Round(time.Since(StartTime).Seconds()),
	}

	marshaledStatus, err := json.MarshalIndent(currentStatus, "", "    ")
	if err != nil {
		log.Println("Error during JSON encoding: " + err.Error())
		http.Error(w, "Error during JSON encoding.", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(marshaledStatus)
	if err != nil {
		log.Println("Failed to write response: " + err.Error())
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func getStatusCode(url string) int {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error making request to external API: " + err.Error())
		return http.StatusInternalServerError
	}

	status := response.StatusCode

	err1 := response.Body.Close()
	if err1 != nil {
		log.Println("Failed to close response body: " + err1.Error())
	}

	return status
}
