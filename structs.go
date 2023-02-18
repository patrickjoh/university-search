package university_search

type Diagnosis struct {
	UniversitiesAPI string `json:"universitiesapi"` // "<http status code for universities API>",
	CountriesAPI    string `json:"countriesapi"`    // "<http status code for restcountries API>",
	Version         string `json:"v1"`              // "v1",
	Uptime          string `json:"uptime"`          // "<time in seconds from the last service restart>"
}

type Univeristies struct {
}

type Neighbours struct {
}
