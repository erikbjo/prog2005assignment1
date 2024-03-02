package shared

// Default path for the server
const DefaultPath = "/"
const Version = "v1"
const LibraryStatsPath = "/librarystats/" + Version
const BookCountPath = LibraryStatsPath + "/bookcount/"
const ReadershipPath = LibraryStatsPath + "/readership/"
const StatusPath = LibraryStatsPath + "/status/"

// External API endpoints hosted by Christopher
const GutendexApi = "http://129.241.150.113:8000/books/"
const RestCountriesApi = "http://129.241.150.113:8080/v3.1"
const LanguageApi = "http://129.241.150.113:3000/language2countries/"

// External API endpoints hosted by owners, while the local server is down
const GutendexApiRemote = "https://gutendex.com/books/"
const RestCountriesApiRemote = "https://restcountries.com/v3.1/"

// Constants for the current API in use, because Christopher's API is not always available
const CurrentGutendexApi = GutendexApi
const CurrentRestCountriesApi = RestCountriesApi
