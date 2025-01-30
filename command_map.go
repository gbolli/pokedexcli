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

	locationList, err := cfg.client.GetLocations(next)
	if err != nil { return err }

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

	locationList, err := cfg.client.GetLocations(*cfg.previous)
	if err != nil { return err }
	
	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}

	cfg.next = locationList.Next
	cfg.previous = locationList.Previous

	return nil
}
