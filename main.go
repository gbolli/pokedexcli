package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	
	for {

		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		clean := cleanInput(input)

		if len(clean) == 0 { 
			fmt.Print("No input recieved")
			continue
		}
		
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}