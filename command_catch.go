package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("you must provide a pokemon name")
		return nil
	}

	pokemonURL := "https://pokeapi.co/api/v2/pokemon/" + args[0]

	pokemonDetails, err := cfg.client.GetPokemon(pokemonURL)
	if err != nil { return err }

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	// fmt.Printf("Base experience: %v\n", pokemonDetails.BaseExperience)

	exp := pokemonDetails.BaseExperience

	var chanceToCatch int
	switch (true) {
	case exp > 300:
		chanceToCatch = 15
	case exp > 200:
		chanceToCatch = 25
	case exp > 150:
		chanceToCatch = 40
	case exp > 100:
		chanceToCatch = 60
	default:
		chanceToCatch = 75
	}

	// fmt.Printf("Chance: %v\n", chanceToCatch)
	catchAttempt := rand.IntN(100)
	// fmt.Printf("Attempt: %v\n", catchAttempt)
	pokemonIsCaught := catchAttempt < chanceToCatch

	if !pokemonIsCaught { 
		fmt.Printf("%s escaped!\n", pokemonDetails.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonDetails.Name)
	cfg.pokedex[pokemonDetails.Name] = pokemonDetails

	return nil
}


// "https://pokeapi.co/api/v2/location-area/canalave-city-area"