package prog2005assignment1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Joke struct {
	IconUrl string `json:"icon_url"`
	Id      string `json:"id"`
	Url     string `json:"url"`
	Value   string `json:"value"`
}

/*
Simple REST client demo
*/
func main() {

	// URL to invoke
	url := "https://api.chucknorris.io/jokes/random"

	// Create new request
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf("Error in creating request:", err.Error())
	}

	// Setting content type -> effect depends on the service provider
	r.Header.Add("content-type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	res, err := client.Do(r)
	//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
	if err != nil {
		fmt.Errorf("Error in response:", err.Error())
	}

	// HTTP Header content
	fmt.Println("Status:", res.Status)
	fmt.Println("Status code:", res.StatusCode)

	fmt.Println("Content type:", res.Header.Get("content-type"))
	fmt.Println("Protocol:", res.Proto)

	// Print raw output - should only be used in development or for debugging
	/*output, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("Error when reading response: ", err.Error())
	}

	fmt.Println(string(output))*/

	// Decoding JSON
	decoder := json.NewDecoder(res.Body)
	var mp Joke
	if err := decoder.Decode(&mp); err != nil {
		log.Fatal(err)
	}

	// Printing decoded output
	fmt.Println(mp)
}
