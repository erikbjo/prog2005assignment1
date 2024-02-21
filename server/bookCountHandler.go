package server

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// BookCountHandler
/*
Handle requests for /bookCount
*/
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

/*
Handle GET request for /bookCount
*/
func handleBookCountGetRequest(w http.ResponseWriter, r *http.Request) {
	defer client.CloseIdleConnections()
	w.Header().Add("content-type", "application/json")

	// Uses /?language={:two_letter_language_code+}/
	languageQuery := r.URL.Query().Get("language")
	if languageQuery == "" {
		log.Println("No language specified.")
		http.Error(w, "No language specified.", http.StatusBadRequest)
		return
	}

	// Split languageQuery into individual languages
	languageQueries := strings.Split(languageQuery, ",")

	responses := make([]*http.Response, len(languageQueries))
	for i, language := range languageQueries {
		responses[i] = makeGutendexRequest(w, r, language)
	}

	totalBooks := getTotalBookCount(w)

	// Array of bookCount structs
	bookCounts := make([]BookCount, len(languageQueries))
	for i, res := range responses {
		mp := decodeJSON(w, res)

		if mp.Count == 0 {
			log.Println("No books found for language: " + languageQueries[i])
			continue
		}

		mp = recursiveGutendexRequest(w, mp)

		output := prettyPrintJSON(w, mp)

		booksOfLanguage := mp.Count
		fraction := float64(booksOfLanguage) / float64(totalBooks)
		uniqueAuthors := getUniqueAuthors(w, output)

		bookCount := BookCount{
			Language: languageQueries[i],
			Books:    booksOfLanguage,
			Authors:  uniqueAuthors,
			Fraction: fraction,
		}

		bookCounts[i] = bookCount
	}

	// Remove empty elements from bookCounts
	bookCounts = removeEmptyElements(bookCounts)

	// Marshal and write to response
	marshaledBookCount, err := json.MarshalIndent(bookCounts, "", "    ")
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

func recursiveGutendexRequest(w http.ResponseWriter, mp GutendexResult) GutendexResult {
	for mp.Next != "" {
		_, err := url.ParseRequestURI(mp.Next)
		if err != nil {
			log.Println("Invalid URL:", err.Error())
			http.Error(w, "Invalid URL", http.StatusInternalServerError)
			return mp
		}

		res, err := client.Get(mp.Next)
		if err != nil {
			log.Println("Error in response:", err.Error())
			http.Error(w, "Error in response", http.StatusInternalServerError)
			return mp
		}

		decoder := json.NewDecoder(res.Body)
		var newMp GutendexResult
		err = decoder.Decode(&newMp)
		if err != nil {
			log.Println("Error during decoding: " + err.Error())
			http.Error(w, "Error during decoding", http.StatusBadRequest)
		}

		mp.Results = append(mp.Results, newMp.Results...)
		mp.Next = newMp.Next
		mp.Previous = newMp.Previous

		log.Println("Current count: " + strconv.Itoa(len(mp.Results)))
	}

	return mp
}

/*
Remove empty elements from bookCounts, i.e. elements with language="", books=0, authors=0, fraction=0
*/
func removeEmptyElements(counts []BookCount) []BookCount {
	var newCounts []BookCount
	for _, count := range counts {
		if count.Language != "" && count.Books != 0 && count.Authors != 0 && count.Fraction != 0 {
			newCounts = append(newCounts, count)
		}
	}
	return newCounts
}

/*
Decode JSON and return as map
*/
func decodeJSON(w http.ResponseWriter, res *http.Response) GutendexResult {
	decoder := json.NewDecoder(res.Body)
	mp := GutendexResult{}

	err := decoder.Decode(&mp)
	if err != nil {
		log.Println("Error during decoding: " + err.Error())
		http.Error(w, "Error during decoding", http.StatusBadRequest)
	}

	return mp
}

/*
Pretty print JSON and return as byte array
*/
func prettyPrintJSON(w http.ResponseWriter, mp GutendexResult) []byte {
	output, err := json.MarshalIndent(mp, "", "  ")
	if err != nil {
		log.Println("Error during pretty printing: " + err.Error())
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return nil
	}

	return output
}

/*
Get total book count from Gutendex API
*/
func getTotalBookCount(w http.ResponseWriter) int {
	r, err1 := http.NewRequest(http.MethodGet, CurrentGutendexApi, nil)
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
	bookCount := mp.Count

	return bookCount
}

/*
Get unique authors from Gutendex API. Authors are distinguished by name and birth and death year.
*/
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

/*
Make request to Gutendex API. Takes languageQuery as parameter, which is a two-letter language code. Returns response.
*/
func makeGutendexRequest(w http.ResponseWriter, r *http.Request, languageQuery string) *http.Response {
	// Create new request
	r, err1 := http.NewRequest(http.MethodGet, CurrentGutendexApi+"?languages="+languageQuery, nil)
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

func GetAuthorsAndBooks(w http.ResponseWriter, twoLetterLanguageCode string) (int, int) {
	res := makeGutendexRequest(w, nil, twoLetterLanguageCode)
	mp := decodeJSON(w, res)

	mp = recursiveGutendexRequest(w, mp)

	output := prettyPrintJSON(w, mp)

	uniqueAuthors := getUniqueAuthors(w, output)

	return uniqueAuthors, mp.Count
}
