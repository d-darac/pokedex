package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (pokeClient *PokeClient) GetLocationArea(locationName string) (PokeAPILocationArea, error) {
	url := baserUrl + "/location-area/" + locationName

	if val, exists := pokeClient.cache.Get(url); exists {
		locationArea := PokeAPILocationArea{}
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return PokeAPILocationArea{}, fmt.Errorf("error unmarshaling json data: %v", err)
		}

		return locationArea, nil
	}

	res, err := pokeClient.sendHttpRequest(http.MethodGet, url, nil, nil)
	if err != nil {
		return PokeAPILocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeAPILocationArea{}, fmt.Errorf("error reading body: %v", err)
	}

	locationArea := PokeAPILocationArea{}
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return PokeAPILocationArea{}, fmt.Errorf("error unmarshaling json data: %v", err)
	}

	pokeClient.cache.Add(url, data)

	return locationArea, nil
}
