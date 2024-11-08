package handlers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"prog2005assignment1/server/shared"
	"time"
)

var client = &http.Client{
	Timeout: 3 * time.Second,
}

var StartTime = time.Now()

// StatusHandler
/*
Handle requests for /status
*/
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

/*
Handle GET request for /status
*/
func handleStatusGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	currentStatus := shared.Status{
		GutendexAPI:  getStatusCode(shared.CurrentGutendexApi, w),
		LanguageAPI:  getStatusCode(shared.LanguageApi, w),
		CountriesAPI: getStatusCode(shared.CurrentRestCountriesApi, w),
		Version:      shared.Version,
		Uptime:       math.Round(time.Since(StartTime).Seconds()),
	}

	marshaledStatus, err := json.MarshalIndent(currentStatus, "", "\t")
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

/*
Get status code from external API. Return 503 if API is not reachable.
*/
func getStatusCode(url string, w http.ResponseWriter) int {
	defer client.CloseIdleConnections()

	// Add language to language API, would get status 204 if not. Add /all to countries API, would get status 404 if not.
	if url == shared.LanguageApi {
		url = url + "/en"
	} else if url == shared.CurrentRestCountriesApi {
		url = url + "/all"
	}

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
