package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func BookCountHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleBookCountGetRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' are supported.", http.StatusNotImplemented)
		return
	}

}

func handleBookCountGetRequest(w http.ResponseWriter, r *http.Request) {
	defer client.CloseIdleConnections()
	w.Header().Add("content-type", "application/json")

	// Uses /?language={:two_letter_language_code+}/
	languageQuery := r.URL.Query().Get("language")
	res := makeGutendexRequest(w, r, languageQuery)

	mp := decodeJSON(w, res)
	output := prettyPrintJSON(w, mp)

	totalBooks := getTotalBookCount(w)
	booksOfLanguage := int(mp["count"].(float64))
	fraction := float64(booksOfLanguage) / float64(totalBooks)
	uniqueAuthors := getUniqueAuthors(w, output)

	bookCount := BookCount{
		Language: languageQuery,
		Books:    booksOfLanguage,
		Authors:  uniqueAuthors,
		Fraction: fraction,
	}

	// Marshal and write to response
	marshaledBookCount, err := json.MarshalIndent(bookCount, "", "    ")
	if err != nil {
		log.Println("Error during JSON encoding: " + err.Error())
		http.Error(w, "Error during JSON encoding.", http.StatusInternalServerError)
		return
	}

	_, err6 := w.Write(marshaledBookCount)
	if err6 != nil {
		log.Println("Failed to write response: " + err6.Error())
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func decodeJSON(w http.ResponseWriter, res *http.Response) map[string]interface{} {
	decoder := json.NewDecoder(res.Body)
	mp := map[string]interface{}{}

	err := decoder.Decode(&mp)
	if err != nil {
		log.Println("Error during decoding: " + err.Error())
		http.Error(w, "Error during decoding", http.StatusBadRequest)
		return nil
	}

	return mp
}

func prettyPrintJSON(w http.ResponseWriter, mp map[string]interface{}) []byte {
	output, err := json.MarshalIndent(mp, "", "  ")
	if err != nil {
		log.Println("Error during pretty printing: " + err.Error())
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return nil
	}

	return output
}

func getTotalBookCount(w http.ResponseWriter) int {
	r, err1 := http.NewRequest(http.MethodGet, GUTENDEX_API_REMOTE, nil)
	if err1 != nil {
		log.Println("Error in creating request:", err1.Error())
		http.Error(w, "Error in creating request", http.StatusInternalServerError)
	}

	r.Header.Add("content-type", "application/json")
	res, err2 := client.Do(r)
	if err2 != nil {
		log.Println("Error in response:", err2.Error())
		http.Error(w, "Error in response", http.StatusInternalServerError)
	}

	mp := decodeJSON(w, res)
	bookCount := int(mp["count"].(float64))

	return bookCount
}

func getUniqueAuthors(w http.ResponseWriter, output []byte) int {
	var result GutendexResult
	err := json.Unmarshal(output, &result)
	if err != nil {
		log.Println("Error during JSON decoding: " + err.Error())
		http.Error(w, "Error during JSON decoding", http.StatusInternalServerError)
	}

	uniqueAuthors := make(map[string]bool)
	for _, book := range result.Results {
		for _, author := range book.Authors {
			// Also using birth and death year to distinguish between authors with the same name
			uniqueAuthors[author.Name+strconv.Itoa(author.BirthYear)+strconv.Itoa(author.DeathYear)] = true
		}
	}

	return len(uniqueAuthors)
}

func makeGutendexRequest(w http.ResponseWriter, r *http.Request, languageQuery string) *http.Response {
	// Create new request
	r, err1 := http.NewRequest(http.MethodGet, GUTENDEX_API_REMOTE+"?languages="+languageQuery, nil)
	if err1 != nil {
		log.Println("Error in creating request:", err1.Error())
		http.Error(w, "Error in creating request", http.StatusInternalServerError)
	}

	// Setting content type -> effect depends on the service provider
	r.Header.Add("content-type", "application/json")

	// Issue request
	res, err2 := client.Do(r)
	if err2 != nil {
		log.Println("Error in response:", err2.Error())
		http.Error(w, "Error in response", http.StatusInternalServerError)
	}

	return res
}
