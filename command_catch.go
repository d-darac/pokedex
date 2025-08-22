package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing argument pokemon_name\nusage:\ncatch <pokemon_name>")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	pokemonSpecies, err := cfg.pokeapiClient.GetPokemonSpecies(pokemon.Species.Name)
	if err != nil {
		return err
	}

	catchPokemon(pokemonName, pokemonSpecies.CaptureRate, 100)

	return nil
}

func catchPokemon(pokemonName string, catchRate, currentHealth int) {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	catchProbability := calcCatchProbability(float64(catchRate), float64(currentHealth))
	numPossibilities := (1 / catchProbability) / 100
	rand := math.Floor(randFloat64n(numPossibilities)*100) / 100
	if rand == 0 {
		fmt.Printf("%s was caught!\n", pokemonName)
		return
	}
	fmt.Printf("%s escaped!\n", pokemonName)
}

func calcCatchProbability(catchRate float64, healthPercent float64) float64 {
	probability := (catchRate * (300 - 2*healthPercent)) / (300 * 255)
	return math.Min(1.0, math.Max(0.0, probability))
}

func randFloat64n(n float64) float64 {
	return 0.00 + (rand.Float64() * (n))
}
