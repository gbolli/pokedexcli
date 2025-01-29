package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type locations struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(url string) locations {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}
	if err != nil {
		log.Fatal(err)
	}

	locationList := locations{}
	if err = json.Unmarshal(data, &locationList); err != nil {
	    log.Fatal(err)
	}

	return locationList
}