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
