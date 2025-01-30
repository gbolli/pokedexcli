package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:		"explore",
			description: "Explore a location",
			callback:	commandExplore,
		},
	}
}