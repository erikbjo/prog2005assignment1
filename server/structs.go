package server

type Coordinates struct {
	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
}

type Location struct {
	Name        string      `json:"name"`
	Postcode    string      `json:"code"`
	Country     string      `json:"country,omitempty"`
	Geolocation Coordinates `json:"location,omitempty"`
}

type Status struct {
	GutendexAPI  int     `json:"gutendexapi"`
	LanguageAPI  int     `json:"languageapi"`
	CountriesAPI int     `json:"countriesapi"`
	Version      string  `json:"version"`
	Uptime       float64 `json:"uptime"`
}
