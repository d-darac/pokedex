package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (pokeClient *PokeClient) ListLocationAreas(pageUrl *string) (PokeAPINamedList, error) {
	url := baserUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, exists := pokeClient.cache.Get(url); exists {
		namedList := PokeAPINamedList{}
		err := json.Unmarshal(val, &namedList)
		if err != nil {
			return PokeAPINamedList{}, fmt.Errorf("error unmarshaling json data: %v", err)
		}

		return namedList, nil
	}

	res, err := pokeClient.sendHttpRequest(http.MethodGet, url, nil, nil)
	if err != nil {
		return PokeAPINamedList{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeAPINamedList{}, fmt.Errorf("error reading body: %v", err)
	}

	namedList := PokeAPINamedList{}
	if err := json.Unmarshal(data, &namedList); err != nil {
		return PokeAPINamedList{}, fmt.Errorf("error unmarshaling json data: %v", err)
	}

	pokeClient.cache.Add(url, data)

	return namedList, nil
}
