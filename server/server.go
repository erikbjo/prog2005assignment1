package server

import (
	"log"
	"net/http"
	"os"
	"prog2005assignment1/server/handlers"
	"prog2005assignment1/server/shared"
)

// Start
/*
Start the server on the port specified in the environment variable PORT. If PORT is not set, the default port 8080 is used.
*/
func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Set up handler endpoints
	http.HandleFunc(shared.DefaultPath, handlers.DefaultHandler)
	http.HandleFunc(shared.StatusPath, handlers.StatusHandler)
	http.HandleFunc(shared.ReadershipPath, handlers.ReadershipHandler)
	http.HandleFunc(shared.BookCountPath, handlers.BookCountHandler)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
