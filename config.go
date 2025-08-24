package main

import (
	"github.com/d-darac/pokedex/internal/pokeapi"
)

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
	pokeapiClient    pokeapi.PokeClient
	caughtPokemon    map[string]pokeapi.PokeAPIPokemon
}
