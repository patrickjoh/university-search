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

		// example.com/neighbourunis/{country_name}/{university_name}
		// Get country data from "RESTcountries" API based country name given in URL
		countryData, err := getCountries(parts[4:])
		if err != nil {
			http.Error(w, "Error during request to CountryAPI", http.StatusInternalServerError)
			log.Println("Error during request to CountryAPI")
			return
		}


		// Get university data from UniversityAPI
		uniData := getUniversities(parts[5:])
		if err != nil {
			http.Error(w, "Error during request to UniversityAPI", http.StatusInternalServerError)
			log.Println("Error during request to UniversityAPI")
			return
		}

		// Loop through each Country in Borders and get the country alpha2 code
		// Append each country to the back of the NEIGHBOUR_PATH + parts[4:] with "," as a delimiter
		for _, country := range countryData[0].Border {
			// Get country data from "RESTcountries" API based on the country's ISO code
			borderData, err := getCountries(country)
			if err != nil {
				http.Error(w, "Error during request to CountryAPI", http.StatusInternalServerError)
				log.Println("Error during request to CountryAPI")
				return
			}
			// Append each country to the back of the NEIGHBOUR_PATH + parts[4:] with "," as a delimiter
			NEIGHBOUR_PATH = NEIGHBOUR_PATH + parts[4:] + "," + borderData[0].Alpha2Code
		}

		// Initialize a slice to hold all of the response objects
		var response []Response

		// Loop through each university in the response, excluding universities
		// from the country given in the URL
		for _, uni := range uniData {

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
