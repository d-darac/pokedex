package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing argument location_area_name\nusage:\n\texplore <location_area_name>")
	}

	locationArea := args[0]

	res, err := cfg.pokeapiClient.GetLocationArea(locationArea)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationArea)
	fmt.Println("Found Pokemon:")
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
