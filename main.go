package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

func main() {

	cfg := &config{}
	commands := GetCommands()
	pcache := pokecache.NewCache(time.Second *8)
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		clean := cleanInput(input)

		if len(clean) == 0 { continue }
		
		command, ok := commands[clean[0]]
		if !ok { 
			fmt.Println("Unknown command") 
			continue
		}

		err := command.callback(cfg, &pcache)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}