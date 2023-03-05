package university_search

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*
Neighbourhandler Entry point for handler of NeighbourUni information
*/
func NeighbourHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetNeighbour(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			"' is supported.", http.StatusNotImplemented)
		return
	}
}

func handleGetNeighbour(w http.ResponseWriter, r *http.Request) {
	// Split the URL into parts delimited by "/"
	parts := strings.Split(r.URL.Path, "/")

	// Check that the URL contains the correct number of parts
	if len(parts) < 6 {
		http.Error(w, "URL does not contain all necessary parts", http.StatusBadRequest)
		log.Println("URL in request does not contain all necessary parts")
		return
	}

	// Get the limit parameter from the URL query parameters
	limitStr := r.URL.Query().Get("limit")
	limitLen, err := strconv.Atoi(limitStr)
	if err != nil && limitStr != "" {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		log.Println("Invalid limit parameter")
		return
	}

	limitBool := true // This is a boolean to check if the limit parameter is set
	// Checks if the limit parameter is not set
	if limitStr == "" {
		limitBool = false
	}

	// Get bordering countries data from "RESTcountries" API
	specCountryURL := COUNTRYAPI_NAME + parts[4] + "?fullText=true"
	countryResponse, err := http.Get(specCountryURL)
	if err != nil {
		http.Error(w, "Error during request to CountryAPI", http.StatusInternalServerError)
		log.Println("Failed to get bordering country data from CountryAPI")
		return
	}
	// Close the response body when the function returns
	defer countryResponse.Body.Close()

	// Struct to hold the response for the specified country
	var specCountryData []Country
	// Decode the response body into the struct
	err = json.NewDecoder(countryResponse.Body).Decode(&specCountryData)
	if err != nil {
		http.Error(w, "Error during request to CountryAPI", http.StatusInternalServerError)
		log.Println("Failed to decode country data from CountryAPI, about the specified country")
		return
	}

	// Get country data from "RESTcountries" API based on the list of border countries
	borders := specCountryData[0].Border
	countryData, err := getCountries(borders)
	if err != nil {
		http.Error(w, "Error during request to CountryAPI", http.StatusInternalServerError)
		log.Println("Failed to get country data from CountryAPI, about the bordering countries")
		return
	}

	// Get university data from UniversityAPI
	uniData, err := getUniversities(parts[5:])
	if err != nil {
		http.Error(w, "Error during request to UniversityAPI", http.StatusInternalServerError)
		log.Println("Failed to get university data from UniversityAPI")
		return
	}

	// Initialize a slice to hold all the response objects
	var response []Response
	limiter := 0
	for _, uni := range uniData {
		if limitBool && limiter >= limitLen {
			break
		}
		// Iterate through the country data to find the country that matches the university's country
		for _, country := range countryData {

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
				// Append the response object to the response slice
				response = append(response, responseObj)
				limiter++ // This is a limiter to limit the number of universities returned
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

	// Ensure that the server interprets requests as JSON from Client (browser)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshallResponse)

}
