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
UniHandler Entry point for handler
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
	if len(parts) < 5 {
		http.Error(w, "URL does not contain all necessary parts", http.StatusBadRequest)
		log.Println("URL in request does not contain all necessary parts")
		return
	}

	// Get university data from "hipo" API
	uniData, err := getUniversities(parts[4:])
	if err != nil {
		http.Error(w, "Error during request to UniversityAPI", http.StatusInternalServerError)
		log.Println("Error during request to UniversityAPI")
		return
	}

	// Initialize a slice to hold all ISO codes
	var isocode []string

	// Loop through each university in the response
	for _, uni := range uniData {
		// Check if the iso code is already in the slice
		found := false
		// Loop through each iso code in the slice
		for _, code := range isocode {
			// If the iso code is found, break out of the loop
			if uni.Alpha2Code == code {
				found = true
				break
			}
		}
		// If the iso code is not found, append it to the slice
		if !found {
			isocode = append(isocode, uni.Alpha2Code)
		}
	}

	// Get country data from "restcountries" API based on the university's ISO code
	countryData, err := getCountries(isocode)
	if err != nil {
		http.Error(w, "Error during request to CountryAPI", http.StatusInternalServerError)
		return
	}

	// Initialize a slice to hold all response objects
	var response []Response

	// Loop through each university in the response
	for _, uni := range uniData {
		// Loop through each country in the response
		for _, country := range countryData {
			// Find the country that matches the current university
			if uni.Alpha2Code == country.Alpha2Code {
				// Create the response object for this university
				responseObj := Response{
					Name:      uni.Name,
					Country:   uni.Country,
					IsoCode:   uni.Alpha2Code,
					WebPages:  uni.WebPages,
					Languages: country.Languages,
					Maps:      country.Maps,
				}
				// Add the response object to the slice of responses
				response = append(response, responseObj)
				// Break out of the loop
				break
			}
		}
	}

	// Marshal the response slice to JSON
	marshallResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error during formatting", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Write(marshallResponse)
}

func getUniversities(name []string) ([]University, error) {
	encodedName := url.QueryEscape(strings.Join(name, " "))
	uniUrl := UNIVERSITYAPI + encodedName
	// Get the response from the API
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

func getCountries(isoCode []string) ([]Country, error) {

	countryUrl := COUNTRYAPI_CODES
	// Loop through each ISO code and append the code the URL
	// Append each code to the URL with a comma delimiter
	isoCodesStr := strings.Join(isoCode, ",")
	countryUrl += isoCodesStr

	countryResponse, err := http.Get(countryUrl)
	if err != nil {
		return nil, err
	}
	defer countryResponse.Body.Close()

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
