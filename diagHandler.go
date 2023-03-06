package university_search

import (
	"encoding/json"
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
	uniURL := "http://universities.hipolabs.com/search?name=university"
	countryURL := "https://restcountries.com/v3.1/alpha/nor,fin,swe,rus"

	diag := Diagnosis{}

	// Issue the requests for UniversitiesAPI and CountriesAPI
	uniRes, err := http.Get(uniURL)
	if err != nil {
		diag.UniversitiesAPI = string(http.StatusServiceUnavailable)
	} else {
		diag.UniversitiesAPI = uniRes.Status
		defer uniRes.Body.Close()
	}

	countryRes, err := http.Get(countryURL)
	if err != nil {
		diag.CountriesAPI = string(http.StatusServiceUnavailable)
	} else {
		diag.CountriesAPI = countryRes.Status
		defer countryRes.Body.Close()
	}

	diag.Version = "v1"
	diag.Uptime = uptime.String()

	// Encode struct to JSON
	jsonBytes, err := json.Marshal(diag)
	if err != nil {
		http.Error(w, fmt.Sprintf("error in encoding to JSON: %s", err), http.StatusInternalServerError)
		return
	}

	// Write JSON to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
