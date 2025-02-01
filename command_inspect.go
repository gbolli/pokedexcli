package main

import "fmt"


func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("you must provide a pokemon name")
		return nil
	}

	pname := args[0]

	pokemonDetail, ok := cfg.pokedex[pname]

	if !ok {
		fmt.Println("You have not caught %s yet.\n", pname)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemonDetail.Name)
	fmt.Printf("Height: %v\n", pokemonDetail.Height)
	fmt.Printf("Weight: %v\n", pokemonDetail.Weight)
	fmt.Print("Stats:\n")
	for _, stat := range pokemonDetail.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, ptype := range pokemonDetail.Types {
		fmt.Printf("  - %s\n", ptype.Type.Name)
	}

	return nil
}