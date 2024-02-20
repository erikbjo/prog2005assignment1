package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

var StartTime time.Time

var client = &http.Client{
	Timeout: 3 * time.Second,
}

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
	http.HandleFunc(DefaultPath, DefaultHandler)
	http.HandleFunc(StatusPath, StatusHandler)
	http.HandleFunc(ReadershipPath, ReadershipHandler)
	http.HandleFunc(BookCountPath, BookCountHandler)

	StartTime = time.Now()

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
