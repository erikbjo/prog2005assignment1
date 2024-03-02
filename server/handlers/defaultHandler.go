package handlers

import (
	"fmt"
	"log"
	"net/http"
	"prog2005assignment1/server/shared"
)

// DefaultHandler
/*
DefaultHandler is the default handler for the shared. It returns a message to the client, informing them that the server
does not provide any functionality on root path level.
*/
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")

	// Offer information for redirection to paths
	output := "This service does not provide any functionality on root path level. <br> Please use paths: " +
		"<ul><li><a href=\"" + shared.ReadershipPath + "\">" + shared.ReadershipPath + "</a></li>" +
		"<li><a href=\"" + shared.BookCountPath + "\">" + shared.BookCountPath + "</a></li>" +
		"<li><a href=\"" + shared.StatusPath + "\">" + shared.StatusPath + "</a></li></ul>"

	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)

	// Deal with error if any
	if err != nil {
		log.Println("Error when returning output: " + err.Error())
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}

}
