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
		GutendexAPI:  getStatusCode(GUTENDEX_API_REMOTE, w),
		LanguageAPI:  getStatusCode(LANGUAGE_API_REMOTE, w),
		CountriesAPI: getStatusCode(COUNTRIES_API_REMOTE, w),
		Version:      VERSION,
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

func getStatusCode(url string, w http.ResponseWriter) int {
	defer client.CloseIdleConnections()
	response, err := client.Get(url)
	if err != nil {
		log.Println("Error making request to external API: " + err.Error())
		return http.StatusServiceUnavailable
	}

	err2 := response.Body.Close()
	if err2 != nil {
		log.Println("Error closing response body: " + err2.Error())
		return http.StatusInternalServerError
	}

	return response.StatusCode
}
