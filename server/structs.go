package server

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
	Count    int    `json:"count"`
	Results  []Book `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type Country struct {
	Iso31661Alpha3 string `json:"ISO3166_1_Alpha_3"`
	Iso31661Alpha2 string `json:"ISO3166_1_Alpha_2"`
	OfficialName   string `json:"Official_Name"`
	RegionName     string `json:"Region_Name"`
	SubRegionName  string `json:"Sub_Region_Name"`
	Language       string `json:"Language"`
}

type CountryFromRestCountries struct {
	Population int `json:"population"`
}
