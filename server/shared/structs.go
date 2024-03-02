package shared

// Status struct, used to return status information about the server
type Status struct {
	GutendexAPI  int     `json:"gutendexapi"`
	LanguageAPI  int     `json:"languageapi"`
	CountriesAPI int     `json:"countriesapi"`
	Version      string  `json:"version"`
	Uptime       float64 `json:"uptime"`
}

// BookCount struct, used to return book count information
type BookCount struct {
	Language string  `json:"language"`
	Books    int     `json:"books"`
	Authors  int     `json:"authors"`
	Fraction float64 `json:"fraction"`
}

// Readership struct, used to return readership information
type Readership struct {
	Country    string `json:"country"`
	Isocode    string `json:"isocode"`
	Books      int    `json:"books"`
	Authors    int    `json:"authors"`
	Readership int    `json:"readership"`
}

// Book struct, used to decode JSON from Gutendex API
type Book struct {
	Id        int      `json:"id"`
	Title     string   `json:"title"`
	Authors   []Person `json:"authors"`
	Languages []string `json:"languages"`
}

// Person struct, used to decode JSON from Gutendex API
type Person struct {
	BirthYear int    `json:"birth_year"`
	DeathYear int    `json:"death_year"`
	Name      string `json:"name"`
}

// GutendexResult struct, used to decode JSON from Gutendex API
type GutendexResult struct {
	Count    int    `json:"count"`
	Results  []Book `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// Country struct, used to decode JSON from Christopher's Language2Countries API
type Country struct {
	Iso31661Alpha3 string `json:"ISO3166_1_Alpha_3"`
	Iso31661Alpha2 string `json:"ISO3166_1_Alpha_2"`
	OfficialName   string `json:"Official_Name"`
	RegionName     string `json:"Region_Name"`
	SubRegionName  string `json:"Sub_Region_Name"`
	Language       string `json:"Language"`
}

// CountryFromRestCountries struct, used to decode JSON from RestCountries API
type CountryFromRestCountries struct {
	Population int `json:"population"`
}
