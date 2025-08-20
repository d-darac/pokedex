package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (pokeClient *PokeClient) ListLocationAreas(pageUrl *string) (PokeAPIListResponse, error) {
	url := baserUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, exists := pokeClient.cache.Get(url); exists {
		listRes := PokeAPIListResponse{}
		err := json.Unmarshal(val, &listRes)
		if err != nil {
			return PokeAPIListResponse{}, fmt.Errorf("error unmarshaling json data: %v", err)
		}

		return listRes, nil
	}

	res, err := sendHttpRequest(pokeClient, http.MethodGet, url, nil, nil)
	if err != nil {
		return PokeAPIListResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeAPIListResponse{}, fmt.Errorf("error reading body: %v", err)
	}

	listRes := PokeAPIListResponse{}
	if err := json.Unmarshal(data, &listRes); err != nil {
		return PokeAPIListResponse{}, fmt.Errorf("error unmarshaling json data: %v", err)
	}

	pokeClient.cache.Add(url, data)

	return listRes, nil
}
