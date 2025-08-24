package main

import (
	"time"

	"github.com/d-darac/pokedex/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := config{
		caughtPokemon: map[string]pokeapi.PokeAPIPokemon{},
		pokeapiClient: pokeapiClient,
	}

	runRepl(&cfg)
}
