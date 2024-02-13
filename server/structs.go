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

type BookCount struct {
	Language string  `json:"language"`
	Books    int     `json:"books"`
	Authors  int     `json:"authors"`
	Fraction float64 `json:"fraction"`
}

type Readership struct {
	Country    string `json:"country"`
	Isocode    string `json:"isocode"`
	Books      int    `json:"books"`
	Authors    int    `json:"authors"`
	Readership int    `json:"readership"`
}

type Book struct {
	Id        int      `json:"id"`
	Title     string   `json:"title"`
	Authors   []Person `json:"authors"`
	Languages []string `json:"languages"`
}

type Person struct {
	BirthYear int    `json:"birth_year"`
	DeathYear int    `json:"death_year"`
	Name      string `json:"name"`
}

type GutendexResult struct {
	Count   int    `json:"count"`
	Results []Book `json:"results"`
}
