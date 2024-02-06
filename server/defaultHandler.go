package server

import (
	"fmt"
	"net/http"
)

/*
Empty handler as default handler
*/
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")

	// Offer information for redirection to paths
	output := "This service does not provide any functionality on root path level. <br>" +
		"Please use paths <a href=\"" + READERSHIP_PATH + "\">" + READERSHIP_PATH + "</a> or <a href=\"" +
		BOOK_COUNT_PATH + "\">" + BOOK_COUNT_PATH + "</a>.<br>" +
		"Or check the <a href=\"" + STATUS_PATH + "\">" + STATUS_PATH + "</a> for server status."
	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)

	// Deal with error if any
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}

}
