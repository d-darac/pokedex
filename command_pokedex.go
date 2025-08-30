package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, v := range cfg.caughtPokemon {
		fmt.Printf("%s-%s\n", getPadding(1, 2), v.Name)
	}

	return nil
}
