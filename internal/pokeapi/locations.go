package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (c *Client) GetLocations(url string) (locations, error) {
	locationList := locations{}

	cacheData, ok := c.gameCache.Get(url)
	if ok {
		err := json.Unmarshal(cacheData, &locationList)
		if err != nil { return locationList, err }

		return locationList, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return locationList, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		errText := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		resError := errors.New(errText)
		return locationList, resError
	}

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &locationList); err != nil {
	    return locationList, err
	}

	return locationList, nil
}