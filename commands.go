package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, args ...string) error
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