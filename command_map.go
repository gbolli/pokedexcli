package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	var next string
	if cfg.next == nil {
		next = "https://pokeapi.co/api/v2/location-area/"
	} else {
		next = *cfg.next
	}

	locationList := cfg.client.GetLocations(next)

	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}

	cfg.next = locationList.Next
	cfg.previous = locationList.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationList := cfg.client.GetLocations(*cfg.previous)
	
	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}

	cfg.next = locationList.Next
	cfg.previous = locationList.Previous

	return nil
}
