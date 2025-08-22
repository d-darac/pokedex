package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (pokeClient *PokeClient) GetPokemon(pokemonName string) (PokeAPIPokemon, error) {
	url := baserUrl + "/pokemon/" + pokemonName

	if val, exists := pokeClient.cache.Get(url); exists {
		pokemon := PokeAPIPokemon{}
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return PokeAPIPokemon{}, fmt.Errorf("error unmarshaling json data: %v", err)
		}

		return pokemon, nil
	}

	res, err := pokeClient.sendHttpRequest(http.MethodGet, url, nil, nil)
	if err != nil {
		return PokeAPIPokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeAPIPokemon{}, fmt.Errorf("error reading body: %v", err)
	}

	pokemon := PokeAPIPokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return PokeAPIPokemon{}, fmt.Errorf("error unmarshaling json data: %v", err)
	}

	pokeClient.cache.Add(url, data)

	return pokemon, nil
}
