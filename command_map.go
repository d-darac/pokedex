package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	if cfg.nextLocationsURL == nil && cfg.prevLocationsURL != nil {
		return errors.New("you have reached the end of the list")
	}

	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = res.Next
	cfg.prevLocationsURL = res.Previous

	for i := range res.Results {
		fmt.Printf("%v", res.Results[i].Name)
		fmt.Println()
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = res.Next
	cfg.prevLocationsURL = res.Previous

	for i := range res.Results {
		fmt.Printf("%v", res.Results[i].Name)
		fmt.Println()
	}

	return nil
}
