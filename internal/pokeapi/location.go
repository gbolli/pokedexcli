package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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
		// errText := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
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