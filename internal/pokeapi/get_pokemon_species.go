package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (pokeClient *PokeClient) GetPokemonSpecies(pokemonSpeciesName string) (PokeAPIPokemonSpecies, error) {
	url := baserUrl + "/pokemon-species/" + pokemonSpeciesName

	if val, exists := pokeClient.cache.Get(url); exists {
		pokemonSpecies := PokeAPIPokemonSpecies{}
		if err := json.Unmarshal(val, &pokemonSpecies); err != nil {
			return PokeAPIPokemonSpecies{}, fmt.Errorf("error unmarshaling json data: %v", err)
		}

		return pokemonSpecies, nil
	}

	res, err := pokeClient.sendHttpRequest(http.MethodGet, url, nil, nil)
	if err != nil {
		return PokeAPIPokemonSpecies{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeAPIPokemonSpecies{}, fmt.Errorf("error reading body: %v", err)
	}

	pokemonSpecies := PokeAPIPokemonSpecies{}
	if err := json.Unmarshal(data, &pokemonSpecies); err != nil {
		return PokeAPIPokemonSpecies{}, fmt.Errorf("error unmarshaling json data: %v", err)
	}

	pokeClient.cache.Add(url, data)

	return pokemonSpecies, nil
}
