package university_search

type Diagnosis struct {
	UniversitiesAPI string `json:"universitiesapi"` // "<http status code for universities API>",
	CountriesAPI    string `json:"countriesapi"`    // "<http status code for restcountries API>",
	Version         string `json:"v1"`              // "v1",
	Uptime          string `json:"uptime"`          // "<time in seconds from the last service restart>"
}

type Universities struct {
	// Oppdater/endre datatypene etter hva som blir mest riktig.
	// Mest URLer
	Name      string   `json:"name"`
	Country   string   `json:"country"`
	Isocode   string   `json:"isocode"`
	Webpages  string   `json:"webpages"`
	Languages []string `json:"languages"`
	Map       string   `json:"map"`
}

type Hipo struct {
}

type Countries struct {
}
