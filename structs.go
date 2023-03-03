package university_search

// Response Diagnosis Struct for responding to the client
type Diagnosis struct {
	UniversitiesAPI string `json:"universitiesapi"` // "<http status code for universities API>",
	CountriesAPI    string `json:"countriesapi"`    // "<http status code for restcountries API>",
	Version         string `json:"v1"`              // "v1",
	Uptime          string `json:"uptime"`          // "<time in seconds from the last service restart>"
}

// Struct for storing data from HIPO Universities API
type University struct {
	IsoCode  string   `json:"alpha_two_code"`
	WebPages []string `json:"web_pages"`
	Name     string   `json:"name"`
	Country  string   `json:"country"`
}

// Struct for storing data from REST Countries API
type Country struct {
	IsoCode   string            `json:"cca2"`
	Languages map[string]string `json:"languages"`
	Maps      map[string]string `json:"maps"`
	Borders   []string          `json:"borders"`
}

// Response Struct for responding to the client
type Response struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	IsoCode   string            `json:"isocode"`
	WebPages  []string          `json:"webpages"`
	Languages map[string]string `json:"languages"`
	Maps      map[string]string `json:"maps"`
}
