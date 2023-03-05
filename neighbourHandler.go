package university_search

import (
	"net/http"
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
	/*
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
		_, err := strconv.Atoi(limitStr)
		if err != nil && limitStr != "" {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			log.Println("Invalid limit parameter")
			return
		}

		// Get bordering countries data from "RESTcountries" API
		specCountryURL := COUNTRYAPI_NAME + parts[4] + "?fields=borders;alpha2Code"
		countryResponse, err := http.Get(specCountryURL)
		if err != nil {
			http.Error(w, "Error during request to CountryAPI", http.StatusInternalServerError)
			log.Println("Failed to get bordering country data from CountryAPI")
			return
		}
		defer countryResponse.Body.Close()

		var specCountryData []Country

		err = json.NewDecoder(countryResponse.Body).Decode(&specCountryData)
		if err != nil {
			http.Error(w, "Error during request to CountryAPI", http.StatusInternalServerError)
			log.Println("Failed to decode country data from CountryAPI, about the specified country")
			log.Println(countryResponse.Body)
			log.Println(specCountryData)
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
			log.Println("Error during request to UniversityAPI")
			return
		}

		// Initialize a slice to hold all of the response objects
		var response []Response

		// Loop through each university in the response, only
		// including universities that are in a border country
		for _, uni := range uniData {
			if uni.Alpha2Code != countryData[0].Alpha2Code {
				// Get country data from "RESTcountries" API based on the university's country
				countryData, err := getCountries([]string{uni.Alpha2Code})
				if err != nil {
					http.Error(w, "Error during request to CountryAPI", http.StatusInternalServerError)
					log.Println("Error during request to CountryAPI")
					return
				}

				// Check if the university's country is a border country
				if contains(borders, countryData[0].Alpha2Code) {
					// Create the response object for this university
					responseObj := Response{
						Name:      uni.Name,
						Country:   uni.Country,
						IsoCode:   uni.Alpha2Code,
						WebPages:  uni.WebPages,
						Languages: countryData[0].Languages,
						Maps:      countryData[0].Maps,
					}
					// Append the response object to the response slice
					response = append(response, responseObj)
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

	*/
}

// Check if a string slice contains a given string
func contains(slice []string, s string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}
