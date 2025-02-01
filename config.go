package main

import (
	"pokedexcli/internal/pokeapi"
)

type config struct {
	client		pokeapi.Client
	next		*string
	previous	*string
	pokedex		map[string]pokeapi.Pokemon
}