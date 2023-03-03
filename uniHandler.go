package university_search

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

/*
Entry point for handler of UniHandler information
*/
func UniHandler(w http.ResponseWriter, r *http.Request) {

	// Ensure that request from client is a GET request
	switch r.Method {
	case http.MethodGet:
		handleGetUni(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

func handleGetUni(w http.ResponseWriter, r *http.Request) {
	// Ensure that the server interprets requests as JSON from Client (browser)
	w.Header().Set("content-type", "application/json")

	// Split the URL into parts delimited by "/"
	parts := strings.Split(r.URL.Path, "/")

	// Check that the URL contains the correct number of parts
	if len(parts) > 5 {
		http.Error(w, "Wrong URL", http.StatusBadRequest)
		log.Println("Wrong URL in request")
		return
	}

	// Get university data from "hipo" API
	uniData, err := getUniversities(parts[4:])
	if err != nil {
		http.Error(w, "Error during request to UniversityAPI", http.StatusInternalServerError)
		log.Println("Error during request")
		return
	}

	// Initialize a slice to hold all of the response objects
	var response []Response

	// Loop through each university in the response
	for _, uni := range uniData {
		// Get country data from "restcountries" API based on the university's ISO code
		countryData, err := getCountries(uni.IsoCode)
		if err != nil {
			http.Error(w, "Error during request to CountryAPI", http.StatusInternalServerError)
			log.Println("Error during request")
			return
		}

		// Create the response object for this university
		responseObj := Response{
			Name:      uni.Name,
			Country:   uni.Name,
			IsoCode:   uni.IsoCode,
			WebPages:  uni.WebPages,
			Languages: countryData[0].Languages,
			Maps:      countryData[0].Maps,
		}

		// Add the response object to the slice of responses
		response = append(response, responseObj)
	}

	// Marshal the response slice to JSON
	marshallResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error during formatting", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Write(marshallResponse)
	w.WriteHeader(http.StatusOK)
}

func getUniversities(name []string) ([]University, error) {
	encodedName := url.QueryEscape(strings.Join(name, " "))
	uniUrl := UNIVERSITYAPI + encodedName
	uniResponse, err := http.Get(uniUrl)
	if err != nil {
		return nil, err
	}
	// Close the response body after the function has returned
	defer uniResponse.Body.Close()
	// Decode the JSON response into a slice of "University" structs
	var uniData []University
	err = json.NewDecoder(uniResponse.Body).Decode(&uniData)
	if err != nil {
		return nil, err
	}

	// Check if any universities were found
	if len(uniData) == 0 {
		return nil, errors.New("No universities found")
	}

	return uniData, nil
}

func getCountries(isoCode string) ([]Country, error) {
	countryUrl := COUNTRYAPI + isoCode
	countryResponse, err := http.Get(countryUrl)
	if err != nil {
		return nil, err
	}
	// defer countryResponse.Body.Close()

	// Decode the JSON response into a slice of "Country" structs
	var countryData []Country
	err = json.NewDecoder(countryResponse.Body).Decode(&countryData)
	if err != nil {
		return nil, err
	}

	// Check if any countries were found
	if len(countryData) == 0 {
		return nil, errors.New("No countries found")
	}

	return countryData, nil
}
