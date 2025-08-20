package main

import (
	"time"

	"github.com/d-darac/pokedex/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second)

	cfg := config{
		pokeapiClient: pokeapiClient,
	}

	runRepl(&cfg)
}
