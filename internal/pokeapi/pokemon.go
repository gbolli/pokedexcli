package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetPokemon(url string) (Pokemon, error) {
	pokemonDetails := Pokemon{}

	// Check cache
	cacheData, ok := c.gameCache.Get(url)
	if ok {
		err := json.Unmarshal(cacheData, &pokemonDetails)
		if err != nil { log.Fatal(err) }

		return pokemonDetails, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return pokemonDetails, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		errText := "Can't find that pokemon. Check spelling."
		resError := errors.New(errText)
		return pokemonDetails, resError
	}

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &pokemonDetails); err != nil {
	    log.Fatal(err)
	}

	return pokemonDetails, nil
}