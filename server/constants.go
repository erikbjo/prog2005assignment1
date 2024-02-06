package server

// Default path for the server
const DEFAULT_PATH = "/"
const VERSION = "v1"
const LIBRARY_STATS_PATH = "/librarystats/" + VERSION
const BOOK_COUNT_PATH = LIBRARY_STATS_PATH + "/bookcount/"
const READERSHIP_PATH = LIBRARY_STATS_PATH + "/readership/"
const STATUS_PATH = LIBRARY_STATS_PATH + "/status/"

// External API endpoints hosted by Christopher
const GUTENDEX_API = "http://129.241.150.113:8000/books/"
const LANGUAGE_API = "http://129.241.150.113:3000/language2countries/"
const COUNTRIES_API = "http://129.241.150.113:8080/v3.1"

// External API endpoints hosted by owners, while the local server is down
const GUTENDEX_API_REMOTE = "https://gutendex.com/books/"
const LANGUAGE_API_REMOTE = "https://restcountries.com/v3.1/all"
const COUNTRIES_API_REMOTE = "https://restcountries.com/v3.1/all"
