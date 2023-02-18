package university_search

import (
	"fmt"
	"net/http"
	"time"
)

// Global variable to hold the start time of the application
var startTime time.Time

func init() {
	// Initialize the start time of the application
	startTime = time.Now()
}

// Handler function for the diagnostics endpoint
func DiagHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetDiag(w)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

func handleGetDiag(w http.ResponseWriter) {

	// Calculate uptime of the application
	uptime := time.Since(startTime)

	// URLs to invoke APIs
	uniURL := "http://universities.hipolabs.com/"
	countryURL := "https://restcountries.com/"

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue the requests for UniversitiesAPI and CountriesAPI
	uniRes, err := client.Get(uniURL)
	if err != nil {
		fmt.Errorf("error in response from UniversitiesAPI: %s", err)
	}

	countryRes, err := client.Get(countryURL)
	if err != nil {
		fmt.Errorf("error in response from CountriesAPI: %s", err)
	}

	// Prepare return info with api-availability
	output := "Service Availability:" + LINEBREAK
	output += "---------------------" + LINEBREAK
	output += fmt.Sprintf("UniversityAPI: %s%s", uniRes.Status, LINEBREAK)
	output += fmt.Sprintf("CountriesAPI: %s%s", countryRes.Status, LINEBREAK)
	output += "Version: " + "v1" + LINEBREAK
	output += "Uptime: " + uptime.String() + LINEBREAK

	// Write the output to the response
	fmt.Fprint(w, output)
}
