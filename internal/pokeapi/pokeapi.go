package pokeapi

import (
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type Client struct {
	gameCache		pokecache.Cache
	httpClient		http.Client
}

// TODO:  break down url components
// const baseURL = "https://pokeapi.co/api/v2"

func NewClient(reapInterval time.Duration) Client {
	return Client{
		gameCache: pokecache.NewCache(reapInterval),
		httpClient: http.Client{},
	}
}