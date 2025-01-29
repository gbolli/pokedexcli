package main

import (
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	var next string
	if cfg.next == nil {
		next = "https://pokeapi.co/api/v2/location-area/"
	} else {
		next = *cfg.next
	}

	locationList := pokeapi.GetLocations(next)

	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}

	cfg.next = locationList.Next
	cfg.previous = locationList.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationList := pokeapi.GetLocations(*cfg.previous)
	
	for _, location := range locationList.Results {
		fmt.Println(location.Name)
	}

	cfg.next = locationList.Next
	cfg.previous = locationList.Previous

	return nil
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:		"help",
			description: "Displays a help message",
			callback:	commandHelp,
		},
		"map": {
			name:		"map",
			description: "Displays the next 20 locations in the Pokemon world",
			callback:	commandMap,
		},
		"mapb": {
			name:		"mapb",
			description: "Displays the previous 20 locations in the Pokemon world",
			callback:	commandMapb,
		},
	}
}