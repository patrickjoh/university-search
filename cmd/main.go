package main

import (
	university_search "Assignment1"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = university_search.DEFAULT_PORT
	}

	// Set up handler endpoints
	http.HandleFunc(university_search.DEFAULT_PATH, university_search.DefaultHandler)
	http.HandleFunc(university_search.UNIINFO_PATH, university_search.UniHandler)
	http.HandleFunc(university_search.NEIGHBOUR_PATH, university_search.NeighbourHandler)
	http.HandleFunc(university_search.DIAG_PATH, university_search.DiagHandler)

	log.Println("Starting server on " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
