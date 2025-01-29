package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	cfg := &config{}
	commands := GetCommands()
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

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}