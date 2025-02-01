package main

import "fmt"


func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty, go catch some Pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	
	for _, pm := range cfg.pokedex {
		fmt.Printf("  - %s\n", pm.Name)
	}

	return nil
}