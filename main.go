package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
	"time"
)

func main() {

	cfg := config{
		client: pokeapi.NewClient(time.Second *8),
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	
	commands := GetCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		clean := cleanInput(input)

		if len(clean) == 0 { continue }
		
		// command
		command, ok := commands[clean[0]]
		if !ok { 
			fmt.Println("Unknown command") 
			continue
		}

		// arguments
		args := []string{}
		if len(clean) > 1 {
			args = clean[1:]
		}

		// callback
		err := command.callback(&cfg, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}


// ideas
// 
// - move to area
// - can only catch pokemon that are in that area
// - they disappear for a time after attemt?
// - look at api for other ideas