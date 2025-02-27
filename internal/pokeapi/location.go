package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocation(url string) (location, error) {
	locationDetails := location{}

	// Check cache
	cacheData, ok := c.gameCache.Get(url)
	if ok {
		err := json.Unmarshal(cacheData, &locationDetails)
		if err != nil { log.Fatal(err) }

		return locationDetails, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return locationDetails, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		errText := "Can't find that location. Check spelling."
		resError := errors.New(errText)
		return locationDetails, resError
	}

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &locationDetails); err != nil {
	    log.Fatal(err)
	}

	return locationDetails, nil
}