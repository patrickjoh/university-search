package university_search

// Diagnosis Struct for responding to the client
type Diagnosis struct {
	UniversitiesAPI string `json:"universities_api"` // "<http status code for universities API>",
	CountriesAPI    string `json:"countries_api"`    // "<http status code for rest_countries API>",
	Version         string `json:"v1"`               // "v1",
	Uptime          string `json:"uptime"`           // "<time in seconds from the last service restart>"
}

// University Struct for storing data from Universities API
type University struct {
	IsoCode  string   `json:"alpha_two_code"`
	WebPages []string `json:"web_pages"`
	Name     string   `json:"name"`
	Country  string   `json:"country"`
}

// Country Struct for storing data from REST Countries API
type Country struct {
	IsoCode    string            `json:"cca2"`
	Alpha3Code string            `json:"cca3"`
	Languages  map[string]string `json:"languages"`
	Maps       map[string]string `json:"maps"`
	Border     []string          `json:"borders"`
}

// Response Struct for responding to the client
type Response struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	IsoCode   string            `json:"iso_code"`
	WebPages  []string          `json:"webpages"`
	Languages map[string]string `json:"languages"`
	Maps      map[string]string `json:"maps"`
}
