package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/d-darac/pokedex/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing argument pokemon_name\nusage:\ncatch <pokemon_name>")
	}

	pokemonName := args[0]

	pokemon, exists := cfg.caughtPokemon[pokemonName]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fields := cfg.commandsDefaultArgs["inspect"]["--raw"]

	if len(args[1:]) == 1 && args[1:][0] == "--raw" {
		if err := displayRawPokemonDetails(pokemon, fields); err != nil {
			return err
		}
		return nil
	}

	if len(args[1:]) > 1 && args[1:][0] == "--raw" {
		if err := displayRawPokemonDetails(pokemon, args[2:]); err != nil {
			return err
		}
		return nil
	}

	if err := displayPokemonDetails(pokemon, cfg); err != nil {
		return err
	}

	return nil
}

func displayPokemonDetails(pokemon pokeapi.PokeAPIPokemon, cfg *config) error {
	pokemonMap, err := structToMap(pokemon, nil, cfg.commandsDefaultArgs["inspect"]["--raw"])
	if err != nil {
		return err
	}

	lines := constructLines(pokemonMap)

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}

func displayRawPokemonDetails(pokemon pokeapi.PokeAPIPokemon, fields []string) error {
	pokemonMap, err := structToMap(pokemon, nil, fields)
	if err != nil {
		return err
	}

	lines := constructLinesRaw(pokemonMap, nil)

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}

func constructLines(pokemonMap map[string]interface{}) []string {
	lines := []string{
		fmt.Sprintf("Name: %s", pokemonMap["Name"]),
		fmt.Sprintf("Height: %d", pokemonMap["Height"]),
		fmt.Sprintf("Weight: %d", pokemonMap["Weight"]),
	}

	lines = append(lines, "Stats:")
	for i := range pokemonMap["Stats"].([]map[string]interface{}) {
		stat := pokemonMap["Stats"].([]map[string]interface{})[i]
		statName := stat["Stat"].(map[string]interface{})["Name"]

		line := fmt.Sprintf("-%s: %d", statName, stat["BaseStat"])
		withPadding := fmt.Sprintf("%s%s", getPadding(1, 2), line)
		lines = append(lines, withPadding)
	}

	lines = append(lines, "Types:")
	for i := range pokemonMap["Types"].([]map[string]interface{}) {
		tp := pokemonMap["Types"].([]map[string]interface{})[i]
		typeName := tp["Type"].(map[string]interface{})["Name"]

		line := fmt.Sprintf("-%s", typeName)
		withPadding := fmt.Sprintf("%s%s", getPadding(1, 2), line)
		lines = append(lines, withPadding)
	}

	return lines
}

func constructLinesRaw(pokemonMap map[string]interface{}, nestLevel *int) []string {
	lines := []string{}

	if nestLevel == nil {
		nestLevel = new(int)
		*nestLevel = 0
	}

	for k, v := range pokemonMap {
		valType := reflect.TypeOf(v)

		if v == nil {
			line := fmt.Sprintf("%s: null", k)
			withPadding := fmt.Sprintf("%s%s", getPadding(*nestLevel, 2), line)
			lines = append(lines, withPadding)
			continue
		}

		if valType.Kind() == reflect.Map {
			line := fmt.Sprintf("%s:", k)
			withPadding := fmt.Sprintf("%s%s", getPadding(*nestLevel, 2), line)
			lines = append(lines, withPadding)

			*nestLevel++

			lns := constructLinesRaw(v.(map[string]interface{}), nestLevel)

			*nestLevel--

			lines = append(lines, lns...)
			continue
		}

		if valType.Kind() == reflect.Slice {
			sliceLen := reflect.ValueOf(v).Len()
			if sliceLen == 0 {
				line := fmt.Sprintf("%s: null", k)
				withPadding := fmt.Sprintf("%s%s", getPadding(*nestLevel, 2), line)
				lines = append(lines, withPadding)
				continue
			}

			line := fmt.Sprintf("%s:", k)
			withPadding := fmt.Sprintf("%s%s", getPadding(*nestLevel, 2), line)
			lines = append(lines, withPadding)

			*nestLevel++
			for i := range sliceLen {
				elementType := reflect.ValueOf(v).Index(i).Type()

				if isNil := reflect.ValueOf(v).Index(i).IsNil(); isNil {
					continue
				}

				if elementType.Kind() == reflect.Map {
					line := fmt.Sprintf("%d:", i)
					withPadding := fmt.Sprintf("%s%s", getPadding(*nestLevel, 2), line)
					lines = append(lines, withPadding)

					*nestLevel++

					lns := constructLinesRaw(v.([]map[string]interface{})[i], nestLevel)

					lines = append(lines, lns...)
					*nestLevel--
					continue
				}

				val := reflect.ValueOf(v).Index(i).Interface()
				line := fmt.Sprintf("%d: %v", i, val)
				withPadding := fmt.Sprintf("%s%s", getPadding(*nestLevel, 2), line)
				lines = append(lines, withPadding)
			}
			*nestLevel--
			continue
		}

		line := fmt.Sprintf("%s: %v", k, v)
		withPadding := fmt.Sprintf("%s%s", getPadding(*nestLevel, 2), line)
		lines = append(lines, withPadding)
	}

	return lines
}
