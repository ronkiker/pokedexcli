package main

import (
	"time"

	"github.com/ronkiker/pokedexcli/internal/pokeapi"
)

// holds stateful data for command functions
type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&cfg)
}
