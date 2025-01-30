package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("you must provide a location name")
		return nil
	}

	locationURL := "https://pokeapi.co/api/v2/location-area/" + args[0]

	locationDetails, err := cfg.client.GetLocation(locationURL)
	if err != nil { return err }

	fmt.Printf("Exploring %s...\n", locationDetails.Name)

	if locationDetails.PokemonEncounters == nil {
		fmt.Println("No Pokemon found...")
		return nil
	}

	fmt.Println("Found Pokemon:")

	for _, p := range locationDetails.PokemonEncounters {
		fmt.Println(p.Pokemon.Name)
	}

	return nil
}


// "https://pokeapi.co/api/v2/location-area/canalave-city-area"